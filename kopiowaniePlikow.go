package main

import (
	"io"
	"os"
	"path/filepath"
)

type PlikUzytkownika struct {
	ID               int
	NazwaUzytkownika string
	NazwaPliku       string
	StrumienDoPliku  io.Reader
}

func kopiujNaDysk(plikUzytkownika PlikUzytkownika) (zakonczono bool, err error) {
	sciezkaDoFolderuUzytkownika := filepath.Join(".", "Repozytorium", plikUzytkownika.NazwaUzytkownika)

	if _, err := os.Stat(sciezkaDoFolderuUzytkownika); os.IsNotExist(err) {
		os.Mkdir(sciezkaDoFolderuUzytkownika, os.ModePerm)
	}

	sciezkaDoPliku := filepath.Join(sciezkaDoFolderuUzytkownika, plikUzytkownika.NazwaPliku)
	file, err := os.Create(sciezkaDoPliku)
	defer file.Close()
	if err != nil {
		return true, err
	}

	//todo: losowy czas spania
	//r := rand.Intn(10)
	//time.Sleep(time.Duration(r) * time.Second)

	_, err = io.Copy(file, plikUzytkownika.StrumienDoPliku)
	if err != nil {
		return true, err
	}
	return true, nil
}
