package main

import ("fmt" 
"bufio"
"os"
 "strings")

 func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
 }

func createBill () bill {
reader := bufio.NewReader(os.Stdin)

// fmt.Print("Create a new bill name: ")
// name, _ := reader.ReadString('\n')
// name = strings.TrimSpace(name)
name, _ := getInput("Create a new bill name: ", reader)

b:= newBill(name)
fmt.Println("Created the bill - ", b.name)

return b
}

func main () {

mybill :=createBill()
fmt.Println(mybill)


/*	myBill := newBill("mario's bill")
	myBill.format()

	myBill.addItem("soup", 4.99)
	myBill.addItem("coffee", 2.99)
	myBill.addItem("pizza", 6.99)
	myBill.updateTip(5.99)
	fmt.Println(myBill.format())*/
}