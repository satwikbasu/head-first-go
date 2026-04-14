package magazine

type Subscriber struct { // Subscriber is capitalized so it can be exported
	Name string // even field names need to be capitalized
	Rate float64 // so that they can be exported
	Active bool // and used in other programs like chapter7/subscribers.go
	Address // anonymous struct with ommitted fields
}

type Employee struct {
	Name string
	Salary float64
	Address //anonymous struct with ommitted fields so it becomes easier to access
}

type Address struct {
    Street string
    City string
    State string
    PostalCode string
}