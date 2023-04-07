package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

//make new bill based on existing struct and giving it initial values
func newBill(name string) bill {
b := bill {
	name: name,
	items: map[string]float64{"pie": 5.99, "cake": 3.99},
	tip:0,
}
return b
}

//format bill - say what struct this applies to, name of func, params and return type
func (b bill) format() string{
fs :="Bill breakdown: \n"
var total float64 =0

for key, value := range b.items {
fs += fmt.Sprintf("%-25v ...$%v \n", key +":",value)
total +=value }

fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

return fs
}

