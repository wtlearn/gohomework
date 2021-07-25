package main

import (
	"fmt"
	"os"
	"week02/mysql"
)

func main() {
	class := 8
	dao := mysql.NewStudentDao()
	students, err := dao.GetStudentByClass(class)
	if err != nil {
		fmt.Println("dao wrong")
		os.Exit(1)
	}
	if len(students) > 0 {
		for i, student := range students {
			fmt.Printf("seq:%d,name:%s,age:%d\n",i, student.Name, student.Age)
		}
	} else {
		fmt.Printf("no student from class %d\n", class)
	}

}

