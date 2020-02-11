package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	type Test struct {
		Hello   string `json:"hello"`
		TestVal int8   `json:"test_val"`
	}
	switch r.Method {
	case "GET":
		test := Test{Hello: "world", TestVal: 12}
		testJson, _ := json.Marshal(&test)
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(testJson); err != nil {
			panic(err)
		}
	case "POST":
		var test Test
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(data, &test); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		testJson, _ := json.Marshal(&test)
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(testJson); err != nil {
			panic(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
