package database

import (
	"CleanArch/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	return r.Db.Create(order).Error
}

func (r *OrderRepository) ListOrders() ([]entity.Order, error) {
	var orders []entity.Order

	r.Db.Find(&orders)

	for _, order := range orders {
		fmt.Println(order)
	}

	return orders, nil
}
