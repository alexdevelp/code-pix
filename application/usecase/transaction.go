package usecase

import (
	"errors"
	"log"

	"github.com/alexdevelp/code-pix/domain/model"
)

type TransctionUseCase struct {
	PixRepository        model.PixKeyRepositoryInterface
	TransctionRepository model.TransactionRepositoryInterface
}

func (t *TransctionUseCase) Register(accountId string, amount float64, pixkeyto string, pixKeyKndTo string, description string) (*model.Transaction, error) {
	account, err := t.PixRepository.FindAccount(accountId)
	// if err != nil : nil, err
	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixRepository.FindKeyByKind(pixkeyto, pixKeyKndTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	t.TransctionRepository.Register(transaction)
	if transaction.ID == "" {
		return nil, errors.New("unable to create new transaction at the moment")
	}

	return transaction, nil
}

func (t *TransctionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransctionRepository.Find(transactionId)
	if err != nil {
		log.Println("transaction not found")
	}

	transaction.Status = model.TransactionConfirmed
	err = t.TransctionRepository.Save(transaction)
	if err != nil {
		return nil, errors.New("unable to save new transaction status at the moment")
	}

	return transaction, nil
}

func (t *TransctionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransctionRepository.Find(transactionId)
	if err != nil {
		log.Println("transaction not found")
	}

	transaction.Status = model.TransactionCompleted
	err = t.TransctionRepository.Save(transaction)
	if err != nil {
		return nil, errors.New("unable to save new transaction status at the moment")
	}

	return transaction, nil
}

func (t *TransctionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := t.TransctionRepository.Find(transactionId)
	if err != nil {
		log.Println("transaction not found")
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason
	err = t.TransctionRepository.Save(transaction)
	if err != nil {
		return nil, errors.New("unable to save new transaction status at the moment")
	}

	return transaction, nil
}
