package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"practice_go/config"
	"practice_go/models"
)

// type ItemParams struct {
// 	ID string `json:"id"`
// 	ItemName string `json:"item_name,omitempty"`
// 	Price int `json:"price,omitempty"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }
type DeleteResponse struct {
	ID string `json:"id"`
}

//ポインタ型でitemsを定義．今回はグローバル変数[配列]がDBの役割をする
var items []*ItemParams

func rootPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

func fetchAllItems(w http.ResponseWriter, r *http.Request){
	var items []models.Item
	// modelの呼び出し
	models.GetAllItems(&items)
	responseBody, err := json.Marshal(items)
	if err := nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(responseBody)
}

func fetchSingleItem(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	var item models.Item
	// modelの呼び出し
	models.GetSingleItem(&item, id)
	responseBody, err := json.Marshal(item)
	if err != nil {
		log.Datal(err)
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	
	var item models.Item
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err)
	}
	// modelの呼び出し
	models.InsertItem(&item)
	responseBody, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
	
	// modelの呼び出し
	models.DeleteItem(id)
	responseBody, err := json.Marshal(DeleteResponse{ID: id})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)

	var updateItem models.Item
	if err := json.Unmarshal(reqBody, &updateItem); err != nil {
		log.Fatal(err)
	}

	// modelの呼び出し
	models.UpdateItem(&updateItem, id)
	convertUintId, _ := strconv.ParseUint(id, 10, 64)
	updateItem.Model.ID = uint(convertUintId)
	responseBody, err := json.Marshal(updateItem)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// StartWebServer サーバーの立ち上げ
func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc({エンドポイント}, {レスポンス関数}).Methods({リクエストメソッド(複数可能)})
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/items", fetchAllItems).Methods("GET")
	router.HandleFunc("/item/{id}", fetchSingleItem).Methods("GET")

	router.HandleFunc("/item", createItem).Methods("POST")
	router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
	router.HandleFunc("/item/{id}", updateItem).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.ServerPort), router)
}