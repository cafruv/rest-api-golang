package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//
type Item struct {
	Item string `json:item`
	Qt   int    `json:qt`
}

var items []Item

var dict = make(map[string]string)

func getItems(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get all items")
	json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Add item is called")

	var item Item
	var check bool
	for _, tp := range r.Header["Content-Type"] {
		if tp == "application/json" {
			check = true
		}
	}

	if check == false {
		http.Error(w, "400 BadRequest", http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&item)
	_, ok := dict[item.Item]
	//fmt.Println(ok)
	if ok {
		http.Error(w, "400 BadRequest", http.StatusBadRequest)
		return
	} else {
		s := strconv.Itoa(item.Qt)
		writeRecord(item.Item + "," + s + "\n")
		w.WriteHeader(http.StatusCreated)
		return

	}

}

func writeRecord(text string) {
	f, err := os.OpenFile("fruit-count.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text); err != nil {
		log.Println(err)
	}

}

func main() {

	file, err := os.Open("fruit-count.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		s := strings.Split(scanner.Text(), ",")
		//fmt.Println(s[0])
		if s[0] != "果物名" {
			dict[s[0]] = s[1]
		}

	}

	// リクエストを裁くルーターを作成
	router := mux.NewRouter()
	for k, v := range dict {
		i, err := strconv.Atoi(v)
		//fmt.Println(v + "[[]]  ")
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		items = append(items, Item{Item: k, Qt: i})
	}

	// エンドポイント
	router.HandleFunc("/fruit", getItems).Methods("GET")

	router.HandleFunc("/fruit", addItem).Methods("POST")

	// Start Server
	log.Println("Listen Server ....")
	// 異常があった場合、処理を停止したいため、log.Fatal で囲む
	log.Fatal(http.ListenAndServe(":8080", router))
}
