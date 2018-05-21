package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"time"
)

type PlikUploadowany struct {
	PlikUzytkownika chan PlikUzytkownika
	KanalZwalnajacy chan bool
}

var KolejkaPlikowDlaUzytkownikow map[string]PlikUploadowany

func main() {
	KolejkaPlikowDlaUzytkownikow = make(map[string]PlikUploadowany)

	router := mux.NewRouter()
	router.
		Path("/upload").
		Methods("POST").
		HandlerFunc(uploadHandler)

	go func() {
		for {
			for _, plikDoUploadowania := range KolejkaPlikowDlaUzytkownikow {
				if len(plikDoUploadowania.PlikUzytkownika) == 0{
					continue
				}
				select {
				case <-time.After(time.Minute):
					fmt.Println("Czas minal")
				case plikUzytkownika := <-plikDoUploadowania.PlikUzytkownika:
					go func() {
						kopiujNaDysk(plikUzytkownika)
						<-plikDoUploadowania.KanalZwalnajacy
					}()
				}
			}
		}
	}()

	err := http.ListenAndServe(":82", router)
	if err != nil {
		fmt.Println("Nie można utworzyć serwera")
	}
}