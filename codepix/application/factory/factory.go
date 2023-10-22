package factory

import (
	"github.com/alexdevelp/code-pix/application/usecase"
	"github.com/alexdevelp/code-pix/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(databese *gorm.DB) usecase.TransctionUseCase {
	pixRepositry := repository.PixKeyRepositoryDb{Db: databese}
	trasactionRepository := repository.TransactionRepositoryDb{Db: databese}

	transactionUsecase := usecase.TransctionUseCase{
		TransctionRepository: trasactionRepository,
		PixRepository:        pixRepositry,
	}

	return transactionUsecase
}
