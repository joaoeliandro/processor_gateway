package factory

import "github.com/joaoeliandro/processor_gateway/domain/repositories"

type RepositoryFactory interface {
	CreateTransactionsRepository() repositories.TransactionsRepository
}
