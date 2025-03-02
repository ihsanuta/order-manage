package usecase

import (
	"context"
	"fmt"
	"order-manage/models"
	"order-manage/repository"
)

type OrderUsecase interface {
	SaveOrder(ctx context.Context, order *models.Order) (err error)
	UpdateOrder(ctx context.Context, order *models.Order) (err error)
	GetOrderByID(ctx context.Context, id uint64) (order models.Order, err error)
	GetOrderList(ctx context.Context, param models.ParamOrder) (response models.ListOrder, err error)
}

type orderUsecase struct {
	customerRepository repository.CustomerRepository
	orderRepository    repository.OrderRepository
}

func NewOrderUsecase(orderRepository repository.OrderRepository, customerRepository repository.CustomerRepository) OrderUsecase {
	return &orderUsecase{
		customerRepository: customerRepository,
		orderRepository:    orderRepository,
	}
}

func (c *orderUsecase) SaveOrder(ctx context.Context, order *models.Order) (err error) {
	customer, err := c.customerRepository.GetCustomerByID(ctx, order.CustomerID)
	if err != nil {
		return err
	}

	if customer.ID == 0 {
		return fmt.Errorf("customer not found")
	}

	err = c.orderRepository.SaveOrder(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (c *orderUsecase) UpdateOrder(ctx context.Context, order *models.Order) (err error) {
	err = c.orderRepository.UpdateOrder(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (c *orderUsecase) GetOrderByID(ctx context.Context, id uint64) (order models.Order, err error) {
	order, err = c.orderRepository.GetOrderByID(ctx, id)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (c *orderUsecase) GetOrderList(ctx context.Context, param models.ParamOrder) (response models.ListOrder, err error) {
	response, err = c.orderRepository.GetListOrder(ctx, param)
	if err != nil {
		return response, err
	}
	return response, nil
}
