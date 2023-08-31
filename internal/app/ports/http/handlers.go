package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type App interface {
	CreateNote(context.Context, string) error
}

type app struct{}

func (c *app) CreateNote(ctx context.Context, in string) error {
	return nil
}

func NewApp() App {
	return &app{}
}

type createNoteRequest struct{}

func createNote(a App) http.HandlerFunc {
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

		err := a.CreateNote(r.Context(), "")
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
