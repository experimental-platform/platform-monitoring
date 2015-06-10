package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 3001, "server port")
	flag.Parse()
	fmt.Println("Port: ", port)

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.JSON(http.StatusOK, "list")
	})

	m.Get("/:name", func(args martini.Params, r render.Render) {
		name := args["name"]
		r.JSON(http.StatusOK, name)
	})

	http.Handle("/", m)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
