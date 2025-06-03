package task

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func main_sqlx1() {
	db := initDbSqlx()
	emps := []Employee{}

	emp := Employee{}
	var err error
	err = db.Select(&emps, "SELECT * FROM ajxd_employee WHERE department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emps)

	err = db.Get(&emp, "SELECT * FROM ajxd_employee   WHERE salary = (SELECT MAX(salary) FROM ajxd_employee) LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emp)

}

func initDbSqlx() *sqlx.DB {
	dsn := "admin:yixin_admin@tcp(10.10.201.89:3306)/dsc?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float32 `db:"salary"`
}
