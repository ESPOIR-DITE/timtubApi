package user

import (
	"github.com/go-chi/chi"
	"net/http"
	"timtubeApi/config"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))

	return mux
}

func homeHandler(app *config.Env) http.Handler {
	return nil
}
