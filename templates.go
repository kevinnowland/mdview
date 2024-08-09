package main

const PageTemplate = `
{{define "PAGE"}}
<!DOCTYPE html>
<html>
<head>
  <title> Markdown </title>
  <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
  <style>
  div.data {
    width: 55%;
    padding-left: 2.5%;
    padding-right: 20%;
    margin-top: 2.5%;
    float: right;
    text-align: justify;
  }
  div.data img { 
    max-width: 100%;
    width: auto;
    height: auto;
    margin: auto;
    display: block;
  }
  div.data code {
    background-color: #e8e8e8;
    padding: 2px;
    font-family: monospace;
  }
  div.data pre {
    margin: auto;
    padding: 1%;
    overflow-x: auto;
    tab-size: 4;
    width: 98%;
    background-color: #e8e8e8;
    display: block;
    text-align: left;
  }
  div.data h1, h2, h3, h4, h5, h6 {
    border-style: none none solid none;
    border-color: #dcdcdc;
  }
  div.nav {
    border-style: hidden double hidden hidden;
    width: 15%;
    padding-left: 2.5%;
    margin-top: 2.5%;
    float: left;
  }
  div.nav ul {
    list-style-type: none;
  }
  div.nav li {
    padding-top: 5px;
    padding-bottom: 5px;
  }

  a {
    text-decoration: none;
  }
  a:link {
    color: black;
  }
  a:visited {
    color: black;
  }
  a:hover {
    color: black;
    font-weight: bold;
  }
  a:active {
    color: grey;
    font-weight: bold;
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

const PageDarkTemplate = `
{{define "PAGE"}}
<!DOCTYPE html>
<html>
<head>
  <title> Markdown </title>
  <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
  <style>
  body {
    background-color: #24292e;
    color: #f5f5f5;
  }
  div.data {
    width: 55%;
    padding-left: 2.5%;
    padding-right: 20%;
    margin-top: 2.5%;
    float: right;
    text-align: justify;
  }
  div.data img { 
    max-width: 100%;
    width: auto;
    height: auto;
    margin: auto;
    display: block;
  }
  div.data code {
    background-color: dimgrey;
    padding: 2px;
    font-family: monospace;
  }
  div.data pre {
    margin: auto;
    padding: 1%;
    overflow-x: auto;
    tab-size: 4;
    width: 98%;
    background-color: dimgrey;
    display: block;
    text-align: left;
  }
  div.data h1, h2, h3, h4, h5, h6 {
    border-style: none none solid none;
    border-color: #404040;
  }
  div.nav {
    border-style: hidden double hidden hidden;
    width: 15%;
    padding-left: 2.5%;
    margin-top: 2.5%;
    float: left;
  }
  div.nav ul {
    list-style-type: none;
  }
  div.nav li {
    padding-top: 5px;
    padding-bottom: 5px;
  }

  a {
    text-decoration: none;
  }
  a:link {
    color: #f5f5f5;
  }
  a:visited {
    color: #f5f5f5;
  }
  a:hover {
    color: #f5f5f5;
    font-weight: bold;
  }
  a:active {
    color: #f5f5f5;
    font-weight: bold;
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
