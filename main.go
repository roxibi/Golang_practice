package main

import "fmt"

func main () {
	myBill := newBill("mario's bill")
	myBill.format()
	fmt.Println(myBill.format())
}