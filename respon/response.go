package respon

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	res := Response{
		Message: err.Error(),
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}
