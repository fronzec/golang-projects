package main

import (
	"context"
	"github.com/go-chi/chi"
	"io"
	"net/http"
)

type postsResource struct{}

func (rs postsResource) Routes() chi.Router {
	r := chi.NewRouter()
	// GET /posts - Read a list of posts
	r.Get("/", rs.List)
	// POST /posts - Create a new post
	r.Post("/", rs.Create)
	//
	r.Route("/{id}", func(r chi.Router) {
		r.Use(PostCtx)
		// GET single post by :id
		r.Get("/", rs.Get)
		// Update a single post by :id
		r.Put("/", rs.Update)
		// Delete a single post by :id
		r.Delete("/", rs.Delete)
	})
	return r
}

// Request Handler for GET /posts
func (rs postsResource) List(writer http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		http.Error(
			writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	writer.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(writer, resp.Body); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler for POST /posts
func (rs postsResource) Create(writer http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts",
		"application/json", r.Body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	writer.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(writer, resp.Body); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Middleware function that checks for the post :id in the URL parameters
func PostCtx(
	next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request) {
			ctx := context.WithValue(
				r.Context(),
				"id", chi.URLParam(r, "id"))
			next.ServeHTTP(
				w, r.WithContext(ctx))
		})
}

// Request Handler for GET /posts/{id}
func (
	rs postsResource) Get(
	w http.ResponseWriter,
	r *http.Request) {
	id := r.Context().Value("id").(string)

	// Fetch post by id
	resp,
		err := http.Get(
		"https://jsonplaceholder.typicode.com/posts/" +
			id)

	if err != nil {
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set(
		"Content-Type",
		"application/json")

	if _, err := io.Copy(
		w, resp.Body); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
}

// Request Handler for PUT /posts/{id}
func (
	rs postsResource) Update(
	w http.ResponseWriter,
	r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest(
		"PUT",
		"https://jsonplaceholder.typicode.com/posts/"+
			id, r.Body)
	req.Header.Add(
		"Content-Type", "application/json")

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set(
		"Content-Type", "application/json")

	if _, err := io.Copy(
		w,
		resp.Body); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}
}

// Request handler for DELETE /posts/{id}

func (
	rs postsResource) Delete(
	w http.ResponseWriter,
	r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		"https://jsonplaceholder.typicode.com/posts/"+
			id, nil)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set(
		"Content-Type",
		"application/json")

	if _, err := io.Copy(
		w, resp.Body); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}
}
