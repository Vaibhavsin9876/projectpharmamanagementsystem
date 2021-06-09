package Routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaibhav/practice/projectpharmamanagementsystem/Controllers"
)

func MakeRouter() {
	routermux := mux.NewRouter()
	routermux.HandleFunc("/Managers", Controllers.GetAllManagers).Methods("GET")
	routermux.HandleFunc("/Managers/{id}", Controllers.GetManager).Methods("GET")
	routermux.HandleFunc("/Managers", Controllers.CreateManager).Methods("POST")
	routermux.HandleFunc("/Managers/{id}", Controllers.DeleteManager).Methods("DELETE")

	routermux.HandleFunc("/Pharmacists", Controllers.GetAllPharmacists).Methods("GET")
	routermux.HandleFunc("/Pharmacists/{id}", Controllers.GetPharmacists).Methods("GET")
	routermux.HandleFunc("/Pharmacists", Controllers.CreatePharmacists).Methods("POST")
	routermux.HandleFunc("/Pharmacists/{id}", Controllers.DeletePharmacists).Methods("DELETE")

	routermux.HandleFunc("/Salesmans",Controllers.GetAllSalesmans).Methods("GET")
	routermux.HandleFunc("/Salesmans/{id}", Controllers.GetSalesmans).Methods("GET")
	routermux.HandleFunc("/Salesmans", Controllers.CreateSalesmans).Methods("POST")
	routermux.HandleFunc("/Salesmans/{id}", Controllers.DeleteSalesmans).Methods("DELETE")

	routermux.HandleFunc("/Medices", Controllers.GetAllMedices).Methods("GET")
	routermux.HandleFunc("/Medices/{Name}", Controllers.GetMedices).Methods("GET")
	routermux.HandleFunc("/Medices", Controllers.CreateMedices).Methods("POST")
	routermux.HandleFunc("/Medices", Controllers.UpdateMedices).Methods("PUT")
	routermux.HandleFunc("/Medices/{Name}", Controllers.DeleteMedices).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080", routermux))
	fmt.Println("it complete")
}
