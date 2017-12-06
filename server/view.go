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

func (view *View) LoadIndex() string {
  buffer := bytes.NewBufferString("")
  rawHtml, _ := LoadFile(view.Pwd + "assets/index.html")
  rawCss, _ := LoadFile(view.Pwd + "assets/css/master.css")
  rawJS, _ := LoadFile(view.Pwd + "assets/js/app.js")

  custom := struct {
    Style templ.CSS
    Script templ.JS
  }{
    Style: templ.CSS(rawCss),
    Script: templ.JS(rawJS),
  }

  template, oops := templ.New("index").Parse(rawHtml)
  if oops != nil {
    panic(oops)
  }
  template.Execute(buffer, custom)
  outlet := buffer.String()

  return outlet
}

