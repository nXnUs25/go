package ckutils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"net/smtp"
	"os"
	"regexp"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const LOCK_PATH = "/tmp/ucd_license_alert.lock"
const LAST_LOCK_PATH = "/tmp/ucd_license_last_alert.lock"
const LAST_ALERT int = 2

func SendEmail(smtpPort int, smtpServer, subject, msg, to string) {
	log.Printf("INFO SendEmail to %v\n", to)
	host, _ := os.Hostname()

	from := "ucd-server@" + host

	smtpHost := fmt.Sprintf("%v:%v", smtpServer, smtpPort)
	log.Printf("INFO Opening connection to SMTP Server %v", smtpHost)
	conn, err := smtp.Dial(smtpHost)
	checkError(err)

	defer conn.Close()
	defer log.Printf("INFO Closing connection to SMTP Server %v\n", smtpHost)

	err = conn.Mail(from)
	checkError(err)
	err = conn.Rcpt(to)
	checkError(err)

	wc, err := conn.Data()
	checkError(err)
	defer wc.Close()

	header_from := "From: " + from
	header_to := "To: " + to
	header_sub := "Subject: " + subject

	body := header_from + "\n" + header_to + "\n" + header_sub + "\n" + msg

	buffer := bytes.NewBufferString(body)
	if _, err := buffer.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
    log.Println("INFO SendEmail done")
}

func ExpireInDays(date string) int {
	end_time, err:= time.Parse("2-Jan-2006", date)
	checkError(err)
	today := time.Now()
	days := math.RoundToEven(end_time.Sub(today).Hours() / 24)
	log.Printf("INFO ExpireInDays %v days\n", days)
	return int(days)
}

func ReadLicence(path string) string {
	log.Printf("INFO ReadLicence path %v\n", path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	regex, err := regexp.Compile(`[0-9]{2}-[a-z]{3}-[0-9]{4}`)

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		if regex.MatchString(scanner.Text()) {
			return regex.FindString(scanner.Text())
		}

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
	}

	return ""
}

func Alert(number int, threshold int) bool {
	if number <= threshold {
		return true
	}
	return false
}

func LockAlerts(path string){
	log.Printf("INFO LockAlerts %v\n" , path)
		f, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		checkError(err)
		defer f.Close()
		t := time.Now()
		content := "Lock created at " + t.String() + "\n"
		_, err = f.Write([]byte(content))
		checkError(err)
}

func LockStillThere(path string) {
	log.Printf("INFO LockStillThere %v \n", path)
	f, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	checkError(err)
	defer f.Close()
	t := time.Now()
	content := "Lock existed and accessed at " + t.String() + "\n"
	_, err = f.Write([]byte(content))
	checkError(err)
}

func CheckIfLockExists(path string) bool {
	log.Printf("INFO CheckIfLockExists %v \n", path)
	if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func RemoveLocks(path string){
	log.Printf("INFO RemoveLocks %v \n", path)
	if err := os.Remove(path); err != nil {
		checkError(err)
	}
}

func SpamThem(number int) bool {
	if number <= 0 {
		return true
	}
	return false
}

func LastAlert(number int, threshold int) bool {
	return Alert(number, threshold - LAST_ALERT)
}

func Main(threshold, port int, email, path, fname, server, data string) {
	date := ReadLicence(path + "/" + fname)
	days := ExpireInDays(date)

	s := strconv.Itoa(days)
	msg := "Please update your license for UCD Server [" + date + "]\nBecause its expire in " + s + " days.\n\n\nRegards\nAdmin"

	if SpamThem(days) {
		log.Printf("INFO SpamThem because no %v days left\n", days)
		if CheckIfLockExists(LOCK_PATH) {
			RemoveLocks(LOCK_PATH)
		}
		if CheckIfLockExists(LAST_LOCK_PATH) {
			RemoveLocks(LAST_LOCK_PATH)
		}
	}

	if days > threshold {
		log.Printf("INFO Clear Locks because %v days are above %v threshold, License Updated!\n", days, threshold)
		if CheckIfLockExists(LOCK_PATH) {
			RemoveLocks(LOCK_PATH)
		}
		if CheckIfLockExists(LAST_LOCK_PATH) {
			RemoveLocks(LAST_LOCK_PATH)
		}
	}

	if Alert(days, threshold) && (threshold - LAST_ALERT) < days {
		log.Printf("INFO Alert because %v days left\n", days)
		if ! CheckIfLockExists(LOCK_PATH) {
			LockAlerts(LOCK_PATH)
			SendEmail(port, server, data, msg, email)
			log.Printf("INFO Lock %v set and email to %v sent because %v days left\n", LOCK_PATH, email, days)
		} else {
			LockStillThere(LOCK_PATH)
			log.Printf("INFO Lock set %v, %v days left low\n", LOCK_PATH, days)
		}
	} else if LastAlert(days, threshold) {
		log.Printf("INFO LastAlert because %v days left\n", days)
		if ! CheckIfLockExists(LAST_LOCK_PATH) {
			LockAlerts(LAST_LOCK_PATH)
			SendEmail(port, server, data, msg, email)
			log.Printf("INFO Last Lock %v set and Last email to %v sent because %v days left\n", LAST_LOCK_PATH, email, days)
		} else {
			LockStillThere(LAST_LOCK_PATH)
			log.Printf("INFO Last Lock set %v, %v days left low\n", LAST_LOCK_PATH, days)
		}
	} else {
		log.Printf("INFO All seems to be good, %v days left to expire date, which is higher than threshold %v\n", days, threshold)
	}
}

func MainDebug(threshold int, path, fname string) {
	log.Println("INFO Running check in quiet mode, Alerts will not be sent, just printed")
	date := ReadLicence(path + "/" + fname)
	days := ExpireInDays(date)

	s := strconv.Itoa(days)
	msg := "\n\tPlease update your license for UCD Server [" + date + "]\n\tBecause its expire in " + s + " days.\n\n\n\tRegards\n\tAdmin"

	if Alert(days, threshold) {
		log.Printf("INFO Alert because %v days left\n", days)
		log.Printf("INFO Alert message:\n %v\n", msg)
	} else {
		log.Printf("INFO All seems to be good, %v days left to expire date, which is higher than threshold %v\n", days, threshold)
	}

}