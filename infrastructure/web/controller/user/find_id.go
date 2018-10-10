package user

import (
	"fmt"
	"net/http"

	"github.com/fahmijufri/jismi/domain"

	"github.com/fahmijufri/jismi/domain/repository"

	"github.com/fahmijufri/jismi/formatter"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func FindByID(repo repository.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		v := mux.Vars(r)
		id := v["user_id"]

		log := logrus.WithFields(logrus.Fields{
			"user_id": id,
		})

		query := domain.User{
			ID: id,
		}
		user, err := repo.FindOne(query)
		if err != nil {
			log.WithError(err).Errorln("Failed get user by ID")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.ObjectResponse(user).Stringify())
	}
}
