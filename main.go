package main

import (
	"fmt"
	"log"
)


func main() {
	client := NewClient("EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI", true, "https://api.edu.cdek.ru/v2")

	size := Size{
		Weight: 1.5,
		Length: 20,
		Width:  15,
		Height: 10,
	}

	addrFrom := "Россия, г. Москва, Cлавянский бульвар д.1"
	addrTo := "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43"

	priceSendings, err := client.Calculate(addrFrom, addrTo, size)
	if err != nil {
		log.Fatalf("failed to calculate: %v", err)
	}

	for _, sending := range priceSendings {
		fmt.Printf("Tariff: %s\n", sending.TariffName)
		fmt.Printf("Description: %s\n", sending.TariffDescription)
		fmt.Printf("Delivery sum: %.2f\n", sending.DeliverySum)
		fmt.Printf("Delivery period: %d - %d days\n\n", sending.PeriodMin, sending.PeriodMax)
	}
}
