package main

import (
	"encoding/json"
	"fmt"

	broker "github.com/TheGoatedDev/MetaTraderClientSDK/pkg"
)

func main() {

	broker := broker.NewBroker()

	companies, err := broker.SearchMT4("Robo")
	if err != nil {
		panic(err)
	}

	fmt.Println("Companies MT4", companies)

	companies5, err := broker.SearchMT5("RoboMarkets")
	if err != nil {
		panic(err)
	}

	companiesJSON, err := json.Marshal(companies5)
	if err != nil {
		panic(err)
	}
	fmt.Println("Companies MT5", string(companiesJSON))
}
