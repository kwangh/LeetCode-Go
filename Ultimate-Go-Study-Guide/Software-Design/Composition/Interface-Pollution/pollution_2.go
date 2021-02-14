package main

type Server struct {
	host string
}

func NewServer(host string) *Server {
	return &Server{host}
}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) Wait() error {
	return nil
}

func main() {
	srv := NewServer("localhost")

	srv.Start()
	srv.Stop()
	srv.Wait()
}
