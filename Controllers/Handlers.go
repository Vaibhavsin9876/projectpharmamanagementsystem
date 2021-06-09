package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaibhav/practice/projectpharmamanagementsystem/Config"
	"github.com/vaibhav/practice/projectpharmamanagementsystem/Models"
)


func GetAllManagers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var manager []Models.Manager
	Config.DB.Find(&manager)
	json.NewEncoder(w).Encode(manager)
	fmt.Println("allrecord enter")
}

func GetManager(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var manager Models.Manager
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&manager)

	Config.DB.Find(&manager, params["id"])
	json.NewEncoder(w).Encode(manager)
	fmt.Println("one record enter")
}

func CreateManager(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var manager Models.Manager
	json.NewDecoder(r.Body).Decode(&manager)
	Config.DB.Create(&manager)
	json.NewEncoder(w).Encode(manager)
	fmt.Println("add someone")
}

func DeleteManager(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var manager Models.Manager
	params := mux.Vars(r)
	Config.DB.Delete(&manager, params["id"])
	json.NewEncoder(w).Encode("delete all record for the database")
	fmt.Println("delete record")
}

func GetAllPharmacists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var pharmacist []Models.Pharmacist
	Config.DB.Find(&pharmacist)
	json.NewEncoder(w).Encode(pharmacist)

}

func GetPharmacists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var pharmacist Models.Pharmacist
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&pharmacist)

	Config.DB.Find(&pharmacist, params["id"])
	json.NewEncoder(w).Encode(pharmacist)
	fmt.Println("one record enter")
}
func CreatePharmacists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var pharmacist Models.Pharmacist
	json.NewDecoder(r.Body).Decode(&pharmacist)
	Config.DB.Create(&pharmacist)
	json.NewEncoder(w).Encode(pharmacist)
	fmt.Println("add someone")
}

func DeletePharmacists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var pharmacist Models.Pharmacist
	params := mux.Vars(r)
	Config.DB.Delete(&pharmacist, params["id"])
	json.NewEncoder(w).Encode("delete all record for the database")
	fmt.Println("delete record")
}

func GetAllSalesmans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var salesman []Models.Salesman
	Config.DB.Find(&salesman)
	json.NewEncoder(w).Encode(salesman)
	fmt.Println("allrecord enter")
}

func GetSalesmans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var salesman Models.Salesman
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&salesman)

	Config.DB.Find(&salesman, params["id"])
	json.NewEncoder(w).Encode(salesman)
	fmt.Println("one record enter")
}

func CreateSalesmans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var salesman Models.Salesman
	json.NewDecoder(r.Body).Decode(&salesman)
	Config.DB.Create(&salesman)
	json.NewEncoder(w).Encode(salesman)
	fmt.Println("add someone")
}

func DeleteSalesmans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var salesman Models.Salesman
	params := mux.Vars(r)
	Config.DB.Delete(&salesman, params["id"])
	json.NewEncoder(w).Encode("delete all record for the database")
	fmt.Println("delete record")
}

func GetAllMedices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var medices []Models.Medices
	Config.DB.Find(&medices)
	json.NewEncoder(w).Encode(medices)
	fmt.Println("allrecord enter")
}

func GetMedices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var medices Models.Medices
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&medices)

	Config.DB.Find(&medices, params["id"])
	json.NewEncoder(w).Encode(medices)
	fmt.Println("one record enter")
}

func CreateMedices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var medices Models.Medices
	json.NewDecoder(r.Body).Decode(&medices)
	Config.DB.Create(&medices)
	json.NewEncoder(w).Encode(medices)
	fmt.Println("add someone")
}
func UpdateMedices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var medices Models.Medices
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&medices)

	Config.DB.First(&medices, params["id"])
	Config.DB.Save(&medices)
	json.NewEncoder(w).Encode(medices)

}

func DeleteMedices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	var medices Models.Medices
	params := mux.Vars(r)
	Config.DB.Delete(&medices, params["id"])
	json.NewEncoder(w).Encode("delete all record for the database")
	fmt.Println("delete record")
}
