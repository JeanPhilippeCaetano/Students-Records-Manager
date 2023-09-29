package studentsrecordmanager

import (
	"sort"
)

type Student struct {
	Name  string  `csv:"Name"`
	Age   int     `csv:"Age"`
	Grade float64 `csv:"Grade"`
	ID    int     `csv:"ID"`
}

type School struct {
	Students []*Student
}

func GetAllGradesOfSchool(school *School) []float64 {
	var allGrades []float64
	for _, student := range school.Students {
		allGrades = append(allGrades, student.Grade)
	}
	return allGrades
}

func AverageGradeOfSchool(school *School) float64 {
	total := 0.0
	for _, v := range school.Students {
		total += v.Grade
	}
	return total / float64(len(school.Students))
}

func BestStudentOfSchool(students []*Student) *Student {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Grade < students[j].Grade
	})
	return students[len(students)-1]
}

func SortStudentsByDescendingGrade(students []*Student) []*Student {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Grade > students[j].Grade
	})
	return students
}
