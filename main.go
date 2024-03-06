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


var mdbClient *mongo.Client

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


    // mongo //

    var err error
    ctxBg := context.Background()
    connStr := "mongodb+srv://fernandomdelacuevamongo:Mongo1234@notekeeper.ojcenvj.mongodb.net/?retryWrites=true&w=majority&appName=NoteKeeper"

    mdbClient, err = mongo.Connect(ctxBg, options.Client().ApplyURI(connStr))
    if err != nil {
        log.Fatal(err)
    }
    if mdbClient == nil{
        log.Fatal("mdbClient es nil")
    }

    // will  run on exit
    defer func() {
        if err = mdbClient.Disconnect(ctxBg); err != nil {
            panic(err)
        }
    }()

    // enable the server, capturing error
    log.Fatal(http.ListenAndServe(serverAddr, nil))
}

func createNote(writer http.ResponseWriter, req *http.Request) {
    var note Note

    decoder := json.NewDecoder(req.Body)
    if err := decoder.Decode(&note); err != nil {
        http.Error(writer, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(writer, "Note: %+v", note)

    // mongo //
    notesCollection := mdbClient.Database("NoteKeeper").Collection("Notes")

    result, err := notesCollection.InsertOne(req.Context(), note)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusBadRequest)
        return
    }
    log.Printf("Id: %v", result.InsertedID)
}
