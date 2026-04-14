package main

import (
	"fmt"
	"head-first-go/src/magazine"
)

func main() {
	// var address magazine.Address // address variable stores all fields under type Address
	// address.Street = "123 Oak St"
	// address.City = "Omaha"
	// address.State = "NE"
	// address.PostalCode = "681069"
	// fmt.Printf("%#v\n %T\n", address, address) // type is Address and underlying type is a struct

	var employee magazine.Employee // employee variable stores all fields under type Employee 
	employee.Name = "John Doe"
	employee.Salary = 690000
	employee.Address.Street = "123 Oak St"
	employee.Address.City = "Omaha"
	employee.Address.State = "NE"
	employee.Address.PostalCode = "681069"
	// employee.Address = address // employee also stores a struct as a field i.e. HomeAddress field of type Addres
	fmt.Printf("%#v\n %T\n", employee, employee) // type is Employee and underlying type is a struct (with a nested struct)

	// two consecutive dot operators are used to access fields of a struct within a struct
	// embedded struct fields can be accessed using the outer struct
	// employee.Address.PostalCode becomes employee.PostalCode in case of anonymous structs
	fmt.Printf("Postal Code of %v = %v\n", employee.Name, employee.PostalCode)
}