package app

import (
    "ebookmod/pkg/api"

    "github.com/go-chi/chi/v5"
)

func APIRouter() chi.Router {

    r := chi.NewRouter()

    r.Route("/", func(r chi.Router) {
        r.Get("/hello", api.DemoHandler)
    })

    return r
}