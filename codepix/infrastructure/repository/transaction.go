package repository

import (
	"fmt"

	"github.com/alexdevelp/code-pix/domain/model"
	"github.com/jinzhu/gorm"
)

// type TransactionRepositoryInterface interface {
// 	Register(transaction *Transaction) error
// 	Save(transaction *Transaction) error
// 	Find(id string) (*Transaction, error)
// }

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (t TransactionRepositoryDb) Register(transction *model.Transaction) error {
	err := t.Db.Create(transction).Error

	if err != nil {
		return err
	}
	return nil
}

func (t TransactionRepositoryDb) Save(transction *model.Transaction) error {
	err := t.Db.Save(transction).Error

	if err != nil {
		return err
	}
	return nil
}

func (t TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transction model.Transaction

	t.Db.Preload("AccountFrom.Bank").First(&transction, "id = ?", id)

	if transction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transction, nil

}
