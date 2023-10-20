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

func AverageGradeOfSchool(school *School) float64 {
	total := 0.0
	for _, student := range school.Students {
		total += student.Grade
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
