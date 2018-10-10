package controller

import (
	"fmt"
	"net/http"

	"github.com/fahmijufri/jismi/formatter"
)

func Liveness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, formatter.SuccessResponse().Stringify())
}
