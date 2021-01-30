package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

	"api/pkg/methods/compilations"
	"api/pkg/methods/recipes"
	"api/pkg/methods/users"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func logHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		log.Debug().
			Str("request", (fmt.Sprintf("%q", x))).
			Msg(r.Method)
		rec := httptest.NewRecorder()
		fn(rec, r)
		log.Debug().
			Str("response", (fmt.Sprintf("%q", rec.Body))).
			Msg(r.Method)

		// this copies the recorded response to the response writer
		for k, v := range rec.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		rec.Body.WriteTo(w)
	}
}

func main() {
	router := EpicMux()

	srv := &http.Server{
		Handler: router,
		Addr:    ":80",
	}

	log.Info().Msg("Starting server")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("Started server")
}

// EpicMux creating our router
func EpicMux() http.Handler {
	router := mux.NewRouter()

	// Users router
	router.HandleFunc("/user/{id}", logHandler(users.GetOne)).Methods("GET")
	router.HandleFunc("/usernameavailable/{username}", logHandler(users.GetUsernameAvailability)).Methods("GET")
	router.HandleFunc("/user/post", logHandler(users.PostOne)).Methods("POST")
	router.HandleFunc("/user/delete", logHandler(users.DeleteOne)).Methods("DELETE")
	router.HandleFunc("/users/search", logHandler(users.Search)).Methods("GET")

	// Recipes router
	router.HandleFunc("/recipe/{id}", logHandler(recipes.GetOne)).Methods("GET")
	router.HandleFunc("/recipe/post", logHandler(recipes.PostOne)).Methods("POST")
	router.HandleFunc("/recipe/delete", logHandler(recipes.DeleteOne)).Methods("DELETE")
	router.HandleFunc("/recipes/search", logHandler(recipes.Search)).Methods("GET")

	// Compilations router
	router.HandleFunc("/compilation/{id}", logHandler(compilations.GetOne)).Methods("GET")
	router.HandleFunc("/compilation/post", logHandler(compilations.PostOne)).Methods("POST")
	router.HandleFunc("/compilation/delete", logHandler(compilations.DeleteOne)).Methods("DELETE")
	router.HandleFunc("/compilations/search", logHandler(compilations.Search)).Methods("GET")

	return router
}
