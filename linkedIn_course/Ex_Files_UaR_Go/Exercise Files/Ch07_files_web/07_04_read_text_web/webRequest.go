package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
	checkError(err)

	fmt.Println("Response:", resp)
	fmt.Printf("Response type %T\n", resp)
	fmt.Println("Response:", resp.Body)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))

}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}