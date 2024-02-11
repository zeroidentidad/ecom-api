package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zeroidentidad/ecom-api/internal/ecommerce" // moduleName/internal/pkgName

	"github.com/gorilla/mux"
)

func (h *Handler) PostItemCart(w http.ResponseWriter, r *http.Request) {
	var c ecommerce.Cart
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, err := h.Service.PostItemCart(r.Context(), c)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(c); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetItemsCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	if userid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, err := h.Service.GetItemsCart(r.Context(), userid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(c); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// @Summary Eliminar Producto Carrito
// @Tags Carts
// @Produce json
// @Param userid path string true "ID Usuario"
// @Param productid path string true "ID Producto"
// @Success 200 {integer} int "OK"
// @Router /api/cart/{userid}/item/{productid} [delete]
func (h *Handler) DeleteItemCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]
	productid := vars["productid"]

	if userid == "" || productid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteItemCart(r.Context(), userid, productid)
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
