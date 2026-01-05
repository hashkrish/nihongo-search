package server

import (
	"fmt"
	"html/template"
	"net/http"
	"nihongo-search/internal/search"
)

type PageData struct {
	Title string
}

type Server struct {
	searchService *search.Service
	templates     map[string]*template.Template
}

func NewServer(searchService *search.Service) *Server {
	s := &Server{
		searchService: searchService,
		templates:     make(map[string]*template.Template),
	}
	s.loadTemplates()
	return s
}

func (s *Server) loadTemplates() {
	s.templates["home"] = template.Must(template.ParseFiles("templates/index.html"))
	s.templates["search"] = template.Must(template.ParseFiles("partials/search.html"))
	s.templates["pango"] = template.Must(template.ParseFiles("partials/pango/search.html"))
	s.templates["text"] = template.Must(template.ParseFiles("partials/text/search.html"))
}

func (s *Server) Start(port string) error {
	http.HandleFunc("GET /", s.handleHome)
	http.HandleFunc("GET /healthcheck", s.handleHealthCheck)
	http.HandleFunc("GET /partial/search", s.handlePartialSearch)
	http.HandleFunc("GET /partial/pango/search", s.handlePangoPartialSearch)
	http.HandleFunc("GET /partial/text/search", s.handleTextPartialSearch)

	fmt.Printf("Server running on port %s\n", port)
	return http.ListenAndServe(":"+port, nil)
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Nihongo search!",
	}
	s.templates["home"].Execute(w, data)
}

func (s *Server) handleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}

func (s *Server) handlePartialSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	fmt.Println("Query:", query)

	data := s.searchService.Search(query)
	s.templates["search"].Execute(w, data)
}

func (s *Server) handlePangoPartialSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	fmt.Println("Query:", query)

	data := s.searchService.Search(query)

	if len(data.KanjiDataList) > 3 {
		data.KanjiDataList = data.KanjiDataList[:3]
	}
	if len(data.WordDataList) > 3 {
		data.WordDataList = data.WordDataList[:3]
	}

	s.templates["pango"].Execute(w, data)
}

func (s *Server) handleTextPartialSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	fmt.Println("Query:", query)

	data := s.searchService.Search(query)

	if len(data.KanjiDataList) > 3 {
		data.KanjiDataList = data.KanjiDataList[:3]
	}
	if len(data.WordDataList) > 3 {
		data.WordDataList = data.WordDataList[:3]
	}

	s.templates["text"].Execute(w, data)
}
