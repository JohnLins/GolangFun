package main

import (
   	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"io"
)

type Data struct {
    Name string  `json:"Name"`
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	queryDataSteamID, ok := r.URL.Query()["requestDataSteamID"]
	queryDataName, ok := r.URL.Query()["requestDataName"]
    
    if !ok || len(queryDataSteamID[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

	requestDataSteamID := queryDataSteamID[0]
	requestDataName := queryDataName[0]

	log.Println("The Key Thing, lets output that real fast: " + string(requestDataSteamID))
	log.Println("The Key Thing, lets output that real fast: " + string(requestDataName))


	data := &Data{/*SteamID: requestDataSteamID,*/ Name: requestDataName}
	e, err := json.Marshal(data)
	if err != nil {
        fmt.Println(err)
        return
    }

	fmt.Fprintln(w, string(e))

	WriteToFile(requestDataSteamID+".json", requestDataName)
	
}



// GO HERE     http://localhost:8080/?requestDataSteamID=6345642654&requestDataName=Will


func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
  
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
  }
