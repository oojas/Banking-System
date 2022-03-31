package domain

type CustomerRepositoryStub struct { // Stub Adaptor
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) { // This will fetch us the list of the customers.
	return s.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{101, "Ojas", 22, "180012", "1", "Jammu"},
		{102, "Ritu", 49, "180012", "1", "Jammu"},
	}
	return CustomerRepositoryStub{
		customers: customers,
	}
}
