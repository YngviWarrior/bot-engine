package utils

import (
	"log"
	"strconv"
)

func ParseFloat(s string) (f float64) {
	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		log.Panicln("PF 01: ", err)
	}

	return
}

func ParseInt(s string) (f uint64) {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		log.Panicln("PI 01: ", err)
	}

	return uint64(i)
}
