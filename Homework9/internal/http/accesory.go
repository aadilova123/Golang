package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-redis/redis/v8"
	"hw8/internal/models"
	"hw8/internal/store"
	"net/http"
	"strconv"
)

type AccesoryResource struct {
	repo store.AccesoriesRepository
	rdb *redis.Client
}

func NewTitleResource(repo store.AccesoriesRepository, rdb *redis.Client) *AccesoryResource {
	return &AccesoryResource{
		repo: repo,
		rdb: rdb,
	}
}

func (tr *AccesoryResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", tr.CreateCategory)
	r.Get("/", tr.AllCategories)
	r.Get("/{id}", tr.ByID)
	r.Put("/", tr.UpdateCategory)
	r.Delete("/{id}", tr.DeleteCategory)

	return r
}

func (tr *AccesoryResource) CreateCategory(w http.ResponseWriter, r *http.Request) {
	title := new(models.Accesory)
	if err := json.NewDecoder(r.Body).Decode(title); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	if err := tr.repo.Create(r.Context(), title); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, "OK")
}

func (tr *AccesoryResource) AllCategories(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	filter := &models.AccesoryFilter{}

	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		goodsFromRDB, err := tr.rdb.Get(context.Background(), searchQuery).Result()
		fmt.Printf("redis = %s\n", goodsFromRDB)
		if err == nil {
			goods := make([]*models.Accesory, 0)
			err := json.Unmarshal([]byte(goodsFromRDB), &goods)
			if err != nil {
				return
			}
			render.JSON(w, r, goods)
			return
		}
		filter.Query = &searchQuery
	}
	goods, err := tr.repo.All(r.Context(), filter)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	if searchQuery != "" {
		fmt.Println(searchQuery)
		goodsMarshal, _ := json.Marshal(goods)
		tr.rdb.Set(context.Background(), searchQuery, goodsMarshal, 0)
	}
	render.JSON(w, r, goods)
}

func (tr *AccesoryResource) ByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	title, err := tr.repo.ByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	render.JSON(w, r, title)
}


func (tr *AccesoryResource) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	title := new(models.Accesory)
	if err := json.NewDecoder(r.Body).Decode(title); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	fmt.Println(title)

	if err := validation.ValidateStruct(title, validation.Field(&title.ID, validation.Required)); err != nil{
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	if err := tr.repo.Update(r.Context(), title); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
	render.JSON(w, r, "OK")
}
func (tr *AccesoryResource) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	if err := tr.repo.Delete(r.Context(), id); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
	render.JSON(w, r, "OK")
}