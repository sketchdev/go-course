package main // main is a special main, not-importable

import (
	"fmt"
	"hello/store"
	"math/rand"
	"time"
)

// ^--- WebStorm/IntelliJ will automatically find imports and add them


func main() { // entry point

	studentStore := store.NewStore()
	students := studentStore.GetStudents()

	//for _, student := range students {
	//	fmt.Printf("%v %v\n", student.FirstName, student.LastName)
	//}

	rand.Seed(time.Now().UnixNano())

	visited := make(map[store.Student]bool)

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
