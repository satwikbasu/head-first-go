package magazine

type Subscriber struct { // Subscriber is capitalized so it can be exported
	Name string // even field names need to be capitalized
	Rate float64 // so that they can be exported
	Active bool // and used in other programs like chapter7/subscribers.go
}

type Employee struct {
	Name string
	Salary float64
	HomeAddress Address // struct as a field of a struct
}

type Address struct {
	// empty struct which can contain whatever fields required
}