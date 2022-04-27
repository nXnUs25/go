package main

import (
	//"database/sql"
	"fmt"
	gdb "github.com/feyeleanor/gosqlite3"
	"os"
	txt "text/template"
)

type Entry struct {
	Number  int
	Double  int
	Squares int
}

var (
	Data    []Entry
	tmpFile string
)

func main() {
	arg := os.Args
	if len(arg) != 3 {
		fmt.Println("Need DB file and Template file")
		fmt.Println("Ex. htmlT.db text.gotemp")
		info := `
			if there is no Sqllitter3 DB file included please run:
			> sqlite3 filename.db
			SQLite version 3.38.2 2022-03-26 13:51:10
			Enter ".help" for usage hints.
			sqlite> create table data (
				...> number integer primary key,
				...> double integer,
				...> square integer );
			sqlite>
			Ctl-D to exit and the DB file should be created and ready to be used.
		`
		fmt.Println(info)
		return
	}

	dbfile := arg[1]
	tmpFile = arg[2]

	db, err := gdb.Open(dbfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Empting database table")
	_, err = db.Execute("delete from data")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Writing database table in file", dbfile)
	stm, err := db.Prepare("insert into data(number, double, square) values(?,?,?)")
	for i := 2; i < 20; i++ {
		_, err = stm.Exec(i, i*2, i*i)
	}

	r, _ := db.runQuery("select * from data")

	var (
		n int
		d int
		s int
	)

	for r.Next() {
		r.Scan(&n, &d, &s)
		obj := Entry{n, d, s}
		Data = append(Data, obj)
	}

	t := txt.Must(txt.ParseGlob(tmpFile))
	t.Execute(os.Stdout, Data)
}
