package usecase

import (
	"CleanArch/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (lo *ListOrdersUseCase) Execute() ([]entity.Order, error) {
	orders, err := lo.OrderRepository.ListOrders()

	if err != nil {
		return nil, err
	}

	return orders, nil
}
