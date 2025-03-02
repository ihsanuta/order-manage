package repository

import (
	"context"
	"order-manage/models"
	"order-manage/utils/pagination"

	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, order *models.Order) (err error)
	UpdateOrder(ctx context.Context, order *models.Order) (err error)
	GetOrderByID(ctx context.Context, id uint64) (order models.Order, err error)
	GetListOrder(ctx context.Context, param models.ParamOrder) (response models.ListOrder, err error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (c *orderRepository) SaveOrder(ctx context.Context, order *models.Order) (err error) {
	err = c.db.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *orderRepository) UpdateOrder(ctx context.Context, order *models.Order) (err error) {
	err = c.db.Save(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *orderRepository) GetOrderByID(ctx context.Context, id uint64) (order models.Order, err error) {
	err = c.db.Model(order).Where("id = ?", id).Take(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (c *orderRepository) GetListOrder(ctx context.Context, param models.ParamOrder) (response models.ListOrder, err error) {
	var order []models.Order

	c.db.Scopes(pagination.Paginate(order, &response.Pagination, c.db)).Find(&order)
	response.Data = order

	return response, nil
}
