package magazine

type Subscriber struct {
	Name string
	Rate float64
	Active bool
}

type Employee struct {
	Name string
	Salary float64
	HomeAddress Address // struct as a field of a struct
}

type Address struct {
	// empty struct which can contain whatever fields required
}