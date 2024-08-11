# Markdown Site

Display nested folders of `.md` files as a static website with support for
LaTeX equations via mathjax and syntax highlighting in code blocks (online only).

## Usage

To serve files from a directory:

```bash
mdview <flags> DIRECTORY
```

![mdview screenshot](screenshot.png? "mdview screenshot")


Set the `-d` flag to use darkmode.

![mdview screenshot darkmode](screenshot_dark.png? "mdview screenshot darkmode")

For help:

```bash
mdview -h
```

Note: embedded images aren't working.

## Installation

```bash
go install github.com/kevinnowland/mdview@latest
```
