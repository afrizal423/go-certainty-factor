package controller

import (
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
	// sql := `select spk_anemia_penyakit.kode_penyakit, spk_anemia_penyakit.nama_penyakit, spk_anemia_rule.bobotnya
	// from spk_anemia_penyakit
	// join spk_anemia_rule on spk_anemia_rule.Penyakitnya_id = spk_anemia_penyakit.id`
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
		// pnykt.Spk_anemia_rules = append(pnykt.Spk_anemia_rules, rulenya)
		sql2 := `select spk_anemia_gejala_penyakit.nama_gejala
		from spk_anemia_penyakit
		join spk_anemia_rule on spk_anemia_rule.Penyakitnya_id = spk_anemia_penyakit.id
		join spk_anemia_gejala_penyakit on spk_anemia_rule.Gejalanya_id = spk_anemia_gejala_penyakit.id
		where spk_anemia_penyakit.Kode_penyakit=?`
		// eksekusi sql statement
		rows2, err2 := db.Query(sql2, pnykt.Kode_penyakit)
		// log.Println(rows2, err2)

		for rows2.Next() {
			rulenya := respon.Gejala{}
			err2 = rows2.Scan(&rulenya.Gejala)
			if err2 != nil {
				log.Panic(err2)
			}
			// log.Println(rulenya)
			pnykt.Gejalanya = append(pnykt.Gejalanya, rulenya)
		}

		// log.Println(pnykt.Kode_penyakit)
		getdt.Penyakitnya = append(getdt.Penyakitnya, pnykt)
	}

	// log.Println("users...")
	// log.Println(getdt)
	respon.GetAllData(w, getdt)
	// json.NewEncoder(w).Encode(getdt)

}
