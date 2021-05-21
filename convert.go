package main

import "github.com/pavelkres/kurs/nbrb"

func convert(currencies *[]nbrb.Currency, value float32,
	from string, to string) float32 {
	var currFrom, currTo nbrb.Currency
	for _, v := range *currencies {
		if from == v.Name {
			currFrom = v
		} else if to == v.Name {
			currTo = v
		}
	}
	if from == "BYN" {
		return currTo.FromBYN(value)
	} else if to == "BYN" {
		return currFrom.ToBYN(value)
	} else {
		return currTo.FromBYN(currFrom.ToBYN(value))
	}
}
