package main

import (
	"ckutils"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	var threshold int
	flag.IntVar(&threshold, "threshold", 5, "the number of days before expiration date")
	flag.IntVar(&threshold, "t", 5, "the number of days before expired date")
	var email string
    flag.StringVar(&email, "email", "sreops@acoustic.com", "<email@address>\temail address to receive alert")
	flag.StringVar(&email, "e", "sreops@acoustic.com", "<email@address> email address to receive alert")
	var path string
	flag.StringVar(&path, "directory", "/opt/IBM/RationalRLKS/config", "path to the license file")
	flag.StringVar(&path, "d", "/opt/IBM/RationalRLKS/config", "path to the license file")
	var fname string
	flag.StringVar(&fname, "file", "server_license.lic", "file name of the license")
	flag.StringVar(&fname, "f", "server_license.lic", "file name of the license")
	var smtpServer string
	flag.StringVar(&smtpServer, "server", "mailproxy.atlis1", "SMTP server name (hostname)")
	flag.StringVar(&smtpServer, "s", "mailproxy.atlis1", "SMTP server name (hostname)")
	var smtpPort int
	flag.IntVar(&smtpPort, "port", 25, "port number of smtp server")
	flag.IntVar(&smtpPort, "p", 25, "port number of smtp server")
	var data string
	flag.StringVar(&data, "message", "UCD Server License is going to expire in few days", "Email Subject")
	flag.StringVar(&data, "m", "UCD Server License is going to expire in few days", "Email Subject")
	var help bool
	flag.BoolVar(&help, "help", false, "display the usage info")
	flag.BoolVar(&help, "h", false, "display the usage info")
	var quiet bool
	flag.BoolVar(&quiet, "quiet", false, "when flag is used, script will only print messages to stdout")
	flag.BoolVar(&quiet, "q", false, "when flag is used, script will only print messages to stdout")

	flag.Parse()

	if help {
		Usage()
	}

	if quiet {
		ckutils.MainDebug(threshold, path, fname)
	} else {
		ckutils.Main(threshold, smtpPort, email, path, fname, smtpServer, data)
	}

}

func Usage(){
	filename:=filepath.Base(os.Args[0])
	fmt.Printf("Usage: \n\t%v [-h|-help] [-e|-email address] [-t|-threshold num] [-d|-path path] [-f|-file filename] [-s|-server host] [-p|-port port] [-m|-message text] [-q|quiet]\n\n", filename)
	flag.PrintDefaults()
	fmt.Printf("\n\nThe script will check the license for UCD Server, and if it is below the threshold\nMail will be send for the notification\nThere are going to be 2 alerts.\n  - 1st at X days equal threshold\n  - 2nd at X - 2 days\n")
	os.Exit(0)
}