package task

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func main_crud() {
	db := initDB()
	stu := Students{}
	stu.name = "李福"
	stu.age = 20
	stu.grade = "三年级"

	//fmt.Println(insetStudents(db, stu))
	//查询年龄大于18岁的
	student := getQueryStudentsByAge(db, 18)
	fmt.Println(student)
	//修改为四年级
	student.grade = "四年级"
	//updateStudents(db, student)
	//删除小于15岁的
	deleteStudentsByAge(db, 15)
}

func initDB() *sql.DB {
	// 配置数据库连接信息
	dsn := "admin:yixin_admin@tcp(10.10.201.89:3306)/dsc?charset=utf8mb4&parseTime=True&loc=Local"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 设置连接池配置
	db.SetMaxOpenConns(25)                 // 最大打开连接数
	db.SetMaxIdleConns(5)                  // 最大空闲连接数
	db.SetConnMaxLifetime(5 * time.Minute) // 连接最大存活时间

	// 测试连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to MySQL database!")

	return db
}

type Students struct {
	id    int `db:"id,omitempty"`
	name  string
	age   int
	grade string
}

func getQueryStudentsByAge(db *sql.DB, minAge int) Students {
	var user Students
	err := db.QueryRow("SELECT id, name, age ,grade FROM ajxd_all WHERE age> ?", minAge).
		Scan(&user.id, &user.name, &user.age, &user.grade)
	if err != nil {
		if err == sql.ErrNoRows {
			return Students{}
		}
		return Students{}
	}
	return user
}

func insetStudents(db *sql.DB, stu Students) sql.Result {
	result, err := db.Exec("INSERT INTO ajxd_all (name, age, grade) VALUES (?, ?, ?)", stu.name, stu.age, stu.grade)
	if err != nil {
		fmt.Println("插入失败》〉》〉", err)
		return result
	}
	return result
}

func updateStudents(db *sql.DB, stu Students) sql.Result {
	result, err := db.Exec("update ajxd_all set grade =? where name= ?", stu.grade, stu.name)
	if err != nil {
		fmt.Println("修改失败》〉》〉", err)
		return result
	}
	return result
}
func deleteStudentsByAge(db *sql.DB, age int) sql.Result {
	result, err := db.Exec("delete from ajxd_all  where age< ?", age)
	if err != nil {
		fmt.Println("删除失败》〉》〉", err)
		return result
	}
	return result
}
