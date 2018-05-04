package main

import (
	"log"

	csvUtil "github.com/samkreter/goDataScience/csv"
)

func main() {
	_, err := csvUtil.ReadDataFrame("data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
}
