package main // main is a special main, not-importable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// ^--- WebStorm/IntelliJ will automatically find imports and add them

type Student struct {
	FirstName string
	LastName  string
}

func main() { // entry point

	b, err := ioutil.ReadFile("roster.json")
	// ^--- having 2 return types like this is a common pattern
	//      - if error, then returned value is nil
	//      - if no error, then err is nil

	if err != nil {
		fmt.Println("Could not open the roster: ", err)
		return
	}

	fmt.Println("err = ", err) // we should nil here
	fmt.Println("👋 Hello, World 🔥")

	var students []Student
	//students := make([]Student, 0) // this is functionally equivalent as the line above

	jsonErr := json.Unmarshal(b, &students)
	//                             ^--- passing a pointer

	if jsonErr != nil {
		fmt.Println("Error unmarshalling: ", jsonErr)
		return
	}

	//for _, student := range students {
	//	fmt.Printf("%v %v\n", student.FirstName, student.LastName)
	//}

	rand.Seed(time.Now().UnixNano())

	visited := make(map[Student]bool)

	//for i := 0; i < 5; {
	//	idx := rand.Intn(len(students))
	//	student := students[idx]
	//
	//	if _, studentHasBeenSeen := visited[student]; !studentHasBeenSeen {
	//		fmt.Printf("%v %v\n", student.FirstName, student.LastName)
	//		visited[student] = true
	//		i++
	//	}
	//}

	// this is functionally the same as ^^^
	for {
		idx := rand.Intn(len(students))
		//           ^---- get a random, non-negative integer
		student := students[idx]

		if _, studentHasBeenSeen := visited[student]; !studentHasBeenSeen {
		// ^--- blank identifier (google it)

			//fmt.Printf("%v %v\n", student.FirstName, student.LastName)
			visited[student] = true
		}

		if len(visited) == 5 {
			break
		}
	}

	for student, _ := range visited {
		fmt.Printf("%v %v\n", student.FirstName, student.LastName)
		//          ^---- formatting (https://golang.org/pkg/fmt/)
	}
}
