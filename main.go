package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Filme struct{
	ID string `json:"id"`
	isbn string `json:"isbn"`
	titulo string `json:"titulo"`
	Diretor *Diretor `json:"diretor"`

}

type Diretor struct{
	Primeironome string `json:"primeironome"`
	Sobrenome string `json:"sobrenome"`
}

var Filmes []Filme



func getFilmes(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Filmes)
}

func deleteFilmes(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Filmes{
		if item.ID == params["id"]{
			Filmes = append(Filmes[:index], Filmes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Filmes)
}


func getFilme(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Filmes{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func criarFilme(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var filme Filme

	_=json.NewDecoder(r.Body).Decode(&filme)
	Filme.ID = strconv.Itoa(rand.Intn(100000000))
	Filmes = append(Filmes, Filme)
	json.NewEncoder(w).Encode(filme)
}

func updateFilmes(w http.ResponseWriter, r *http.Request){
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params :=mux.Vars(r)
	//loop over the movies, range
	//delete the movie with the i.d that youi've sent
	for index, item := range Filmes{
		if item.ID ==params["id"]{
			Filmes = append(Filmes[:index], Filmes[index+1:]...)
			var filme Filme
			_=json.NewDecoder(r.Body).Decode(&filme)
			Filme.ID = strconv.Itoa(rand.Intn(100000000))
			Filmes = append(Filmes, Filme)
			json.NewEncoder(w).Encode(filme)
			return
		}
	}
}


func main(){
	r := mux.NewRouter()


	Filmes = append(Filmes, Filme{ID: "1", isbn: "438227", titulo : "Filme um", Diretor: &Diretor{Primeironome: "João", Sobrenome: "Silva"}})
	Filmes = append(Filmes, Filme{ID: "2", isbn: "454555", titulo : "Filme dois", Diretor: &Diretor{Primeironome: "Gabriel", Sobrenome: "Almeida"}})
	r.HandleFunc("/filmes", getFilmes).Methods("GET")
	r.HandleFunc("/filmes/{id}", getFilme).Methods("GET")
	r.HandleFunc("/filmes", criarFilme).Methods("POST")
	r.HandleFunc("/filmes/{id}", updateFilmes).Methods("PUT")
	r.HandleFunc("/filmes/{id}", deleteFilmes).Methods("DELETE")
	

	fmt.Printf("Iniciando server na porta 8000\n")
	log.Fatal(http.ListenAndServer(":8000"), r)
}