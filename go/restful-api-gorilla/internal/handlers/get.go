package handlers

import (
	"net/http"

	"gitlab.com/prasetiyohadi/go-usvc/internal/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
//
// responses:
// 	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("Get all records")
	rw.Header().Add("Content-Type", "application/json")

	cur := r.URL.Query().Get("currency")

	// fetch the products from the database
	prods, err := p.productDB.GetProducts(cur)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// serialize the list to JSON
	err = data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Error("Unable to serialize product", "error", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Returns a product from the database
//
// responses:
// 	200: productResponse
// 	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(r)
	cur := r.URL.Query().Get("currency")

	p.l.Debug("Get record", "id", id)

	prod, err := p.productDB.GetProductByID(id, cur)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Error("Unable to fetch product", "error", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Error("Unable to fetch product", "error", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Error("Unable to serialize product", "error", err)
	}
}
