package main

import (
	"encoding/json"
	"net/http"
)

type book struct {
	Titulo    string `json:"titulo"`
	Autor     string `json:"autor"`
	Oferta    bool   `json:"oferta"`
	Cantidad  int    `json:"quantity"`
	Editorial string `json:"editorial"`
}

var books = []book{
	{
		Titulo:    "Does androids dream of electric sheeps?",
		Autor:     "Philip K Dick",
		Oferta:    false,
		Cantidad:  5,
		Editorial: "De bolsillo",
	},
	{
		Titulo:    "Perdida",
		Autor:     "Gillian Flynn",
		Oferta:    true,
		Cantidad:  15,
		Editorial: "Si",
	},
	{
		Titulo:    "Dune",
		Autor:     "Frank Herbert",
		Oferta:    false,
		Cantidad:  10,
		Editorial: "Patito",
	},
}

type album struct {
	Titulo  string `json: titulo`
	Artista string `json: artista`
	Año     int    `json: año`
	Url     string `json: url`
}

var albumes = []album{{
	Titulo:  "Album de oro",
	Artista: "Los Acosta",
	Año:     1995,
	Url:     "https://images-na.ssl-images-amazon.com/images/I/414HS3CT8FL.jpg",
}, {
	Titulo:  "Bocanada",
	Artista: "Cerati, G.",
	Año:     1999,
	Url:     "https://upload.wikimedia.org/wikipedia/en/1/14/Bocanada_%28Cerati%29.jpg",
}, {
	Titulo:  "OK, COMPUTER",
	Artista: "Radiohead",
	Año:     1997,
	Url:     "https://upload.wikimedia.org/wikipedia/en/thumb/b/ba/Radioheadokcomputer.png/220px-Radioheadokcomputer.png",
}}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(albumes)
	}

}

func getBooks(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(books)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	http.HandleFunc("/getAlbums", getAlbums)
	http.HandleFunc("/getBooks", getBooks)

	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
