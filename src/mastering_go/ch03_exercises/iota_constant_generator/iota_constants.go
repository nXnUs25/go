package iota_constant_generator

type Day int

type Power4 int

const (
	P4_0 Power4 = 1 << iota
	_
	P4_1
	_
	P4_2
	_
	P4_3
	_
	P4_4
	_
	P4_5
	_
	P4_6
	_
	P4_7
	_
	P4_8
	_
	P4_9
	_
	P4_10
	_
	P4_11
	_
	P4_12
	_
	P4_13
	_
	P4_14
)

const (
	Monday Day = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	OOD Day = -1 // out of days
)

/// slice elements
func (d Day) String() string {
	return [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}[d-1]
}

/// array
func (d Day) English() string {
	return [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}[d-1]
}

func (d Day) Polish() string {
	return [7]string{"Poniedzialek", "Wtorek", "Sroda", "Czwartek", "Piatek", "Sobota", "Niedziela"}[d-1]
}

/// array pointer
func WeekDaysStar() *[7]Day {
	return &[7]Day{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}
}

/// copy of array
func WeekDays() [7]Day {
	return [7]Day{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}
}

/// slice "are passed by ref"
func WeekDaysSlice() []Day {
	return []Day{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}
}

/// slice ref
func WeekDaysSliceStar() *[]Day {
	return &[]Day{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}
}

func WeekDay(day int) Day {
	return [7]Day{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}[day]
}

// Note the [...] instead of []: it ensures you get a (fixed size) array instead of a slice. So the values aren't fixed but the size is.
func WeekDaySlice(day int) Day {
	return [...]Day{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}[day]
}

func Get(day Day) string {
	switch day {
	case Monday:
		return day.String()
	case Tuesday:
		return day.String()
	case Wednesday:
		return day.String()
	case Thursday:
		return day.String()
	case Friday:
		return day.String()
	case Saturday:
		return day.String()
	case Sunday:
		return day.String()
	default:
		return day.String()
	}
}

func GetDay(day int) Day {
	switch day {
	case 1:
		return Monday
	case 2:
		return Tuesday
	case 3:
		return Wednesday
	case 4:
		return Thursday
	case 5:
		return Friday
	case 6:
		return Saturday
	case 7:
		return Sunday
	default:
		return OOD
	}
}
