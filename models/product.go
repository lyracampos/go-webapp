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
	query, err := db.Query("select * from products order by id asc")
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

func Get(id string) Product {
	db := db.DatabaseConnect()
	query, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	for query.Next() {

		var id, quantity int
		var name, description string
		var price float64

		err = query.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}
	defer db.Close()
	return product
}

func Edit(id int, name, description string, price float64, quantity int) {
	db := db.DatabaseConnect()
	query, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(name, description, price, quantity, id)
	defer db.Close()
}
