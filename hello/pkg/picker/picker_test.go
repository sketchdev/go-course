package picker

import (
	"github.com/stretchr/testify/assert"
	"hello/pkg/store"
	"testing"
)

type storeMock struct {
	students []store.Student
}

func (sm storeMock) GetStudents() []store.Student {
	return sm.students
}

func TestRandom(t *testing.T) {
	// mock the store
	studentStore = storeMock{}

	t.Run("no students exist in store", func(t *testing.T) {
		students := Random(1)

		assert.Empty(t, students)
	})

	t.Run("one student in the store", func(t *testing.T) {
		studentStore = storeMock{
			students: []store.Student{
				{
					FirstName: "Sung",
					LastName:  "Kang",
				},
			},
		}
		students := Random(1)
		assert.Len(t, students, 1)
	})
}


func TestRandomTableDriven(t *testing.T) {
	tests := []struct{
		name string
		storeMock store.Store
		expectedNumberOfStudents int
		size int
	}{
		{
			name:"no students in the store",
			storeMock: storeMock{},
			expectedNumberOfStudents: 0,
			size: 1,
		},
		{
			name:"one student in the store",
			storeMock: storeMock{
				students: []store.Student{
					{
						FirstName: "Sung",
						LastName: "Kang",
					},
				},
			},
			expectedNumberOfStudents: 1,
			size: 1,
		},
		{
			name:"two students in the store",
			storeMock: storeMock{
				students: []store.Student{
					{
						FirstName: "Sung",
						LastName: "Kang",
					},
					{
						FirstName: "Paul",
						LastName: "Chang",
					},
				},
			},
			expectedNumberOfStudents: 2,
			size: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			studentStore = test.storeMock
			students := Random(test.size)
			assert.Equal(t, test.expectedNumberOfStudents, len(students))
		})
	}
}
