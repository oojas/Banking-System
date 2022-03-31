package domain

type Customer struct {
	ID      int
	Name    string
	Age     int
	Zipcode string
	Status  string
	City    string
}

type CustomerRepository interface { // Secondary Port which implements an interface.
	FindAll() ([]Customer, error)
}
