package memory

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/victoorraphael/school-plus-BE/domain/entities"
	"github.com/victoorraphael/school-plus-BE/internal/repositories/student"
)

func TestMemory_GetStudent(t *testing.T) {
	assert := require.New(t)
	repo := New()
	std, _ := student.New(entities.Person{Name: "Raphael", Email: "raphael@email.com"})

	_ = repo.Add(std)

	type test struct {
		test        string
		id          primitive.ObjectID
		expectedErr error
	}

	testCases := []test{
		{
			test:        "get std",
			id:          std.GetID(),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			st, err := repo.Get(tc.id)
			assert.Nil(err)
			assert.Equal(st.GetName(), std.GetName())
			assert.Equal(st.GetEmail(), std.GetEmail())
		})
	}
}

func TestMemory_AddStudent(t *testing.T) {
	assert := require.New(t)
	type test struct {
		test        string
		person      entities.Person
		expectedErr error
	}

	testCases := []test{
		{
			test: "create student",
			person: entities.Person{
				Name:  "Raphael",
				Email: "raphael@emal.com",
			},
			expectedErr: nil,
		},
	}

	repo := New()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			student, err := student.New(tc.person)
			assert.Nil(err)
			errRepoAdd := repo.Add(student)
			assert.Equal(errRepoAdd, tc.expectedErr)
		})
	}
}

func BenchmarkMemory_AddStudent(b *testing.B) {
	st, _ := student.New(entities.Person{
		Name:  "Raphael",
		Email: "raphael@email.com",
	})

	testCases := []struct {
		test string
		student.Student
	}{
		{
			test:    "person1",
			Student: st,
		},
		{
			test:    "person2",
			Student: st,
		},
		{
			test:    "person3",
			Student: st,
		},
		{
			test:    "person4",
			Student: st,
		},
		{
			test:    "person5E",
			Student: st,
		},
	}

	repo := New()

	for _, tc := range testCases {
		b.Run(tc.test, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				repo.Add(tc.Student)
			}
		})
	}
}
