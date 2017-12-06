package server

import (
  "io/ioutil"
  templ "html/template"
  "bytes"
)

type View struct {
  Pwd string
}

func NewView() *View {
  view := View {
    Pwd: "./src/github.com/ishiikurisu/homeserver/",
  }
  return &view
}

func LoadFile(path string) (string, error) {
  raw, oops := ioutil.ReadFile(path)
  outlet := ""
  if oops == nil {
    outlet = string(raw)
  }
  return outlet, oops
}

func (view *View) LoadPage(page string) string {
  buffer := bytes.NewBufferString("")
  rawHtml, _ := LoadFile(view.Pwd + page)
  rawCss, _ := LoadFile(view.Pwd + "assets/css/master.css")
  rawJS, _ := LoadFile(view.Pwd + "assets/js/app.js")

  custom := struct {
    Style templ.CSS
    Script templ.JS
  }{
    Style: templ.CSS(rawCss),
    Script: templ.JS(rawJS),
  }

  template, _ := templ.New("index").Parse(rawHtml)
  template.Execute(buffer, custom)
  outlet := buffer.String()

  return outlet
}

func (view *View) LoadIndex() string {
  return view.LoadPage("assets/index.html")
}

