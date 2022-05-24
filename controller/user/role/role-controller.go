package role

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"timtubeApi/config"
	"timtubeApi/domain"
	repository "timtubeApi/storage/user/role-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", getRole(app))
	r.Post("/create", createRole(app))
	r.Get("/getAll", getRoles(app))

	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))
	return r
}

func getRoles(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetRoles()
		result, err := json.Marshal(user)
		if err != nil {
			fmt.Println("couldn't marshal")
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func createRole(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.Role{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		var role = repository.GetRoleObject(data)
		response := repository.CreateRole(role)
		if response.Id == "" {
			fmt.Println("error creating role")
		}
		result, err := json.Marshal(repository.GetRoleObject(response))
		if err != nil {
			fmt.Println("couldn't marshal")
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}

func getRole(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.GetRole(id)
			result, err := json.Marshal(role)
			if err != nil {
				fmt.Println("couldn't marshal")
			}
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
	}
}
