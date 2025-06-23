package handler

import (
	"html/template"
	"net/http"

	"url-shortener/internal/shortener"
)

type Handler struct {
	tmpl    *template.Template
	service *shortener.Service
}

func New(service *shortener.Service) *Handler {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	return &Handler{tmpl: tmpl, service: service}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		original := r.FormValue("url")
		short, err := h.service.Shorten(r.Context(), original)
		if err != nil {
			http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
			return
		}
		h.tmpl.Execute(w, map[string]string{
			"ShortURL": "http://" + r.Host + "/" + short,
		})
		return
	}
	h.tmpl.Execute(w, nil)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	original, err := h.service.Resolve(r.Context(), code)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, original, http.StatusFound)
}
