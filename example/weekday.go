package example

//go:generate textenums -type=WeekDay

type WeekDay int

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (d WeekDay) String() string {
	switch d {
	case Monday:
		return "Dilluns"
	case Tuesday:
		return "Dimarts"
	case Wednesday:
		return "Dimecres"
	case Thursday:
		return "Dijous"
	case Friday:
		return "Divendres"
	case Saturday:
		return "Dissabte"
	case Sunday:
		return "Diumenge"
	default:
		return "invalid WeekDay"
	}
}
