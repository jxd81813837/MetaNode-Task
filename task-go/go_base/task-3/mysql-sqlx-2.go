package task

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db := initDbSqlx()
	books := []Book{}

	var err error
	err = db.Select(&books, "SELECT * FROM ajxd_employee WHERE Price > ?", 50)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)

}

type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}
