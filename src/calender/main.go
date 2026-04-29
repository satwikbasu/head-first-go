package calender

import "errors"
	
// struct fields not capitalized so they remain unexported
type Date struct {
	year int
	month int
	day int
}

// however, setter methods remain capitalized so they can be used to set the struct fields
// -> private fields but public setters
func (d *Date) SetYear (year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth (month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay (day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
