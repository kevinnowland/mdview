# Markdown Site

Display nested folders of `.md` files as a static website with
support for mathjax (online only).

## Usage

```bash
mdview PATH/TO/MARKDOWN/FOLDERS
```

![mdview screenshot](screenshot.png?raw=true "mdview screenshot")


Set the `-d` flag to use darkmode.

![mdview screenshot darkmode](screenshot_dark.png?raw=true "mdview screenshot darkmode")

For help:

```bash
mdview -h
```

Note: embedded images aren't working.

## Installation

```bash
go install github.com/kevinnowland/mdview@latest
```

## TODO

- Add support for images
