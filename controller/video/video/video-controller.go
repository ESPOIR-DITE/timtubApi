package video_comment

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"timtubeApi/config"
	"timtubeApi/controller/util"
	"timtubeApi/domain"
	repository "timtubeApi/storage/video/video-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", get(app))
	r.Get("/delete/{id}", delete(app))
	r.Post("/create", create(app))
	r.Post("/create", update(app))
	r.Get("/getAll", getAll(app))
	return r
}

func delete(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteVideo(id)
			result, err := json.Marshal(role)
			if err != nil {
				fmt.Println("couldn't marshal")
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
	}
}

func getAll(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetVideos()
		result, err := json.Marshal(user)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error writing bytes")))
			return
		}
	}
}

func create(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.Video{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetVideoObject(data)
		response := repository.CreateVideo(object)
		if response.Id == "" {
			fmt.Println("error creating video")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating video")))
			return
		}
		result, err := json.Marshal(repository.GetVideoObject(response))
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
	}
}

func update(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.Video{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetVideoObject(data)
		response := repository.UpdateVideo(object)
		if response.Id == "" {
			fmt.Println("error creating Video")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Video")))
			return
		}
		result, err := json.Marshal(repository.GetVideoObject(response))
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
	}
}

func get(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object := repository.GetVideo(id)
			result, err := json.Marshal(object)
			if err != nil {
				fmt.Println("couldn't marshal")
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
	}
}
