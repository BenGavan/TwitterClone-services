package main

func (s *server) setupRoutes() {
	s.router.HandleFunc("/", s.handleIndex)
	s.router.HandleFunc(s.routes.NewAuthProfile, s.handleNewAuthProfile)
}