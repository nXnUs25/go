package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // root@localhost: itRrx&AAR7EX
	"os"
	"text/template"
)

type Entry struct {
	Number  int
	Doubled int
	Square  int
}

var (
	data []Entry
	tf   string
)

func main() {
	host := "localhost"
	port := "3306"
	dbc := "root:local_pwd@tcp(" + host + ":" + port + ")/"
	dbname := "data"

	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("Template file")
		return
	}

	tf := arg[1]

	fmt.Printf("Connecting to MySQL database on [%v:%v]\n", host, port)
	db, err := sql.Open("mysql", dbc)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Connection refused verify port 3306")
		return
	}
	defer db.Close()

	fmt.Println("Creating database", dbname)
	_, err = db.Exec("create database if not exists " + dbname)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Cannot create database", dbname)
		return
	}

	fmt.Println("Setting database to use:", dbname)
	_, err = db.Exec("use " + dbname)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Cannot use database", dbname)
		return
	}

	fmt.Println("create table numbers in database name: " + dbname)
	_, err = db.Exec("create table if not exists numbers(number int not null primary key, doubled int, square int);")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Cannot create table data")
		return
	}

	fmt.Println("Empting database table")
	_, err = db.Exec("delete from numbers")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Connecting to MySQL database on [%v:%v]/%v\n", host, port, dbname)
	db, err = sql.Open("mysql", dbc+dbname)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Connection refused verify port 3306")
		return
	}
	defer db.Close()

	fmt.Println("Writing data into table")
	stm, err := db.Prepare("insert into numbers(number, doubled, square) values (?,?,?)")
	for i := 2; i < 20; i++ {
		_, err := stm.Exec(i, 2*i, i*i)
		//		fmt.Printf("exec: n:%v d:%v s:%v\n", i, 2*i, i*i)
		defer insert.Close()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Problem inserting data into table")
			return
		}
	}
	stm.Close()

	fmt.Println("Selecting data..")
	r, _ := db.Query("select * from numbers")
	defer r.Close()
	var (
		n int
		d int
		s int
	)

	for r.Next() {
		r.Scan(&n, &d, &s)
		obj := Entry{Number: n, Doubled: d, Square: s}
		//		fmt.Printf("Entry %v, %v, %v\n", obj.Number, obj.Doubled, obj.Square)
		data = append(data, obj)
	}

	//fmt.Println(data)

	t := template.Must(template.ParseGlob(tf))
	t.Execute(os.Stdout, data)

}
