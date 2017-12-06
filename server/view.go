package server

import "io/ioutil"

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
  html, _ := LoadFile(view.Pwd + "assets/index.html")
  return html
}
