package http

import (
	"context"
	"encoding/json"
	"fmt"
	"hw7/internal/models"
	"hw7/internal/store"
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

	// REST
	// сущность/идентификатор
	// /accesories/bags
	// /accesories/bracelets
	r.Post("/accesories/bags", func(w http.ResponseWriter, r *http.Request) {
		bag := new(models.Bag)
		if err := json.NewDecoder(r.Body).Decode(bag); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Bags().Create(r.Context(), bag)
	})
	r.Get("/accesories/bags", func(w http.ResponseWriter, r *http.Request) {
		bags, err := s.store.Bags().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, bags)
	})
	r.Get("/accesories/bags/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		bag, err := s.store.Bags().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, bag)
	})
	r.Put("/accesories/bags", func(w http.ResponseWriter, r *http.Request) {
		bag := new(models.Bag)
		if err := json.NewDecoder(r.Body).Decode(bag); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Bags().Update(r.Context(), bag)
	})
	r.Delete("/accesories/bags/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Bags().Delete(r.Context(), id)
	})
	//Bracelet
	
	r.Get("/accesories/bracelets", func(w http.ResponseWriter, r *http.Request) {
		bracelets, err := s.store.Bracelets().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, bracelets)
	})
	r.Get("/accesories/bracelets/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		bracelet, err := s.store.Bracelets().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, bracelet)
	})
	
	r.Delete("/accesories/bracelets/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Bracelets().Delete(r.Context(), id)
	})
	r.Post("/accesories/bracelets", func(w http.ResponseWriter, r *http.Request) {
		bracelet := new(models.Bracelet)
		if err := json.NewDecoder(r.Body).Decode(bracelet); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Bracelets().Create(r.Context(), bracelet)
	})
	r.Put("/bracelets", func(w http.ResponseWriter, r *http.Request) {
		bracelet := new(models.Bracelet)
		if err := json.NewDecoder(r.Body).Decode(bracelet); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		s.store.Bracelets().Update(r.Context(), bracelet)
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
	<-s.ctx.Done() // блокируемся, пока контекст приложения не отменен

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	// блок до записи или закрытия канала
	<-s.idleConnsCh
}
