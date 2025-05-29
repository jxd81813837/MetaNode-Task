package 面向对象

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	psn        Person
	employeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工姓名: %s\n员工年龄: %d\n员工id: %d\n", e.psn.Name, e.psn.Age, e.employeeID)

}

func maindx1() {
	e := Employee{Person{"JXD", 18}, 2131231}
	e.PrintInfo()
}
