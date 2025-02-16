package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"streamserver/internal/download"
	"streamserver/internal/google"
	"streamserver/internal/soundcloud"
)

type Video struct {
	Title string   `json:"title"`
	URL   *url.URL `json:"url"`
}

func Server() http.Handler {
	mux := http.NewServeMux()

	soundcloud, err := soundcloud.New()
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/yt/{query}", func(w http.ResponseWriter, r *http.Request) {
		query := r.PathValue("query")

		if query == "" {
			http.Error(w, "invalid query", http.StatusNotAcceptable)
			return
		}

		results, err := google.Search(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		out := make([]Video, len(results))
		for i := range results {
			u, err := url.Parse("https://youtube.com/watch")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			u.RawQuery = url.Values{"v": []string{results[i].VideoID}}.Encode()
			out[i] = Video{
				Title: results[i].Title,
				URL:   u,
			}
		}

		json.NewEncoder(w).Encode(out)
	})

	mux.HandleFunc("/sc/{query}", func(w http.ResponseWriter, r *http.Request) {
		query := r.PathValue("query")

		if query == "" {
			http.Error(w, "invalid query", http.StatusNotAcceptable)
			return
		}

		results, err := soundcloud.Search(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		out := make([]Video, len(results))
		for i := range results {
			u, err := url.Parse(results[i].Url)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			out[i] = Video{
				Title: results[i].Title,
				URL:   u,
			}
		}

		json.NewEncoder(w).Encode(results)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "invalid url", http.StatusNotAcceptable)
			return
		}

		w.Header().Set("Content-Type", "audio/mpeg")
		w.Header().Set("Transfer-Encoding", "chunked")

		if err := download.Download(url, ctx, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return mux
}
