package repository

import (
	"context"
	"order-manage/models"
	"order-manage/utils/pagination"

	"github.com/jinzhu/gorm"
)

type CustomerRepository interface {
	SaveCustomer(ctx context.Context, customer *models.Customer) (err error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) (err error)
	GetCustomerByID(ctx context.Context, id uint64) (customer models.Customer, err error)
	GetListCustomer(ctx context.Context, param models.ParamCustomer) (response models.ListCustomer, err error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (c *customerRepository) SaveCustomer(ctx context.Context, customer *models.Customer) (err error) {
	err = c.db.Create(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) UpdateCustomer(ctx context.Context, customer *models.Customer) (err error) {
	err = c.db.Save(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) GetCustomerByID(ctx context.Context, id uint64) (customer models.Customer, err error) {
	err = c.db.Model(customer).Where("id = ?", id).Take(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (c *customerRepository) GetListCustomer(ctx context.Context, param models.ParamCustomer) (response models.ListCustomer, err error) {
	var customer []models.Customer

	c.db.Scopes(pagination.Paginate(customer, &response.Pagination, c.db)).Find(&customer)
	response.Data = customer

	return response, nil
}
