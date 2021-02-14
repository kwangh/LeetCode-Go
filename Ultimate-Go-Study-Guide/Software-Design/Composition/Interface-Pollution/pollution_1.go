package main

type Server interface {
	Start() error
	Stop() error
	Wait() error
}

type server struct {
	host string
}

func NewServer(host string) Server {
	// SMELL - Storing an unexported type pointer in the interface.
	return &server{host}
}

func (s *server) Start() error {
	return nil
}

func (s *server) Stop() error {
	return nil
}

func (s *server) Wait() error {
	return nil
}

func main() {
	srv := NewServer("localhost")

	srv.Start()
	srv.Stop()
	srv.Wait()
}
