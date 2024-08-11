# Markdown Site

Display nested folders of `.md` files as a static website with support for
LaTeX equations via mathjax and syntax highlighting in code blocks (online only).

Why do this when I could just view the files in VSCode or another modern
text editor? Good question. I like (neo)vim instead and while I have used the
markdown viewer I wanted to navigate a little more easily. This isn't a good reason,
but hey, it was a fun little project.

## Usage

To serve files from a directory:

```bash
mdview <flags> DIRECTORY
```

![mdview screenshot](screenshot.png?raw=true "mdview screenshot")


Set the `-d` flag to use darkmode.

![mdview screenshot darkmode](screenshot_dark.png?raw=true "mdview screenshot darkmode")

For help:

```bash
mdview -h
```

## Installation

```bash
go install github.com/kevinnowland/mdview@latest
```
