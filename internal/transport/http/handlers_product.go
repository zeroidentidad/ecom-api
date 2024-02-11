package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zeroidentidad/ecom-api/internal/ecommerce" // moduleName/internal/pkgName

	"github.com/gorilla/mux"
)

func (h *Handler) UpsertProduct(w http.ResponseWriter, r *http.Request) {
	var p ecommerce.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := h.Service.UpsertProduct(r.Context(), p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := h.Service.GetProduct(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	order := r.URL.Query().Get("order")
	search := r.URL.Query().Get("search")

	p, err := h.Service.GetProducts(r.Context(), order, search)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// @Summary Eliminar Producto
// @Tags Products
// @Produce json
// @Param id path string true "ID Producto"
// @Success 200 {integer} int "OK"
// @Router /api/product/{id} [delete]
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteProduct(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(
		struct{ message string }{"successfully deleted"},
	); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
