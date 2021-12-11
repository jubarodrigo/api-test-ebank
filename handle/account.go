package handle

import (
	"ebanx/domain"
	"ebanx/services"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	deposit  = "deposit"
	withdraw = "withdraw"
	transfer = "transfer"
)

var account domain.Account

func Reset(w http.ResponseWriter, r *http.Request) {

	account = domain.Account{}
	json.NewEncoder(w).Encode(0)

	return
}

func GetBalance(w http.ResponseWriter, r *http.Request) {

	result, err := services.GetBalance(r, &account)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(0)
		return
	}

	json.NewEncoder(w).Encode(&result)
}

func Account(w http.ResponseWriter, r *http.Request) {

	var request domain.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(0)
		return
	}

	if request.Type == deposit {
		Deposit(w, &request)
	}
	if request.Type == withdraw {
		Withdraw(w, &request)
	}
	if request.Type == transfer {
		Transfer(w, &request)
	}
}

func Deposit(w http.ResponseWriter, request *domain.Request) {

	services.Deposit(request, &account)

	json.NewEncoder(w).Encode(domain.Destination{&account})
}

func Withdraw(w http.ResponseWriter, request *domain.Request) {
	err := services.Withdraw(request, &account)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(0)
		return
	}

	json.NewEncoder(w).Encode(domain.Origin{&account})
}

func Transfer(w http.ResponseWriter, request *domain.Request) {

	var destination domain.Account

	err := services.Transfer(request, &account)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(0)
		return
	}

	destination.ID, _ = strconv.Atoi(request.Destination)
	destination.Balance = request.Amount

	json.NewEncoder(w).Encode(domain.Transfer{
		Origin:      &account,
		Destination: destination})
}
