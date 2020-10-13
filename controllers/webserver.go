package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// ItemParams はjson型で取得
type ItemParams struct {
	ID string `json:"id"`
	ItemName string `json:"item_name,omitempty"`
	Price int `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//ポインタ型でitemsを定義．今回はグローバル変数[配列]がDBの役割をする
var items []*ItemParams

func rootPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

func fetchAllItems(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func fetchSingleItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, item := range items {
		if item.ID == key {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var item ItemParams
    if err := json.Unmarshal(reqBody, &item); err != nil {
        log.Fatal(err)
    }

    items = append(items, &item)
    json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, item := range items {
        if item.ID == id {
            items = append(items[:index], items[index+1:]...)
        }
    }
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    reqBody, _ := ioutil.ReadAll(r.Body)
    var updateItem ItemParams
    if err := json.Unmarshal(reqBody, &updateItem); err != nil {
        log.Fatal(err)
    }

    for index, item := range items {
        if item.ID == id {
            items[index] = &ItemParams{
                ID:           item.ID,
                ItemName:     updateItem.ItemName,
                Price:        updateItem.Price,
                CreatedAt:    item.CreatedAt,
                UpdatedAt:    updateItem.UpdatedAt,
            }
        }
    }
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

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

//モックデータを初期値として読み込む
func init() {
	items = []*ItemParams{
		&ItemParams{
			ID: "1",
			ItemName: "item_1",
			Price: 2500,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		&ItemParams{
			ID: "2",
			ItemName: "item_2",
			Price: 3000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		&ItemParams{
			ID: "3",
			ItemName: "item_3",
			Price: 3400,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}