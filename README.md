<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>

# Expert System With Certainty Factor Using GoLang

Sistem pendukung keputusan dengan certainty factor menggunakan GoLang database MySql. Sistem ini berbasis REST API, jadi untuk komunikasi data menggunakan JSON.

## Installation
- Silahkan download/clone repository ini
- Import data [.sql](go_cf.sql) terlebih dahulu.
- ubah file [.env](.env) sesuai konfigurasi kalian
- Jalankan main.go dengan perintah ```go run main.go ```

## NOTE!!!
- Data yang didapat bukanlah data dari pakar untuk pembobotan MB (<i>measure of belief / ukuran kepercayaan</i>) maupun MD (<i> measure of disbelief / ukuran ketidakpercayaan</i>)
- Data penyakit, rule dan gejala didapat dari [jurnal ini](https://prpm.trigunadharma.ac.id/public/fileJurnal/hpxu6%20-%20Trinanda.pdf)
- Silahkan ubah data tersebut sesuai studi kasus masing-masing

## Input Output
1. Input
Contoh POST data JSON seperti dibawah ini:(<i>Data harus array</i>)
```
[
	{
		"kode_gejala": "G01",
		"persentase_user": 0.8
	},
	{
		"kode_gejala": "G02",
		"persentase_user": 0.5
	},
	{
		"kode_gejala": "G03",
		"persentase_user": 0.6
	},
	{
		"kode_gejala": "G08",
		"persentase_user": 0.4
	}
]
```
Setelah itu akan diproses oleh sistem.

2. Output
Hasil perhitungan akan berupa JSON, seperti dibawah ini:
```
{
  "status": 200,
  "pesan": "Sukses",
  "list_penyakit": [
    {
      "nama_penyakit": "Anemia Aplastik",
      "hasil_perhitungan": 0.29776
    },
    {
      "nama_penyakit": "Anemia Defisiensi Zat besi",
      "hasil_perhitungan": 0.4145536
    },
    {
      "nama_penyakit": "Anemia Kremis/Kronik",
      "hasil_perhitungan": 0.24400002
    }
  ],
  "hasil_akhir": {
    "nama_penyakit": "Anemia Defisiensi Zat besi",
    "hasil_perhitungan": 0.4145536
  }
}
```

## Hitung Manual
- Silahkan lihat file [.ods](manual.ods) ini untuk perhitungan manual dari contoh inputan diatas 

## Referensi
1. [SISTEM PAKAR PENDIAGNOSA PENYAKIT ANAK MENGGUNAKAN CERTAINTY FACTOR (CF)](https://ejournal.unsrat.ac.id/index.php/JIS/article/view/705/0)
2. [Perancangan Aplikasi Sistem Pakar Penyakit Roseola Dengan Menggunakan Metode Certainty Factor](https://www.ejurnal.stmik-budidarma.ac.id/index.php/JSON/article/view/1956)
3. [SISTEM PAKAR DIAGNOSA PENYAKIT PADA AYAM MENGGUNAKAN METODE CERTAINTY FACTOR](https://www.scribd.com/document/431280343/Dokumentasi-Sistem-Pakar-Ayam-Skripsi)
4. [SISTEM PAKAR penyelesaian metode Certainty Factor
](http://ariecandra02.blogspot.com/2017/05/sistem-pakar-penyelesaian-metode_64.html)


## Disclaimer

* <b>Dilarang keras</b> di perjual-belikan, source ini saya publikasi untuk keperluan belajar saja.

## Donation

* Bagi yang ingin berdonasi terbentuknya sistem ini, siapapun, berapapun, saya ucapkan terimakasih sebanyak-banyaknya. Via Gopay / Dana.

### Gopay<br>
<img src="img/gpy.png" height="400"> <br>

### Dana<br>
<img src="img/dana.png" height="350">
