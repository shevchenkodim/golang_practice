package database

import (
	"database/sql"
	"fmt"
)

type Product struct {
	Id      int64
	Model   string
	Company string
	Price   int
}

func (p Product) Create(db *sql.DB) (id, count int64, err error) {
	result, err := db.Exec("insert into products (model, company, price) values ($1, $2, $3)",
		p.Model, p.Company, p.Price)
	if err != nil {
		panic(err)
	}

	id, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}

	count, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return id, count, err
}

func (p Product) AllRows(db *sql.DB) (qs []Product) {
	rows, err := db.Query("select * from Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Println(p.Id, p.Model, p.Company, p.Price)
	}

	return products
}

func (p Product) Update(db *sql.DB) (status bool, err error) {
	result, err := db.Exec("update Products set model = $1, company = $2, price = $3 where id = $4",
		&p.Model, &p.Company, &p.Price, &p.Id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())

	return true, err
}

func (p Product) Remove(db *sql.DB) (status bool, err error) {
	result, err := db.Exec("delete from Products where id = $1", &p.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())

	return true, nil
}

//db, err := sql.Open("sqlite3", "db.sqlite3")
//if err != nil {
//	panic(err)
//}
//defer db.Close()

//var prod1 = database.Product{Model: "m1", Company: "com1", Price: 1}
//fmt.Println(prod1.Create(db))

//products := database.Product{}.AllRows(db)
//fmt.Println(products)

//updateProd := database.Product{Id: 1, Model: "test", Company: "test", Price: 11111111}
//fmt.Println(updateProd.Update(db))

//removeRow := database.Product{Id: 1}
//fmt.Println(removeRow.Remove(db))
