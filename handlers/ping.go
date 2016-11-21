package handlers

import "net/http"

func Ping(ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ren.Render(w, http.StatusOK, "ping", struct{}{}, "layout")
	})
}
