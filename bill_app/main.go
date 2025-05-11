package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func get_gst(value int) float64 {
	return 0.05 * float64(value)
}

func main() {
	var c_texts []string

	fmt.Printf("SpecialDay: ")
	sd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if sd == "Yes\n" {
		c_texts = append(c_texts, "Special Days")
	}
	fmt.Printf("Items: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	items := strings.Split(strings.Replace(text, "\n", "", 1), ",")

	fmt.Printf("Peak hour: ")
	ph, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if ph == "Yes\n" {
		c_texts = append(c_texts, "Peak hour")
	}
	fmt.Printf("Night order: ")
	no, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if no == "Yes\n" {
		c_texts = append(c_texts, "Night Charges")
	}

	food_price := get_food_price(items)
	fmt.Println("total = ", get_food_price(items))

	fmt.Println("charge = ", get_charges(c_texts))

	total_price := float64(food_price+get_charges(c_texts)) + get_gst(food_price)
	fmt.Println("Total = ", total_price)
}
