package picker

import (
	"hello/pkg/store"
	"math/rand"
	"time"
)

func Random(size int) (randomStudents []store.Student) {

	studentStore := store.NewStore()
	students := studentStore.GetStudents()

	rand.Seed(time.Now().UnixNano())

	visited := make(map[store.Student]bool)

	for {
		idx := rand.Intn(len(students))
		student := students[idx]

		if _, studentHasBeenSeen := visited[student]; !studentHasBeenSeen {
			visited[student] = true
		}

		if len(visited) == size {
			break
		}
	}

	for student, _ := range visited {
		randomStudents = append(randomStudents, student)
	}

	return
}
