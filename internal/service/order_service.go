package service

import (
	"bookstore/internal/models"
	"bookstore/internal/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
}

func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	order.Status = "pending" // Устанавливаем статус по умолчанию
	return s.OrderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrders() ([]models.Order, error) {
	return s.OrderRepo.GetOrders()
}

func (s *OrderService) GetOrderByID(orderID int) (*models.Order, error) {
	return s.OrderRepo.GetOrderByID(orderID)
}

func (s *OrderService) UpdateOrderStatus(orderID int, status string) error {
	return s.OrderRepo.UpdateOrderStatus(orderID, status)
}
