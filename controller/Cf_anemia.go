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
	var dt_json []respon.GetHitung
	for dec.More() {
		var GetData respon.GetHitung
		// decode an array value (Message)
		err := dec.Decode(&GetData)
		if err != nil {
			log.Fatal(err)
		}
		// log.Print(GetData)
		dt_json = append(dt_json, GetData)
	}
	var hasil_combine = make(map[string]float32)
	// kita buat perulangan untuk mendapatkan values
	rows, err := db.Query("select kode_penyakit, nama_penyakit from spk_anemia_penyakit")
	iterasi := 0
	for rows.Next() {
		pnykt := respon.Penyakit{}
		if err := rows.Scan(&pnykt.Kode_penyakit, &pnykt.Nama_penyakit); err != nil {
			log.Fatal(err.Error())

		}

		sql2 := `select spk_anemia_gejala_penyakit.nama_gejala, spk_anemia_gejala_penyakit.kode_gejala, spk_anemia_rule.md
			from spk_anemia_penyakit
			join spk_anemia_rule on spk_anemia_rule.Penyakitnya_id = spk_anemia_penyakit.id
			join spk_anemia_gejala_penyakit on spk_anemia_rule.Gejalanya_id = spk_anemia_gejala_penyakit.id
			where spk_anemia_penyakit.Kode_penyakit=?`
		rows2, err2 := db.Query(sql2, pnykt.Kode_penyakit)
		iterasi2 := 0
		fmt.Println("Nama Penyakit ", pnykt.Nama_penyakit)
		fmt.Println("=====================================")
		var kombin []float32
		for rows2.Next() {
			rulenya := respon.Gejala{}
			err2 = rows2.Scan(&rulenya.Gejala, &rulenya.Kode_gejala, &rulenya.MD)
			if err2 != nil {
				log.Panic(err2)
			}
			for _, value := range dt_json {
				if value.Kode_gejala == rulenya.Kode_gejala {
					fmt.Printf("kode %v : %v x %v = %v\n", rulenya.Kode_gejala, value.Persentase_user, rulenya.MD, value.Persentase_user*rulenya.MD)
					kombin = append(kombin, value.Persentase_user*rulenya.MD)
				}
			}

			iterasi2++
		}
		fmt.Println("============Hasil Combine============")
		var hasilcombin float32 = 0
		for key := range kombin {
			if key == 0 {
				hasilcombin = kombin[key] + kombin[key+1]*(1.0-kombin[key])
				// fmt.Printf("%v + %v * (1 - %v) = %v \n", kombin[key], kombin[key+1], kombin[key], hasilcombin)
				if len(kombin)-1 == 1 {
					hasil_combine[pnykt.Nama_penyakit] = hasilcombin
					// fmt.Println("============Hasil Kombin============")
					fmt.Println("Hasil combin = ", hasilcombin)
					break
				}
			} else {
				if key+1 == len(kombin) {
					hasil_combine[pnykt.Nama_penyakit] = hasilcombin
					// fmt.Println("============Hasil Kombin============")
					fmt.Println("Hasil combin = ", hasilcombin)
					break
				}
				hasilcombin = hasilcombin + kombin[key+1]*(1.0-hasilcombin)
				// fmt.Printf("%v + %v * (1 - %v) = %v \n", hasilcombin, kombin[key+1], hasilcombin, hasilcombin)
			}
		}
		// fmt.Println(len(kombin))
		fmt.Println("=====================================")
		iterasi++
	}
	var hitung []respon.Hasil_hitung
	var terbesar respon.Hasil_hitung
	var tmp float32 = 0
	for key, value := range hasil_combine {
		var i respon.Hasil_hitung
		i.Nama_penyakit = key
		i.Hasil_perhitungan = value
		if tmp < value {
			terbesar.Nama_penyakit = key
			terbesar.Hasil_perhitungan = value
			tmp = value
		}
		hitung = append(hitung, i)

	}
	respon.MessageResponse(w, "Sukses", hitung, terbesar, 200)

}
