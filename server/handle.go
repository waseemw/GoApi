package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler[B any, R any] func(body *B) R

func Handle[B any, R any](url string, handler Handler[B, R]) {
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(404)
			return
		}
		b := new(B)
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&b); err != nil {
			log.Fatal(err)
		}
		encoder := json.NewEncoder(w)
		var res R
		res = handler(b)

		if err := encoder.Encode(res); err != nil {
			log.Fatal(err)
		}

	})
}
