package main

import (
	"io"
	"net/http"
	"os"
	"mime"
)

func (s Serwer) uploadHandler(w http.ResponseWriter, r *http.Request) {
	contentDisposition := r.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil{
		return
	}
	nazwaPliku := params["filename"]
	nazwaUzytkownika := params["username"]

	file, err := os.Create("./Repozytorium/" + nazwaUzytkownika + "/" + nazwaPliku)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}
	//w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}
