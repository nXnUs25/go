package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
)

var (
	pf      = fmt.Printf
	pln     = fmt.Println
	buf     bytes.Buffer
	l       = log.New(&buf, "gen-pwd: ", log.LstdFlags|log.Lshortfile)
	I       = "[INFO] "
	D       = "[DEBUG] "
	E       = "[ERROR] "
	GetFunc = func() string {
		pc, _, _, _ := runtime.Caller(1)
		return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	}
	lf         = log.Printf
	lln        = log.Println
	utils      = "/src/mastering_go/ch04_exercises/json-xml"
	pwd, _     = os.Getwd()
	MIN        = 0
	MAX        = 94
	LENGTH int = 4
)

const PASSMAX = 5

/*
	Try to change the logic behind generatePassword.go by picking the password from a list of passwords
	found in a Go slice combined with the current system time or date.
*/
func main() {
	l.SetOutput(&buf)
	SetSeedRand(GetSeedUnixNano())
	option := rand.Intn(2) + 1
	SetOptionsToRendPWD(option)
	passwords := [PASSMAX]string{}

	for p := 0; p < PASSMAX; p++ {
		passwords[p] = GeneratePassword() + GetSalt()
	}
	plen := 8 + LENGTH + 2
	pln("Proposed passwords are: ")
	pf("[%3s]. -> %-[2]*s<-\n", "Opt", plen, "Password")
	for indexPWD, valuePWD := range passwords {
		pf("[%3v]. -> %-[2]*v<-\n", indexPWD+1, plen, valuePWD)
	}

}

func GetSeedUnix() int64 {
	seed := time.Now().Unix()
	lf(I+"Seed Unix base - Seed [%v] - func(): < %v >\n", seed, GetFunc())
	return seed
}

func GetSeedSystemTime() int64 {
	h, m, s := time.Now().Local().Clock()
	seed := int64(h + m + s)
	lf(I+"Seed System Time - Seed [%v] - h[%v]:m[%v]:s[%v] -func(): < %v >\n", seed, h, m, s, GetFunc())
	return seed
}

func GetSeedUnixNano() int64 {
	seed := time.Now().UnixNano()
	lf(I+"Seed UnixNano - Seed [%v] - func(): < %v >\n", seed, GetFunc())
	return seed
}

func SetSeedRand(seed int64) {
	lf(I+"Set Seed - [%v] - func(): < %v >\n", seed, GetFunc())
	rand.Seed(seed)
}

func SetOptionsToRendPWD(option int) {
	lf(I+"Seed Unix base - func(): < %v >\n", GetFunc())
	switch option {
	case 1:
		LENGTH = 6
		SetSeedRand(GetSeedUnix())
		lf(I+"Password Len: [%v] - func(): < %v >", LENGTH, GetFunc())
	case 2:
		LENGTH = 9
		SetSeedRand(GetSeedSystemTime())
		lf(I+"Password Len: [%v] - func(): < %v >", LENGTH, GetFunc())
	case 3:
		SetSeedRand(GetSeedUnixNano())
		lf(I+"Password Len: [%v] - func(): < %v >", LENGTH, GetFunc())
	default:
		lf(I+"Using Default Settings - Password Len: [%v] - func(): < %v >", LENGTH, GetFunc())
	}
}

func GeneratePassword() string {
	lf(I+"Password Len: [%v] - func(): < %v >\n", LENGTH, GetFunc())
	char := "!"
	password := ""
	for i := 0; i < LENGTH; i++ {
		rand := rand.Intn(MAX-MIN) + MIN
		newChar := string(char[0] + byte(rand))
		password += newChar
	}
	lf(I+"Password: [%v] - func(): < %v >\n", password, GetFunc())
	return password
}

func GetSalt() string {
	lf(I+"func(): < %v >\n", GetFunc())
	salt := ""
	day := time.Now().UTC().Day()
	month := int(time.Now().UTC().Month())
	year := time.Now().UTC().Year()
	hour := time.Now().UTC().Hour()
	minute := time.Now().UTC().Minute()
	second := time.Now().UTC().Second()
	salt = fmt.Sprintf("%v%v%v%v%v%v", day, month, year, hour, minute, second)
	lf(I+"Salt Settings: [%v][%v][%v][%v][%v][%v] - Salt [%v] - func(): < %v >\n", day, month, year, hour, minute, second, salt, LENGTH, GetFunc())
	return salt
}
