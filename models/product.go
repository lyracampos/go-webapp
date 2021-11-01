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
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)
	}
	defer db.Close()
	return products
}

func Add(name, description string, price float64, quantity int) {
	db := db.DatabaseConnect()
	query, err := db.Prepare("insert into products (name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(name, description, price, quantity)
	defer db.Close()
}

func Delete(id string) {
	db := db.DatabaseConnect()
	query, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)
	defer db.Close()
}
