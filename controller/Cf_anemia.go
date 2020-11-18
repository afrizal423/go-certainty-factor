package controller

import (
	"encoding/json"
	"fmt"
	"go_certainty_factor/config"
	"go_certainty_factor/respon"
	"log"
	"net/http"
)

func ListPenyakit(w http.ResponseWriter, r *http.Request) {

	db, err := config.CreateConnection()
	defer db.Close()

	if err != nil {
		respon.ErrorResponse(w, 500, err)
		return
	}

	getdt := &respon.GetAll{}
	getdt.Status = "Sukses"
	getdt.Penyakitnya = make([]respon.Penyakit, 0)
	sql := `select spk_anemia_penyakit.kode_penyakit, spk_anemia_penyakit.nama_penyakit
	from spk_anemia_penyakit `
	// db.Select(&user, sql)
	rows, err := db.Query(sql)

	for rows.Next() {
		pnykt := respon.Penyakit{}
		// rulenya := Spk_anemia_rule{}
		err = rows.Scan(&pnykt.Kode_penyakit, &pnykt.Nama_penyakit)
		if err != nil {
			log.Panic(err)
		}
		sql2 := `select spk_anemia_gejala_penyakit.nama_gejala, spk_anemia_gejala_penyakit.kode_gejala
		from spk_anemia_penyakit
		join spk_anemia_rule on spk_anemia_rule.Penyakitnya_id = spk_anemia_penyakit.id
		join spk_anemia_gejala_penyakit on spk_anemia_rule.Gejalanya_id = spk_anemia_gejala_penyakit.id
		where spk_anemia_penyakit.Kode_penyakit=?`
		// eksekusi sql statement
		rows2, err2 := db.Query(sql2, pnykt.Kode_penyakit)
		// log.Println(rows2, err2)

		for rows2.Next() {
			rulenya := respon.Gejala{}
			err2 = rows2.Scan(&rulenya.Gejala, &rulenya.Kode_gejala)
			if err2 != nil {
				log.Panic(err2)
			}
			pnykt.Gejalanya = append(pnykt.Gejalanya, rulenya)
		}
		getdt.Penyakitnya = append(getdt.Penyakitnya, pnykt)
	}

	respon.GetAllData(w, getdt)

}

func HitungCF(w http.ResponseWriter, r *http.Request) {

	db, err := config.CreateConnection()
	defer db.Close()
	if err != nil {
		log.Panic(err)
	}

	dec := json.NewDecoder(r.Body)

	// kita unmarshall bracket array dulu

	dec.Token()

	// kita buat perulangan untuk mendapatkan values

	// var dt_penyakit []respon.Penyakit
	// var dt_rule []respon.Gejala
	// fmt.Println(len(dt_penyakit))
	// fmt.Println(dt_rule)

	for dec.More() {
		var GetData respon.GetHitung
		// decode an array value (Message)
		err := dec.Decode(&GetData)
		if err != nil {
			log.Fatal(err)
		}
		// log.Print(GetData)
		fmt.Printf("%v: %v\n", GetData.Kode_gejala, GetData.Persentase_user)
		rows, err := db.Query("select kode_penyakit, nama_penyakit from spk_anemia_penyakit")
		iterasi := 0
		for rows.Next() {
			pnykt := respon.Penyakit{}
			if err := rows.Scan(&pnykt.Kode_penyakit, &pnykt.Nama_penyakit); err != nil {
				log.Fatal(err.Error())

			}

			// dt_penyakit = append(dt_penyakit, pnykt)
			sql2 := `select spk_anemia_gejala_penyakit.nama_gejala, spk_anemia_gejala_penyakit.kode_gejala
			from spk_anemia_penyakit
			join spk_anemia_rule on spk_anemia_rule.Penyakitnya_id = spk_anemia_penyakit.id
			join spk_anemia_gejala_penyakit on spk_anemia_rule.Gejalanya_id = spk_anemia_gejala_penyakit.id
			where spk_anemia_penyakit.Kode_penyakit=?`
			rows2, err2 := db.Query(sql2, pnykt.Kode_penyakit)
			iterasi2 := 0
			fmt.Println(pnykt.Nama_penyakit)
			for rows2.Next() {
				rulenya := respon.Gejala{}
				err2 = rows2.Scan(&rulenya.Gejala, &rulenya.Kode_gejala)
				if err2 != nil {
					log.Panic(err2)
				}
				// dt_rule = append(dt_rule, rulenya)
				// fmt.Println(rulenya.Kode_gejala)
				if rulenya.Kode_gejala == GetData.Kode_gejala {
					fmt.Printf("%v: %v = %v\n", GetData.Kode_gejala, GetData.Persentase_user, rulenya.Kode_gejala)
				}
				iterasi2++
			}

			iterasi++
		}
	}

}
