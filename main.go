package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"

	"api-gateway/http/middlewares/cors"
)

func init() {
	// load our environmental variables.
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("API_URL_REST_PORT"))
	fmt.Println("Starting API Gateway...")

	// initialize http router
	r := chi.NewRouter()

	// initialize middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Init().Handler)

	// default route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := &HTTPResponseVM{
			Status:  http.StatusOK,
			Success: true,
			Message: "alive",
		}

		response.JSON(w)
	})

	r.Route("/{version:v[0-9]+}", func(r chi.Router) {
		r.Route("/{service}", func(r chi.Router) {
			r.HandleFunc("/", GatewayController)

			// workaround for https://github.com/go-chi/chi/issues/569
			const (
				depth          int    = 8
				pathSlugFormat string = "/{path%d:.*}"
			)

			var pathAccumulator string

			for i := 1; i <= depth; i++ {
				slug := fmt.Sprintf(pathSlugFormat, i)
				pathAccumulator = pathAccumulator + slug

				r.HandleFunc(pathAccumulator, GatewayController)
			}
		})
	})

	fmt.Println("Server is listening on " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
