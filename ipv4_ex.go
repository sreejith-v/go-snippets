package main

import "fmt"

func incrementIP(ip [4]int) [4]int {
	octetNumber := len(ip)
	var ipIncremented = false
	for i := octetNumber - 1; i >= 0; i-- {
		if ipIncremented {
			break
		}
		if ip[i] == 255 {
			ip[i] = 0
		} else {
			ip[i] = ip[i] + 1
			ipIncremented = true
		}
	}
	return ip
}

func main() {
	ipV4 := [4]int{255, 255, 255, 255}

	//sIp := ipV4[:4]

	incrementedIP := incrementIP(ipV4)
	fmt.Println(incrementedIP)
}
