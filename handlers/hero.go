package handlers

import "net/http"

func Hero(ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ren.Render(w, http.StatusOK, "hero", struct{}{}, "layout")
	})
}
