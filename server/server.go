package server

type Server struct {
  Port string
  Router *ActionRouter
}

func NewServer(port string) *Server {
  server := Server {
    Port: port,
    Router: NewRouter(),
  }

  server.Router.Route("/", server.Index)

  return &server
}

func (server *Server) Run() {
  server.Router.Serve(server.Port)
}

func (server *Server) Index() string {
  return "Hello Joe!"
}
