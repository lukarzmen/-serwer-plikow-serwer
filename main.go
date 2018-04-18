package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	serwer := Serwer{
		NazwaKlienta:                "lmedyk",                //os.Getenv("NAZWA_KLIENTA"),
		SciezkaDoFolderuUzytkownika: "/Users/lmedyk/Desktop", //os.Getenv("SCIEZKA"),
	}
	router := mux.NewRouter()
	router.
		Path("/upload").
		Methods("POST").
		HandlerFunc(serwer.uploadHandler)
	err := http.ListenAndServe(":80", router)
	if err != nil {

	}
}
