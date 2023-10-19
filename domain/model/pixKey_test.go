package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/alexdevelp/code-pix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)
	require.NoError(t, err) // Verifique se não há erro na criação do banco

	accountNumber := "123456" // Use um número de conta válido, não "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(bank, accountNumber, ownerName)
	require.NoError(t, err) // Verifique se não há erro na criação da conta

	kind := "email"
	key := "j@j.com"
	pixKey, err := model.NewPixKey(kind, key, account)
	require.NoError(t, err) // Verifique se não há erro na criação da chave Pix

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	key = "12345678901" // Use um CPF válido como chave
	_, err = model.NewPixKey(kind, key, account)
	require.NoError(t, err) // Verifique se não há erro na criação da chave Pix com tipo "cpf"

	kind = "nome"
	_, err = model.NewPixKey(kind, key, account)
	require.Error(t, err) // Verifique se ocorre um erro na criação da chave Pix com tipo inválido
}
