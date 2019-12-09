package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"httpGo/dbmethods"
)

func main() {

	//standart db actions
	/*db, err := sql.Open("sqlite3", "skateshop.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	prdct := make([]*ProductsInfo.Product, 0)

	for rows.Next() {
		p := new(ProductsInfo.Product)
		err := rows.Scan(&p.Id, &p.Item, &p.Company, &p.Price, &p.Amount)
		if err != nil {
			log.Fatal(err)
		}
		prdct = append(prdct, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, p := range prdct {
		fmt.Printf("%s %d\n %s %s\n %s %s\n %s %d\n %s %d\n\n","item id:", p.Id, "item:",  p.Item, "company:", p.Company, "price(BLR):", p.Price, "amount:", p.Amount)
	}

	rows, err = db.Query("SELECT * FROM info")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	info := make([]*ProductsInfo.Info, 0)
	for rows.Next() {
		i := new(ProductsInfo.Info)
		err := rows.Scan(&i.Id, &i.Company, &i.Information, &i.Rating)
		if err != nil {
			log.Fatal(err)
		}
		info = append(info, i)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, i := range info {
		fmt.Printf("%s %d\n %s %s\n %s %s\n %s %d\n\n","company id:", i.Id, "company:", i.Company, "information:", i.Information, "rating:", i.Rating)
	}*/

	//new table if not exists
	db, err := dbmethods.NewItemTable()
	if err != nil {
		panic(err)
	}

	//adding new products
	/*tmp := ProductsInfo.Product{}
	tmp.Item = "Wheels"
	tmp.Company = "Enjoy"
	tmp.Price = 60
	tmp.Amount = 80
	err = db.AddItem(&tmp)
	if err != nil {
		panic(err)
	}*/

	//adding new info
	/*tmp2 := ProductsInfo.Info{}
	tmp2.Rating = 10
	tmp2.Company = "Enjoy"
	tmp2.Information = "LULZ LULW"
	db.AddInfo(&tmp2)
	*/

	s1, _ := db.GetItem("Shoes", "Vans")
	fmt.Println(s1)
	s2, _ := db.GetInfo("Enjoy")
	//db.ShowAll()
	//all methods checked

	//serv
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!\n"+s2)
	})
	e.Logger.Fatal(e.Start(":1323")) // http://localhost:1323

}
