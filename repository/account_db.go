package repository

import "github.com/jmoiron/sqlx"

type AccountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return AccountRepositoryDB{db: db}
}

func (r AccountRepositoryDB) Create(acc Account) (*Account, error) {
	query := "insert into accounts (customer_id, opening_date, account_type, amount, status) values(?, ?, ?, ?, ?)"
	result, err := r.db.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AmountType,
		acc.Amount,
		acc.Status,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	acc.AccountID = int(id)

	return &acc, nil

}

func (r AccountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id=?"
	account := []Account{}
	err := r.db.Select(&account, query, customerID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
