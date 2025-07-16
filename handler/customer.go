package handler

import (
	"basic/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	cusSrv service.CustomerService
}

func NewCustomerHandler(cusSrv service.CustomerService) customerHandler {
	return customerHandler{cusSrv: cusSrv}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customer, err := h.cusSrv.GetCustomers()
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.cusSrv.GetCustomer(customerID)
	if err != nil {
		handleError(w, err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
