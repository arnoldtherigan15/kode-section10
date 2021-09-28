package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/arnoldtherigan15/kode-section10/status"
)

func autoUpdate(status *status.Status) {
	for range time.Tick(time.Second * 1) {
		status.Status.Water = rand.Intn(99) + 1
		status.Status.Wind = rand.Intn(99) + 1

		b, err := json.Marshal(status)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile("data.json", b, 0755)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	file, _ := ioutil.ReadFile("data.json")
	var status status.Status
	json.Unmarshal([]byte(file), &status)

	go autoUpdate(&status)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var path = path.Join("view", "index.html")
		var client, err = template.ParseFiles(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = client.Execute(w, status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening on PORT 8080")
	http.ListenAndServe(":8080", nil)
}
