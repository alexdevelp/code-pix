package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20);not null" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	acccount := Account{
		Bank:      bank,
		Number:    number,
		OwnerName: ownerName,
	}

	acccount.ID = uuid.NewV4().String()
	acccount.CreatedAt = time.Now()

	err := acccount.isValid()

	if err != nil {
		return nil, err
	}

	return &acccount, nil
}
