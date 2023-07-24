package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (httpa *Adapter) InitRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/addition", httpa.DtoHandler(httpa.GetAddition))
	r.Get("/subtraction", httpa.DtoHandler(httpa.GetSubtraction))
	r.Get("/multiplication", httpa.DtoHandler(httpa.GetMultiplication))
	r.Get("/division", httpa.DtoHandler(httpa.GetDivision))

	return r
}
