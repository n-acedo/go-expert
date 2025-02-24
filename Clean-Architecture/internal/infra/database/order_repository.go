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

	//////
	// rows, err := r.Db.Query("SELECT * FROM orders")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// orders := []entity.Order{}
	// for rows.Next() {
	// 	var id string
	// 	var price, tax, final_price float64
	// 	if err := rows.Scan(&id, &price, &tax, &final_price); err != nil {
	// 		return nil, err
	// 	}
	// 	orders = append(orders, entity.Order{ID: id, Price: price, Tax: tax, FinalPrice: final_price})
	// }
	// return orders, nil
}
