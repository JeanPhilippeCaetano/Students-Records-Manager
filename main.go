package main

import (
	"fmt"
	"log"
	"studentsrecordmanager/studentsrecordmanager"
)

func main() {
	school, err := studentsrecordmanager.LoadSchool("students.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(studentsrecordmanager.AverageGradeOfSchool(school))
	fmt.Println(studentsrecordmanager.BestStudentOfSchool(school.Students))
	studentsrecordmanager.SortStudentsByDescendingGrade(school.Students)
	for _, student := range school.Students {
		fmt.Println(student)
	}
	studentsrecordmanager.SaveSchool("currentstudents.csv", school)
}
