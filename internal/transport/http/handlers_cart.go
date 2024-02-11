package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zeroidentidad/ecom-api/internal/ecommerce" // moduleName/internal/pkgName

	"github.com/gorilla/mux"
)

// @Summary Add product to cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param request body http.Cart true "Cart data"
// @Success 200 {object} http.Cart "OK"
// @Router /api/cart [post]
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

// @Summary List products in cart
// @Tags Cart
// @Produce json
// @Param userid path string true "user Id"
// @Success 200 {array} http.Cart "OK"
// @Router /api/cart/{userid} [get]
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

// @Summary Delete product in cart
// @Tags Cart
// @Produce json
// @Param userid path string true "user Id"
// @Param productid path string true "product Id"
// @Success 200 {object} http.message "OK"
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
		message{Message: "successfully deleted"},
	); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
