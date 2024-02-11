package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zeroidentidad/ecom-api/internal/ecommerce" // moduleName/internal/pkgName

	"github.com/gorilla/mux"
)

// @Summary Upsert: Agregar/modificar producto si recibe id en body
// @Tags Product
// @Accept json
// @Produce json
// @Param request body http.Product true "Product data"
// @Success 200 {object} http.Product "OK"
// @Router /api/product [patch]
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

// @Summary Obtener un producto
// @Tags Product
// @Produce json
// @Param id path string true "Id producto"
// @Success 200 {object} http.Product "OK"
// @Router /api/product/{id} [get]
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

// @Summary Listar productos con parametros opcionales de: orden por precio y buscar por nombre
// @Tags Product
// @Produce json
// @Param order query string false "order options" Enums(ASC, DESC)
// @Param search query string false "string name" example(string)
// @Success 200 {array} http.Product "OK"
// @Router /api/products [get]
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

// @Summary Eliminar producto
// @Tags Product
// @Produce json
// @Param id path string true "Id Producto"
// @Success 200 {object} http.message "OK"
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
		message{Message: "successfully deleted"},
	); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
