package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pavelkres/kurs/nbrb"
)

var neededCurrencies = []string{"USD", "EUR"}

func main() {
	value, from, to, ok := getArgs()
	currencies := nbrb.Get()
	fmt.Println()

	if ok {
		result := convert(currencies, value, from, to)
		fmt.Printf("%.2f %s = %.2f %s\n", value, from, result, to)
		return
	}

	nbrb.PrintNeeded(currencies, neededCurrencies)
}

func getArgs() (float32, string, string, bool) {
	args := os.Args[1:]
	if len(args) != 3 {
		return 0, "", "", false
	}
	value, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		log.Fatalln(err)
	}
	from := strings.ToUpper(args[1])
	to := strings.ToUpper(args[2])
	return float32(value), from, to, true
}
