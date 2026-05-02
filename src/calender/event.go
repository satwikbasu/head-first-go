package calender

// Event type is created as an outer type that contains an embedded type Date
type Event struct {
	Title string
	Date // anonymous field since this type is a struct of its own
}