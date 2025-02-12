package repository

import (
	"bookstore/internal/models"
	"database/sql"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	query := "INSERT INTO orders (user_id, total_price, status) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := r.DB.QueryRow(query, order.UserID, order.TotalPrice, order.Status).
		Scan(&order.ID, &order.CreatedAt)
	return err
}

func (r *OrderRepository) GetOrders() ([]models.Order, error) {
	rows, err := r.DB.Query("SELECT id, user_id, total_price, status, created_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderByID(orderID int) (*models.Order, error) {
	query := "SELECT id, user_id, total_price, status, created_at FROM orders WHERE id = $1"
	row := r.DB.QueryRow(query, orderID)

	var order models.Order
	err := row.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) UpdateOrderStatus(orderID int, status string) error {
	_, err := r.DB.Exec("UPDATE orders SET status = $1 WHERE id = $2", status, orderID)
	return err
}
