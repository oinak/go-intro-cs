package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mdbClient *mongo.mdbClient

func main() {
	fmt.Println("Hola Caracola")
	// ahora hagamos un servidor web
	const serverAddr string = "127.0.0.1:8081"

	// handler de peticiones
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tengop get")
		w.Write([]byte("HTTP Caracola"))
	})
	http.HandleFunc("POST /notes", createNote)

	// enable the server, capturing error
	log.Fatal(http.ListenAndServe(serverAddr, nil))

	var err error
	ctxBg := context.Background()
	const connStr string = "mongodb+srv://fernandomdelacuevamongo:Mongo1234!@cluster0.sutkl9e.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	mdbClient, err = mongo.Connect(ctxBg, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}
}

type Scope struct {
	Project string
	Area    string
}

type Note struct {
	Title string
	Tags  []string
	Text  string
	Scope Scope
}

func createNote(writer http.ResponseWriter, req *http.Request) {
	var note Note
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&note); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(writer, "Note: %+v", note)
}
