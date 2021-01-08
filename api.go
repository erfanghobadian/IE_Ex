package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/thedevsaddam/renderer"
	"net/http"
	"strconv"
)

func getCustomers(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	var decoder = schema.NewDecoder()
	decoder.Decode(&customer, r.URL.Query())
	fmt.Println(customer.Name)
	render := renderer.New()
	var customers []Customer
	DB.Where("Name LIKE ?", "%"+customer.Name+"%").Where("Address LIKE ?", "%"+customer.Address+"%").Find(&customers)
	var json_out = make(map[string]interface{})
	if customers == nil {
		json_out["msg"] = "error"
		render.JSON(w, http.StatusBadRequest, json_out)

	} else {
		json_out["size"] = len(customers)
		json_out["customers"] = customers
		json_out["msg"] = "success"
		render.JSON(w, http.StatusOK, json_out)
	}

}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	render := renderer.New()
	var customer Customer
	_ = json.NewDecoder(r.Body).Decode(&customer)
	DB.Create(&customer)
	customer_json, _ := json.Marshal(customer)
	var json_out map[string]interface{}
	json.Unmarshal(customer_json, &json_out)
	json_out["msg"] = "success"
	render.JSON(w, http.StatusCreated, json_out)
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	render := renderer.New()
	params := mux.Vars(r)
	cID := params["id"]
	var json_out = make(map[string]interface{})
	var customer Customer
	DB.Find(&customer, cID)
	if customer.ID == 0 {
		json_out["msg"] = "error"
		render.JSON(w, http.StatusBadRequest, json_out)
	} else {
		customer_json, _ := json.Marshal(customer)
		json.Unmarshal(customer_json, &json_out)
		json_out["msg"] = "success"
		render.JSON(w, http.StatusOK, json_out)
	}

}
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	render := renderer.New()
	params := mux.Vars(r)
	cID := params["id"]
	var json_out = make(map[string]interface{})
	var customer Customer
	var updatedCustomer Customer
	_ = json.NewDecoder(r.Body).Decode(&updatedCustomer)
	DB.Find(&customer, cID).UpdateColumns(updatedCustomer)
	if customer.ID == 0 {
		json_out["msg"] = "error"
		render.JSON(w, http.StatusBadRequest, json_out)
	} else {
		customer_json, _ := json.Marshal(customer)
		json.Unmarshal(customer_json, &json_out)
		json_out["msg"] = "success"
		render.JSON(w, http.StatusOK, json_out)
	}
}
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	render := renderer.New()
	params := mux.Vars(r)
	cID := params["id"]
	var json_out = make(map[string]interface{})
	var customer Customer
	DB.Find(&customer, cID)
	if customer.ID == 0 {
		json_out["msg"] = "error"
		render.JSON(w, http.StatusBadRequest, json_out)
	} else {
		DB.Delete(&customer, cID)
		json_out["msg"] = "success"
		render.JSON(w, http.StatusOK, json_out)
	}
}
func reportCustomers(w http.ResponseWriter, r *http.Request) {

	render := renderer.New()
	params := mux.Vars(r)
	month, _ := strconv.Atoi(params["month"])
	month += 1
	var customers []Customer
	DB.Raw("SELECT * FROM customers WHERE EXTRACT(MONTH FROM register_date) =  " + strconv.Itoa(month)).Scan(&customers)
	var json_out = make(map[string]interface{})
	json_out["totalCustomers"] = len(customers)
	json_out["period"] = 1
	json_out["msg"] = "success"
	render.JSON(w, http.StatusOK, json_out)
}
