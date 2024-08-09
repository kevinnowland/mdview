package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

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

const pageTemplate = `
{{define "PAGE"}}
<!DOCTYPE html>
<html>
<head>
  <title> Markdown </title>
  <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
  <style>
  div.nav {
    border-style: hidden double hidden hidden;
    width: 15%;
    padding-left: 2.5%;
    margin-top: 2.5%;
    float: left;
  }
  div.data {
    width: 65%;
    padding-left: 2.5%;
    padding-right: 10%;
    margin-top: 2.5%;
    float: right;
  }
  </style>
</head>
<body>
<div class="body">
  <div class="nav">
    <ul>
      {{range $link := .Nav.Links}}
      <li><a href="{{$link.Href}}">{{$link.Text}}</a></li>
      {{end}}
    </ul>
  </div>
  <div class="data">
    {{.Data}}
  </div>
</div>
</body>
</html>
{{end}}
`

var dirPath string

func init() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("must provide exactly one argument: directory from which to serve markdown")
		os.Exit(1)
	}

	dirPath = args[0]
	if stat, err := os.Stat(dirPath); err != nil || !stat.IsDir() {
		fmt.Printf("provide path is not a directory: %s\n", dirPath)
		os.Exit(1)
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
		fmt.Printf("error getting markdown paths: %s\n", err.Error())
		os.Exit(1)
	}

	nav, err := GetNav(dirPath, paths)
	if err != nil {
		fmt.Printf("error getting navs: %s\n", err.Error())
		os.Exit(1)
	}

	t := template.Must(template.New("page").Parse(pageTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexPage := Page{
			Nav:  nav,
			Data: "<p>Welcome! Click a link in the nav to view markdown</p>",
		}
		err = t.ExecuteTemplate(w, "PAGE", indexPage)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 - Internal Server Error: %s", err.Error())
			return
		}
	})

	for _, path := range paths {
		p := path
		url, err := ConvertPathToUrl(dirPath, p)
		if err != nil {
			fmt.Printf("error converting path to url: %s\n", err.Error())
			os.Exit(1)
		}
		http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
			mdBytes, err := os.ReadFile(p)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "500 - Internal Server Error: %s", err.Error())
				return
			}

			mdHtml := bytes.Buffer{}
			if err := md.Convert(mdBytes, &mdHtml); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "500 - Internal Server Error: %s", err.Error())
				return
			}

			markdownPage := Page{
				Nav:  nav,
				Data: mdHtml.String(),
			}

			err = t.ExecuteTemplate(w, "PAGE", markdownPage)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "500 - Internal Server Error: %s", err.Error())
				return
			}
		})
	}

	err = http.ListenAndServe(":8080", nil)
	// TODO: handle expected shutdown more gracefully
	if err != nil {
		fmt.Printf("shutting down: %s\n", err.Error())
		os.Exit(1)
	}
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

// Assumes path of form {dirPath}/foo/bar/name.md
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

func GetMarkdownSite(dirPath string, md goldmark.Markdown) error {
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".md" {
			// parse into url
			relative, err := filepath.Rel(dirPath, path)
			if err != nil {
				return err
			}
			_ = fmt.Sprintf("/%s", relative[:len(relative)-3])

			// read and add to map
			mdBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			mdHtml := bytes.Buffer{}
			if err := md.Convert(mdBytes, &mdHtml); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
