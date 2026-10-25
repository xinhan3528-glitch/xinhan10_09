package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"xinhan10_09/Internal/app"
	"xinhan10_09/Internal/routes"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("APP已开始使用")

	var port int

	flag.IntVar(&port, "port", 8080, "use to change server port")
	flag.Parse()

	r := routes.SetupRoutes(app)

	s := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.Logger.Printf("Server开始跑: port为 %d\n", port)

	err = s.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
