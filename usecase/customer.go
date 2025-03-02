package usecase

import (
	"context"
	"order-manage/models"
	"order-manage/repository"
)

type CustomerUsecase interface {
	SaveCustomer(ctx context.Context, customer *models.Customer) (err error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) (err error)
	GetCustomerByID(ctx context.Context, id uint64) (customer models.Customer, err error)
	GetCustomerList(ctx context.Context, param models.ParamCustomer) (response models.ListCustomer, err error)
}

type customerUsecase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUsecase(customerRepository repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepository: customerRepository,
	}
}

func (c *customerUsecase) SaveCustomer(ctx context.Context, customer *models.Customer) (err error) {
	err = c.customerRepository.SaveCustomer(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerUsecase) UpdateCustomer(ctx context.Context, customer *models.Customer) (err error) {
	err = c.customerRepository.UpdateCustomer(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerUsecase) GetCustomerByID(ctx context.Context, id uint64) (customer models.Customer, err error) {
	customer, err = c.customerRepository.GetCustomerByID(ctx, id)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (c *customerUsecase) GetCustomerList(ctx context.Context, param models.ParamCustomer) (response models.ListCustomer, err error) {
	response, err = c.customerRepository.GetListCustomer(ctx, param)
	if err != nil {
		return response, err
	}
	return response, nil
}
