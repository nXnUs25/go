package myPackage

const version = "1.0.0"

func OtherFileInSamePackage() string {
	return "Its other file location for this file"
}

func privatePackage(name string) string {
	return name + " " + name
}
