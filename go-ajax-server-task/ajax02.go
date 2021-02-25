package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type album struct {
	Titulo  string `json: titulo`
	Artista string `json: artista`
	Año     int    `json: año`
}

var albumes = []album{{
	Titulo:  "Album de oro",
	Artista: "Los Acosta",
	Año:     1995,
}, {
	Titulo:  "Bocanada",
	Artista: "Cerati, G.",
	Año:     1999,
}, {
	Titulo:  "OK, COMPUTER",
	Artista: "Radiohead",
	Año:     1997,
}}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ajax02.html")
}

func reto(w http.ResponseWriter, r *http.Request) {
	words := [3]string{"hello1", "hello2", "hello3"}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(words[0])
	case "POST":
		r.ParseForm()
		position, _ := strconv.ParseInt(r.Form.Get("index"), 10, 64)
		dato := r.Form.Get("n")
		words[position] = dato
		fmt.Println(words)
	}

}

func reto2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		json.NewEncoder(w).Encode(albumes)

	}

}
func reto3(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		titulo := r.Form.Get("Titulo")
		artista := r.Form.Get("Artista")
		año, _ := strconv.Atoi(r.Form.Get("Año"))

		newAlbum := album{
			Titulo:  titulo,
			Artista: artista,
			Año:     año,
		}
		albumes = append(albumes, newAlbum)
		fmt.Println("new data: ", titulo)
		json.NewEncoder(w).Encode(len(albumes))

	}

}
func reto3b(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		num := r.Form.Get("num")
		n, _ := strconv.Atoi(r.Form.Get("num"))
		fmt.Println("num:")
		fmt.Println(num)
		var albumSolo = albumes[n:len(albumes)]
		fmt.Println(albumSolo)
		json.NewEncoder(w).Encode(albumSolo)

	}

}

func main() {
	http.HandleFunc("/reto3b", reto3b)
	http.HandleFunc("/reto3", reto3)
	http.HandleFunc("/reto2", reto2)
	http.HandleFunc("/reto", reto)
	http.HandleFunc("/", mostrarHTML)

	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
