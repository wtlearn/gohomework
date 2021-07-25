package main

import (
	"fmt"
	"os"
	"week02/mysql"
)

func main() {
	dao := mysql.NewStudentDao()
	students, err := dao.GetStudentByClass(8)
	if err != nil {
		fmt.Println("dao wrong")
		os.Exit(1)
	}
	for i, student := range students {
		fmt.Printf("seq:%d,name:%s,age:%d\n",i, student.Name, student.Age)
	}
}

