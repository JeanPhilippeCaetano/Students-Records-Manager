package studentsrecordmanager

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func LoadSchool(filename string) (*School, error) {
	file, err := os.Open(filename)
	if err != nil {
		return &School{}, err
	}
	defer file.Close()

	school := School{}

	if err := gocsv.UnmarshalFile(file, &school.Students); err != nil {
		return &School{}, err
	}
	for _, student := range school.Students {
		fmt.Println("Hello, ", student.Name)
	}
	return &school, nil
}

func SaveSchool(school *School) error {
	file, err := os.Create("currentstudents.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	gocsv.MarshalFile(&school.Students, file)
	return nil
}
