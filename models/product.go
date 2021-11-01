package models

import (
	"go-webapp/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAll() []Product {
	db := db.DatabaseConnect()
	query, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for query.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = query.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)
	}
	defer db.Close()
	return products
}
