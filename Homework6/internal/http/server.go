package http

import (
	"context"
	"encoding/json"
	"fmt"
	"hw6/project/internal/models"
	"hw6/project/internal/store"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store
	Address string
}

func NewServer(ctx context.Context, address string, store store.Store) *Server {
	return &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
		store:       store,
		Address: address,
	}
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	r.Post("/books", func(w http.ResponseWriter, r *http.Request) {
		book := new(models.Book)
		if err := json.NewDecoder(r.Body).Decode(book); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Create(r.Context(), book)
	})
	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		books, err := s.store.All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, books)
	})
	r.Get("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		book, err := s.store.ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, book)
	})
	r.Put("/books", func(w http.ResponseWriter, r *http.Request) {
		book := new(models.Book)
		if err := json.NewDecoder(r.Body).Decode(book); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Update(r.Context(), book)
	})
	r.Delete("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Delete(r.Context(), id)
	})
	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)
	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done() 
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}
	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	<-s.idleConnsCh
}