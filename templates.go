package main

const PageTemplate = `
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
    color: white;
  }
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
    color: white;
  }
  a:visited {
    color: white;
  }
  a:hover {
    color: white;
    font-weight: bold;
  }
  a:active {
    color: white;
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
