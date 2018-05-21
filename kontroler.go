package main

import (
	"mime"
	"net/http"
)

var IDPliku int = 0

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	IDPliku++
	contentDisposition := r.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		http.Error(w, http.ErrBodyNotAllowed.Error(), http.StatusNotAcceptable)
		return
	}
	nazwaPliku := params["filename"]
	nazwaUzytkownika := params["username"]

	plikUzytkownika := PlikUzytkownika{
		ID:               IDPliku,
		NazwaPliku:       nazwaPliku,
		NazwaUzytkownika: nazwaUzytkownika,
		StrumienDoPliku:  r.Body,
	}
	plikDoZuploadowaina, czyJestJuzTakiWpisUzytkownika := KolejkaPlikowDlaUzytkownikow[nazwaUzytkownika]
	if czyJestJuzTakiWpisUzytkownika {
		plikDoZuploadowaina.PlikUzytkownika <- plikUzytkownika
		plikDoZuploadowaina.KanalZwalnajacy <- true
		return
	}
	kanalZwalniajacy := make(chan bool, 1)
	kanalDoWrzucanai := make(chan PlikUzytkownika, 1)
	kanalDoWrzucanai <- plikUzytkownika
	kanalZwalniajacy <- true
	KolejkaPlikowDlaUzytkownikow[nazwaUzytkownika] = PlikUploadowany{
		PlikUzytkownika: kanalDoWrzucanai,
		KanalZwalnajacy: kanalZwalniajacy,
	}

}