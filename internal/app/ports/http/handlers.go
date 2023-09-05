package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/samarec1812/segmentation-service/internal/app/service"
)

type createNoteRequest struct{}

func createNote(a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody createNoteRequest
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&reqBody); err != nil {
			err = json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		err := a.CreateUser()
		if err != nil {
			err = json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		err = json.NewEncoder(w).Encode("")
		if err != nil {
			log.Println(err)
			return
		}
	}
}
