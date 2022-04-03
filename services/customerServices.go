package services

import "Banking_System/domain"

type CustomerService interface { //Service interface. ( Primary Port) // This will handle the api calls.
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
} //This is basically used for initiating the default customer service.
