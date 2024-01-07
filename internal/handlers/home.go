package handlers

import (
	"construccion_demo/internal/models"
	"construccion_demo/web/template"
	"log/slog"
	"net/http"
)

func (s service) Home(w http.ResponseWriter, r *http.Request) {
	var products *[]models.Product
	err := r.ParseForm()
	if err != nil {
		slog.Error(err.Error())
		slog.Error("Error parsing form values", "URL", r.URL, "host", r.Host, "remote_address", r.RemoteAddr, "request_uri", r.RequestURI)
		return
	}

	if r.FormValue("search") != "" {
		products, err = s.ProductService.SearchProductsByName(r.FormValue("search"))
	} else {
		products, err = s.ProductService.GetProducts()
	}

	if err != nil {
		return
	}

	err = templates.Home(*products).Render(r.Context(), w)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	return
}
