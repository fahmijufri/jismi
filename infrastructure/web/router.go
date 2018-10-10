package web

import (
	"fmt"
	"net/http"

	"github.com/fahmijufri/jismi/domain/repository"
	"github.com/fahmijufri/jismi/infrastructure/web/controller/user"

	"github.com/fahmijufri/jismi/formatter"

	"github.com/fahmijufri/jismi/infrastructure/web/controller"
	"github.com/fahmijufri/jismi/infrastructure/web/middleware"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Router(userRepository repository.UserInterface) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()

	commonHandlers := negroni.New(
		middleware.HTTPStatLogger(),
	)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, formatter.FailResponse("Not Found").Stringify())
	})

	// Probes
	router.
		Handle("/liveness", commonHandlers.With(
			negroni.WrapFunc(controller.Liveness),
		)).Methods(http.MethodGet)
	router.
		Handle("/liveness", commonHandlers.With(
			negroni.WrapFunc(controller.Liveness),
		)).Methods(http.MethodHead)
	router.
		Handle("/readiness", commonHandlers.With(
			negroni.WrapFunc(controller.Readiness),
		)).Methods(http.MethodGet)
	router.
		Handle("/readiness", commonHandlers.With(
			negroni.WrapFunc(controller.Readiness),
		)).Methods(http.MethodHead)

	// Users
	router.
		Handle("/users/{user_id}", commonHandlers.With(
			negroni.WrapFunc(user.FindByID(userRepository)),
		)).Methods(http.MethodGet)

	return router
}
