package main

type food struct {
	item  string
	price int
}

type menu []food

func get_food_price(items []string) int {
	food_menu := menu{
		food{item: "Pizza", price: 150},
		food{item: "Burger", price: 100},
		food{item: "Coke", price: 50},
		food{item: "Brownies", price: 60},
	}

	total := 0
	for _, it := range items {
		for _, food := range food_menu {
			if food.item == it {
				total += food.price
			}
		}
	}

	return total
}
