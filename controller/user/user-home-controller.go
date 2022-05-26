package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"timtubeApi/config"
	roleController "timtubeApi/controller/user/role"
	userController "timtubeApi/controller/user/user"
	userAccountController "timtubeApi/controller/user/user-account"
	userSubscriptionController "timtubeApi/controller/user/user-account"
	userVideoController "timtubeApi/controller/user/user-video"
	"timtubeApi/controller/util"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/role", roleController.Home(app))
	mux.Mount("/user", userController.Home(app))
	mux.Mount("/user-account", userAccountController.Home(app))
	mux.Mount("/user-subscription", userSubscriptionController.Home(app))
	mux.Mount("/user-video", userVideoController.Home(app))

	return mux
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := json.Marshal("data")
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(response))
		if err != nil {
			return
		}
	}
}
