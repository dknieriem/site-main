package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var router *chi.Mux

type Page struct {
	ID      int           `json:"id"`
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
}

func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	var err error
	router.Get("/", GetHomepage)
	router.Route("/pages", func(r chi.Router) {
		r.Route("/{pageID}", func(r chi.Router) {
			r.Use(PageCtx)
			r.Get("/", GetPage) // GET /pages/1234
		})
	})
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))
	router.Handle("/favicon.ico", fs)
	err = http.ListenAndServe(":8005", router)
	catch(err)
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func PageCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageID := chi.URLParam(r, "pageID")
		page, err := fileGetPage(pageID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "page", page)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetHomepage(w http.ResponseWriter, r *http.Request) {
	homepage, err := fileGetHomepage()
	catch(err)

	t, _ := template.ParseFiles("templates/html.go.tmpl", "templates/homepage.go.tmpl")
	err = t.Execute(w, homepage)
	catch(err)
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	page := r.Context().Value("page").(*Page)
	t, _ := template.ParseFiles("templates/html.go.tmpl", "templates/page.go.tmpl")
	err := t.Execute(w, page)
	catch(err)
}
