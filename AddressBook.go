package main

import (
	"bufio"
	"fmt"
	"os"
)

type Friend struct {
	Name         string
	MobileNumber string
	AltMobile    string
	Address      string
	City         string
}

func main() {
	friends := []Friend{
		{
			Name:         "John Smith",
			MobileNumber: "123-456-7890",
			AltMobile:    "",
			Address:      "123 Main St",
			City:         "New York",
		},
		{
			Name:         "Jane Doe",
			MobileNumber: "456-789-0123",
			AltMobile:    "789-012-3456",
			Address:      "456 Park Ave",
			City:         "Chicago",
		},
		{
			Name:         "Bob Johnson",
			MobileNumber: "789-012-3456",
			AltMobile:    "",
			Address:      "789 Maple St",
			City:         "Los Angeles",
		},
		{
			Name:         "Charlie",
			MobileNumber: "123-456-7892",
			AltMobile:    "111-222-3333",
			Address:      "789 Park Avenue",
			City:         "Los Angeles",
		},
		{
			Name:         "Tony",
			MobileNumber: "143-456-4324",
			Address:      "Chicago",
			City:         "California",
		},
		{
			Name:         "Alice",
			MobileNumber: "123-456-7892",
			AltMobile:    "111-666-3333",
			Address:      "789 Park Avenue",
			City:         "Los Angeles",
		},
		{
			Name:         "Charlie",
			MobileNumber: "123-456-7892",
			AltMobile:    "111-222-3333",
			Address:      "789 Park Avenue",
			City:         "Los Angeles",
		}, {
			Name:         "Roman",
			MobileNumber: "123-556-6892",
			Address:      "789 Park Avenue",
			City:         "Los Angeles",
		},
	}

	file, err := os.Create("address_book.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, f := range friends {
		fmt.Fprintf(w, "Name: %s\nMobile Number: %s\nAlternate Mobile Number: %s\nAddress: %s\nCity: %s\n\n", f.Name, f.MobileNumber, f.AltMobile, f.Address, f.City)
	}

	w.Flush()
}
