# Markdown Site

Display nested folders of `.md` files as a static website with
support for mathjax (online).

Note: that this does require an internet connection for the mathjax
support.

## Usage

```bash
mdview PATH/TO/MARKDOWN/FOLDERS
```

Set the `-d` flag to use darkmode.

For help:

```bash
mdview -h
```

## Installation

```bash
go install github.com/kevinnowland/mdview@latest
```

## TODO

- Add offline mathjax support
- Add 404 for random URLs
