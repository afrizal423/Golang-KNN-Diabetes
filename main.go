package main

import (
	"fmt"
	"fp_datmin_03_19_92/core"
)

func main() {
	asd := core.Buka_file("diabetes.csv")
	fmt.Println(asd)
}
