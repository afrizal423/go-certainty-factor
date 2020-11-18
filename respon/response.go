package respon

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"list_penyakit"`
	Hasil  interface{} `json:"hasil_akhir"`
}

type GetAll struct {
	Status      string     `json:"status"`
	Penyakitnya []Penyakit `json:"penyakitnya"`
}

type Gejala struct {
	Kode_gejala string  `db:"kode_gejala" json:"kode_gejala"`
	Gejala      string  `db:"nama_gejala" json:"nama_gejala"`
	MD          float32 `db:"md" json:"md"`
}

type Penyakit struct {
	Nama_penyakit string   `db:"nama_penyakit" json:"nama_penyakit"`
	Kode_penyakit string   `db:"kode_penyakit" json:"kode_penyakit"`
	Gejalanya     []Gejala `json:"gejalanya"`
}

type GetHitung struct {
	Kode_gejala     string  `json:"kode_gejala"`
	Persentase_user float32 `json:"persentase_user"`
}

type Hasil_hitung struct {
	Nama_penyakit     string  `db:"nama_penyakit" json:"nama_penyakit"`
	Hasil_perhitungan float32 `db:"hasil_perhitungan" json:"hasil_perhitungan"`
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	res := Response{
		Pesan: err.Error(),
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func MessageResponse(w http.ResponseWriter, message string, data interface{}, data2 interface{}, status int) {
	res := Response{
		Status: status,
		Pesan:  message,
		Data:   data,
		Hasil:  data2,
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
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
