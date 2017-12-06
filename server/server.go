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

  return &server
}

func (server *Server) Run() {
  server.Router.Serve(server.Port)
}

func (server *Server) Index() string {
  return server.View.LoadIndex()
}
