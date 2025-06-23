package main

import (
	"context"
	"log"
	"net/http"

	"url-shortener/internal/handler"
	"url-shortener/internal/middleware"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	db, err := storage.NewPostgresStorage(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	service := shortener.New(db)
	h := handler.New(service)
	mux := http.NewServeMux()

	mux.HandleFunc("/login", h.ShowLogin)
	mux.HandleFunc("/login_submit", h.HandleLogin)
	mux.HandleFunc("/register", h.ShowRegister)
	mux.HandleFunc("/register_submit", h.HandleRegister)

	mux.Handle("/", middleware.JWTAuth(http.HandlerFunc(h.Index)))
	mux.Handle("/s/", http.HandlerFunc(h.Redirect))

	mux.HandleFunc("/favicon.ico", http.NotFound)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
