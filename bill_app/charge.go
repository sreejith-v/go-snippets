package main

type rule struct {
	name   string
	charge int
}

type charges []rule

func get_charges(rules []string) int {
	hour_charges := charges{
		rule{name: "Peak hour", charge: 30},
		rule{name: "Special Days", charge: 50},
		rule{name: "Night Charges", charge: 20},
		rule{name: "Normal Delivery charge", charge: 20},
	}
	charges := 0
	for _, c := range rules {
		for _, r := range hour_charges {
			if r.name == c {
				charges += r.charge
			}
		}
	}

	if charges==0 {
		charges=hour_charges[3].charge
	}

	return charges
}
