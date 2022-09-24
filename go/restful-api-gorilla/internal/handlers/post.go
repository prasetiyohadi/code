package handlers

import (
	"net/http"

	"gitlab.com/prasetiyohadi/go-usvc/internal/data"
)

// swagger:route POST /products products createProduct
// Creates a new product
//
// responses:
// 	200: productResponse
// 	422: errorValidation
// 	502: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Debug("Inserting product: %#v\n", prod)
	p.productDB.AddProduct(prod)
}
