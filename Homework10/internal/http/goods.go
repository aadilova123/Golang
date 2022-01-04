package http


import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	lru "github.com/hashicorp/golang-lru"
	"hw10/internal/message_broker"
	"hw10/internal/models"
	"hw10/internal/store"
	"net/http"
	"strconv"
)

type AccesoryResource struct {
	store  store.Store
	broker message_broker.MessageBroker
	cache  *lru.TwoQueueCache
}

func NewAccesoryResource(store store.Store, broker message_broker.MessageBroker, cache *lru.TwoQueueCache) *AccesoryResource {
	return &AccesoryResource{
		store:  store,
		broker: broker,
		cache:  cache,
	}
}

func (cr *AccesoryResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", cr.CreateAccesory)
	r.Get("/", cr.AllAccesories)
	r.Get("/{id}", cr.ByID)
	r.Get("/categories/{category_id}", cr.ByCategoryID)
	r.Put("/", cr.UpdateGood)
	r.Delete("/{id}", cr.DeleteAccesory)

	return r
}

func (cr *AccesoryResource) CreateAccesory (w http.ResponseWriter, r *http.Request) {
	accesory := new(models.Accesory)
	if err := json.NewDecoder(r.Body).Decode(accesory); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Accesories().Create(r.Context(), accesory); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	// Правильно пройтись по всем буквам и всем словам
	cr.broker.Cache().Purge() // в рамках учебного проекта полностью чистим кэш после создания новой категории

	w.WriteHeader(http.StatusCreated)
}

func (cr *AccesoryResource) AllAccesories(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	filter := &models.AccesoryFilter{}

	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		categoriesFromCache, ok := cr.cache.Get(searchQuery)
		if ok {
			render.JSON(w, r, categoriesFromCache)
			return
		}

		filter.Query = &searchQuery
	}

	categories, err := cr.store.Accesories().All(r.Context(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	if searchQuery != "" {
		cr.cache.Add(searchQuery, categories)
	}
	render.JSON(w, r, categories)
}

func (cr *AccesoryResource) ByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	categoryFromCache, ok := cr.cache.Get(id)
	if ok {
		render.JSON(w, r, categoryFromCache)
		return
	}

	accesory, err := cr.store.Accesories().ByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	cr.cache.Add(id, accesory)
	render.JSON(w, r, accesory)
}

func (cr *AccesoryResource) UpdateGood(w http.ResponseWriter, r *http.Request) {
	category := new(models.Accesory)
	if err := json.NewDecoder(r.Body).Decode(category); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	err := validation.ValidateStruct(
		category,
		validation.Field(&category.ID, validation.Required),
		validation.Field(&category.Name, validation.Required),
	)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Accesories().Update(r.Context(), category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	cr.broker.Cache().Remove(category.ID)
}

func (cr *AccesoryResource) DeleteAccesory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Accesories().Delete(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	cr.broker.Cache().Remove(id)
}

func (cr *AccesoryResource) ByCategoryID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "category_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	categoryFromCache, ok := cr.cache.Get(id)
	if ok {
		render.JSON(w, r, categoryFromCache)
		return
	}

	accesory, err := cr.store.Accesories().ByCategoryID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	cr.cache.Add(id, accesory)
	render.JSON(w, r, accesory)
}