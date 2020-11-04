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

type GetAll struct {
	Status      string     `json:"status"`
	Penyakitnya []Penyakit `json:"penyakitnya"`
}

type Gejala struct {
	// Id       int64  `db:"pId"`
	Kode_gejala string `db:"kode_gejala" json:"kode_gejala"`
	Gejala      string `db:"nama_gejala" json:"nama_gejala"`
}

type Penyakit struct {
	// Id       int64  `db:"pId"`
	Nama_penyakit string   `db:"nama_penyakit" json:"nama_penyakit"`
	Kode_penyakit string   `db:"kode_penyakit" json:"kode_penyakit"`
	Gejalanya     []Gejala `json:"gejalanya"`
	// Bobotnya string `db:"bobotnya"`
	// Nama_gejala Coba
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

func MessageResponse(w http.ResponseWriter, message string, data interface{}, status int) {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func GetAllData(w http.ResponseWriter, data interface{}) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Fatal(err)
	}
}
