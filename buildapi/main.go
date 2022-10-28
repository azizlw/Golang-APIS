package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for Inventory Management - file

type ShopItems struct {
	ItemId    string  `json:"itemid"`
	ItemName  string  `json:"itemname"`
	ItemDesc  string  `json:"desc"`
	ItemPrice float64 `json:"price"`
	// ItemQuantity int `json:"quantity"`
}

type ShopOwner struct {
	AdminName string
	AdminId   string
}

type Customer struct {
	UserName string
	UserId   string
}

// fake database
var items []ShopItems

// helper function - file
func (i *ShopItems) IsEmpty() bool {
	//return i.ItemId == "" && i.ItemName == ""
	// not checking because we will generate ourself
	// and it will not send by customer
	return i.ItemName == ""
}

func handleRequests() {
	r := mux.NewRouter()

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/items", getAllItems).Methods("GET")
	r.HandleFunc("/item/{id}", getOneItem).Methods("GET")
	r.HandleFunc("/item", createOneItem).Methods("POST")
	r.HandleFunc("/item/{id}", updateOneItem).Methods("PUT")
	r.HandleFunc("/item/{id}", deleteOneItem).Methods("POST")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	fmt.Println("Welcome to golang API")

	//seeding
	items = append(items, ShopItems{
		ItemId:    "546",
		ItemName:  "Mobile",
		ItemDesc:  "Apple Iphone 14 Pro Max",
		ItemPrice: 130000,
	})
	items = append(items, ShopItems{
		ItemId:    "873",
		ItemName:  "Laptop",
		ItemDesc:  "Apple Macbook Pro",
		ItemPrice: 230000,
	})

	handleRequests()
}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Aziz</h1>"))
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Items from Inventory")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getOneItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Item from Inventory")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// fmt.Printf("Params value is %v and Type is %T", params, params)
	//loop through items, find matching id and return the response
	for _, item := range items {
		if item.ItemId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("No Item found with given id")
	return
}

func createOneItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Item")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about - {}

	var item ShopItems
	_ = json.NewDecoder(r.Body).Decode(&item)
	if item.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, and convert into string
	// append new item in items

	rand.Seed(time.Now().UnixNano())
	item.ItemId = strconv.Itoa(rand.Intn(100))
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
	return
}

func updateOneItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Item")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	//loop, id, remove, add with my ID

	for index, item := range items {
		if item.ItemId == params["id"] {
			items = append(items[:index], items[index+1:]...)
			var item ShopItems
			_ = json.NewDecoder(r.Body).Decode(&item)
			return
		}
	}
	items = append(items)
	//send response when id is not found
	json.NewEncoder(w).Encode("No Item found with given id")
	return
}

func deleteOneItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Item")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	//loop, id, remove

	for index, item := range items {
		if item.ItemId == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	//send response when id is not found
	json.NewEncoder(w).Encode("No Item found with given id")
	return
}

// func Register() {

// }

// func Login() {

// }

// func Order() {

// }

// func BillGenerate() {

// }
