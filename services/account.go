package services

import (
	"ebanx/domain"
	"errors"
	"net/http"
	"strconv"
)

func Deposit(request *domain.Request, account *domain.Account) error {

	if account.ID == 0 {
		Create(request, account)
		return nil
	}

	if strconv.Itoa(account.ID) != request.Destination {
		return errors.New("error")
	}

	account.Balance = account.Balance + request.Amount

	return nil

}

func Create(request *domain.Request, account *domain.Account) {

	destination, _ := strconv.Atoi(request.Destination)

	account.ID = destination
	account.Balance = request.Amount

	return

}

func GetBalance(request *http.Request, account *domain.Account) (*int, error) {
	id := request.FormValue("account_id")

	if account.ID == 0 || strconv.Itoa(account.ID) != id {
		return nil, errors.New("error")
	}

	return &account.Balance, nil
}

func Withdraw(request *domain.Request, account *domain.Account) error {

	if account.ID == 0 || strconv.Itoa(account.ID) != request.Origin {
		return errors.New("error")
	}

	account.Balance = account.Balance - request.Amount

	return nil
}

func Transfer(request *domain.Request, account *domain.Account) error {

	if account.ID == 0 || strconv.Itoa(account.ID) != request.Origin {
		return errors.New("error")
	}

	account.Balance = account.Balance - request.Amount

	return nil
}
