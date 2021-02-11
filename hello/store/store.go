package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Student struct {
	FirstName string
	LastName  string
}

type Store interface {
	GetStudents() []Student
}

type fileStore struct {}

func (f fileStore) GetStudents() []Student {
	b, err := ioutil.ReadFile("roster.json")
	// ^--- having 2 return types like this is a common pattern
	//      - if error, then returned value is nil
	//      - if no error, then err is nil

	if err != nil {
		fmt.Println("Could not open the roster: ", err)
		return nil
	}

	fmt.Println("err = ", err) // we should nil here
	fmt.Println("👋 Hello, World 🔥")

	var students []Student
	//students := make([]Student, 0) // this is functionally equivalent as the line above

	jsonErr := json.Unmarshal(b, &students)
	//                             ^--- passing a pointer

	if jsonErr != nil {
		fmt.Println("Error unmarshalling: ", jsonErr)
		return nil
	}
	return students
}

func NewStore() Store {
	return fileStore{}
}