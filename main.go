package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"text/template"
	"time"

	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type Link struct {
	Href string
	Text string
}

type Nav struct {
	Links []Link
}

type Page struct {
	Nav  Nav
	Data string
}

var (
	dirPath      string
	port         int
	logger       *slog.Logger
	pageTemplate *template.Template
)

func init() {
	var verbose bool
	var darkmode bool
	flag.IntVar(&port, "port", 8080, "port to run server on")
	flag.BoolVar(&verbose, "v", false, "log verbosely")
	flag.BoolVar(&darkmode, "d", false, "use darkmode")

	flag.Parse()
	args := flag.Args()

	programLevel := new(slog.LevelVar)
	if verbose {
		programLevel.Set(slog.LevelDebug)
	} else {
		programLevel.Set(slog.LevelWarn)
	}
	logger = slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: programLevel},
	))

	if len(args) != 1 {
		logger.Error("must provide exactly one argument at end of command", "numArgs", len(args))
		os.Exit(1)
	}

	dirPath = args[0]
	if stat, err := os.Stat(dirPath); err != nil || !stat.IsDir() {
		logger.Error("provide path is not a directory", "argument", dirPath)
		os.Exit(1)
	}

	if darkmode {
		pageTemplate = template.Must(template.New("page").Parse(PageDarkTemplate))
	} else {
		pageTemplate = template.Must(template.New("page").Parse(PageTemplate))
	}
}

func main() {
	md := goldmark.New(
		goldmark.WithExtensions(mathjax.MathJax),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithXHTML(),
		),
	)

	paths, err := GetMarkdownPaths(dirPath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	nav, err := GetNav(dirPath, paths)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	http.HandleFunc("/", IndexHandler(nav))

	for _, path := range paths {
		url, err := ConvertPathToUrl(dirPath, path)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		http.HandleFunc(url, MarkdownHandler(nav, path, md))
	}

	go func() {
		fmt.Printf("\n\tServing %s at http://localhost:%d/\n\n", dirPath, port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error())
		}
		logger.Info("Stopped serving new connections")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Graceful shutdown complete")
}

func GetMarkdownPaths(dirPath string) ([]string, error) {
	paths := []string{}

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && filepath.Ext(path) == ".md" {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}

	return paths, nil
}

func GetNav(dirPath string, paths []string) (Nav, error) {
	links := make([]Link, len(paths))
	for i, path := range paths {
		href, err := ConvertPathToUrl(dirPath, path)
		if err != nil {
			return Nav{}, nil
		}

		links[i] = Link{
			Href: href,
			Text: href,
		}
	}
	return Nav{Links: links}, nil
}

// assumes path of form {dirPath{/foo/bar/name.md
func ConvertPathToUrl(dirPath string, path string) (string, error) {
	relative, err := filepath.Rel(dirPath, path)
	if err != nil {
		return "", err
	}

	if filepath.Ext(path) != ".md" {
		return "", fmt.Errorf("path doesn't end in .md: %s", path)
	}

	return fmt.Sprintf("/%s", relative[:len(relative)-3]), nil
}

func IndexHandler(nav Nav) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		indexPage := Page{
			Nav:  nav,
			Data: "<p>Welcome! Click a link in the nav to view markdown</p>",
		}

		err := pageTemplate.ExecuteTemplate(w, "PAGE", indexPage)
		if err != nil {
			WriteInternalServerError(w, err)
			return
		}
	}
}

func MarkdownHandler(nav Nav, path string, md goldmark.Markdown) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mdBytes, err := os.ReadFile(path)
		if err != nil {
			WriteInternalServerError(w, err)
			return
		}

		mdHtml := bytes.Buffer{}
		if err := md.Convert(mdBytes, &mdHtml); err != nil {
			WriteInternalServerError(w, err)
			return
		}

		markdownPage := Page{
			Nav:  nav,
			Data: mdHtml.String(),
		}

		err = pageTemplate.ExecuteTemplate(w, "PAGE", markdownPage)

		if err != nil {
			WriteInternalServerError(w, err)
			return
		}
	}
}

func WriteInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "500 - Internal Server Error: %s", err.Error())
	logger.Error(err.Error())
}
