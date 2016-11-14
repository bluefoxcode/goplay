package main

import (
	"fmt"
	"net/http"

	"github.com/bluefoxcode/goplay/boot"
	"github.com/bluefoxcode/goplay/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	info := boot.LoadConfig()
	dev := true

	renderer := handlers.NewRenderRenderer("views", []string{".html"}, handlers.Funcs, dev)

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Handle("/", handlers.Index(renderer)).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, http.StatusNotFound, "not_found", map[string]string{
			"url": r.URL.String(),
		}, "layout")
	})

	n := negroni.Classic()
	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%s", info.Port)
	n.Run(hostStr)

}
