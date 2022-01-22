package inMemo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		id   int
		name string
	}

	testCases := []struct {
		title string
		args  args
		want  error
	}{
		{"Successfull Test", args{id: 1, name: "Cengiz"}, nil},
		{"Successfull Test 2", args{id: 2, name: "Alameddin"}, nil},
		{"Fail Test for Blank Title", args{id: 3, name: ""}, errors.New("name is empty")},
		{"Fail Test for Duplicated ID", args{id: 1, name: "Deniz"}, errors.New("this id is exist")},
	}
	mockInMemory := NewUserInMemo()
	for _, tt := range testCases {
		t.Run(tt.title, func(t *testing.T) {
			err := mockInMemory.Insert(tt.args.id, tt.args.name)
			assert.Equal(t, tt.want, err)
		})
	}
}
