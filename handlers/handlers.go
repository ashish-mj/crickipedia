package handlers

import (
	"crickipedia/db"
	"crickipedia/models"
	"encoding/json"
	"net/http"
	"strings"
	"log"
)

func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var player[] models.Player
	player,err := db.GetAllDocuments()
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(player)
	w.WriteHeader(http.StatusOK)
}

func GetPlayerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.Split(r.URL.Path, "/")[3]
	var player models.Player
	player,err := db.GetDocument(id)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(player)
	w.WriteHeader(http.StatusOK)
}

func CreatePlayer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var player models.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = db.InsertDocument(player)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(player)
	w.WriteHeader(http.StatusCreated)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id := strings.Split(r.URL.Path, "/")[3]
	err := db.DeleteDocument(id)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
    w.WriteHeader(http.StatusOK)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id := strings.Split(r.URL.Path, "/")[3]
	var player,request models.Player 
	player,err := db.GetDocument(id)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	player.Contract = request.Contract
	err = db.UpdateDocument(player)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
    w.WriteHeader(http.StatusOK)
}