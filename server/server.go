package server

type Server struct {
  Port string
  Router *ActionRouter
  View *View
}

func NewServer(port string) *Server {
  server := Server {
    Port: port,
    Router: NewRouter(),
    View: NewView(),
  }

  server.Router.Route("/", server.Index)
  server.Router.Route("/ytdl", server.Ytdl)

  return &server
}

func (server *Server) Run() {
  server.Router.Serve(server.Port)
}

func (server *Server) Index() string {
  return server.View.LoadIndex()
}

func (server *Server) Ytdl() string {
  return server.View.LoadYtdl()
}
