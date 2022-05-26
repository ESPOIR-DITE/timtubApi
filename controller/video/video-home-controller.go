package video

import (
	"github.com/go-chi/chi"
	"net/http"
	"timtubeApi/config"
	videoController "timtubeApi/controller/video/video"
	categoryController "timtubeApi/controller/video/video-category"
	videoCategoryController "timtubeApi/controller/video/video-category"
	videoCommentController "timtubeApi/controller/video/video-comment"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/category", categoryController.Home(app))
	mux.Mount("/video-category", videoCategoryController.Home(app))
	mux.Mount("/video", videoController.Home(app))
	mux.Mount("/comment", videoCommentController.Home(app))
	return mux
}

func homeHandler(app *config.Env) http.Handler {
	return nil
}
