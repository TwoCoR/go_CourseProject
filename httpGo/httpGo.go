package main

import (
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	memorycache "httpGo/cache"
	"httpGo/dbmethods"
	"net/http"
	"os"
	"strings"
	"time"
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
	tmp.Item = "Shoes"
	tmp.Company = "Vans"
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

	fmt.Println("Работать ли с cache?\n да - введите 1\n нет - введите 0")
	var check int
	fmt.Fscan(os.Stdin, &check)
	if check == 1 {
		fmt.Println("Сохраним в кеше информацию о компании с бд. Время жизни контейнера одна минута, как и нашей записи.")
		cache := memorycache.New(time.Minute, 10*time.Minute)
		str, _ := db.GetInfo("Enjoy")
		cache.Set("myKey", str, time.Minute)
		fmt.Println("Хотим вывести информацию о компании?\n да - введите 1\n нет - введите 0")
		fmt.Fscan(os.Stdin, &check)
		if check == 1 {
			i, b := cache.Get("myKey")
			fmt.Printf("%s %t", "Информация существует?", b)
			fmt.Printf("%s %s", "\nИнформация: ", i)
		} else {
			fmt.Println("Хотим удалить из кеша или подождать пока закончиться время жизни?\n Удалить - 1\n Ждать - 0")
			fmt.Fscan(os.Stdin, &check)
			if check == 0 {
				fmt.Println("Подождем минуту, чтобы отчистился кэш")
				time.Sleep(time.Minute)
				i, b := cache.Get("myKey")
				fmt.Printf("%s %t", "Информация существует?", b)
				fmt.Printf("%s %s", "\nИнформация: ", i)
				fmt.Println("\nПроверим действительно ли у объекта закончилось время жизни\nвыведем массив ключей у которых закончилось время")
				fmt.Println(cache.ExpiredKeys())
			} else if check == 1 {
				cache.Delete("myKey")
				i, b := cache.Get("myKey")
				fmt.Printf("%s %t", "Информация существует?", b)
				fmt.Printf("%s %s", "\nИнформация: ", i)
				fmt.Println("\nПроверим действительно ли удалили то что хранилось в кеше под нашим ключем или у объекта закончилось время жизни\nвыведем массив ключей у которых закончилось время")
				fmt.Println(cache.ExpiredKeys())
			}
		}
	}


	s2, _ := db.GetInfo("Enjoy")

	//db.ShowAll()
	//all methods checked

	//serv
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!\n"+s2)
	})
	e.Logger.Fatal(e.Start(":1323")) // http://localhost:1323

	handler := http.NewServeMux()

	s := &http.Server{
		Addr:           "localhost:8181",
		Handler:        handler,          // if nil use default http.DefaultServeMux
		ReadTimeout:    10 * time.Second, // max duration reading entire request
		WriteTimeout:   10 * time.Second, // max timing write response
		IdleTimeout:    15 * time.Second, // max time wait for the next request
		MaxHeaderBytes: 1 << 20,          // 2^20 or 128kbytes
	}

	go func() {
		log.Printf("Listening on http://%s\n", s.Addr)
		log.Fatal(s.ListenAndServe())
	}()
}

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		//b := len(auth) != 2
		if  auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		//i := len(pair) != 2
		if  !validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func validate(username, password string) bool {
	if username == "test" && password == "test" { //Basic dGVzdDp0ZXN0
		return true
	}
	return false
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		log.Printf("server [net/http] method [%s]  connection from [%v]", r.Method, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}