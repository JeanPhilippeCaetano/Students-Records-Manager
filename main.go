package main

import (
	"fmt"
	"studentsrecordmanager/studentsrecordmanager"
)

func main() {
	school, _ := studentsrecordmanager.LoadSchool("students.csv")
	fmt.Println(school)
	// school.Students = append(school.Students, studentsrecordmanager.Student{
	// 	Name:  "JP",
	// 	Age:   12,
	// 	Grade: 17.5,
	// 	ID:    1,
	// })
	// school.Students = append(school.Students, studentsrecordmanager.Student{
	// 	Name:  "Jordan",
	// 	Age:   7,
	// 	Grade: 2.4,
	// 	ID:    2,
	// })
	// school.Students = append(school.Students, studentsrecordmanager.Student{
	// 	Name:  "Romain",
	// 	Age:   20,
	// 	Grade: 12.7,
	// 	ID:    3,
	// })
	// school.Students = append(school.Students, studentsrecordmanager.Student{
	// 	Name:  "Alexandre",
	// 	Age:   16,
	// 	Grade: 13.7,
	// 	ID:    4,
	// })
	// school.Students = append(school.Students, studentsrecordmanager.Student{
	// 	Name:  "Olivier",
	// 	Age:   29,
	// 	Grade: 2.0,
	// 	ID:    5,
	// })
	fmt.Println(studentsrecordmanager.AverageGradeOfSchool(school))
	fmt.Println(studentsrecordmanager.BestStudentOfSchool(school.Students))
	fmt.Println(studentsrecordmanager.SortStudentsByDescendingGrade(school.Students))
	studentsrecordmanager.SaveSchool(school)
}
