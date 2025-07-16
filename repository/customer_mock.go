package repository

import "errors"

type CustomerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepositoryMock {
	customers := []Customer{
		//json:data
	}
	return CustomerRepositoryMock{customers: customers}
}

func (r CustomerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r CustomerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
