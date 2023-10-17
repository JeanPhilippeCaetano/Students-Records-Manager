package main

import (
	"fmt"
	"studentsrecordmanager/studentsrecordmanager"
)

func main() {
	school, _ := studentsrecordmanager.LoadSchool("students.csv")
	fmt.Println(studentsrecordmanager.AverageGradeOfSchool(school))
	fmt.Println(studentsrecordmanager.BestStudentOfSchool(school.Students))
	fmt.Println(studentsrecordmanager.SortStudentsByDescendingGrade(school.Students))
	for _, student := range school.Students {
		fmt.Println(student)
	}
	studentsrecordmanager.SaveSchool(school)
}
