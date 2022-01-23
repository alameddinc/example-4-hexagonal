package inMemo

import (
	"errors"
	"hexagonal/internal/core/ports"
	"sync"
)

type UserInMemo struct {
	memory map[int]map[string]interface{}
	mutex  *sync.Mutex
}

func NewUserInMemo() ports.UserRepository {
	return &UserInMemo{
		memory: make(map[int]map[string]interface{}),
		mutex:  new(sync.Mutex),
	}
}

func (u *UserInMemo) Get(id int) (map[string]interface{}, error) {
	if value, ok := u.memory[id]; ok {
		return value, nil
	}
	return nil, errors.New("handler not found")
}
func (u *UserInMemo) GetAll() ([]map[string]interface{}, error) {
	tmpUserList := []map[string]interface{}{}
	for _, v := range u.memory {
		tmpUserList = append(tmpUserList, v)
	}
	return tmpUserList, nil
}
func (u *UserInMemo) Insert(id int, name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	if _, ok := u.memory[id]; ok {
		return errors.New("this id is exist")
	}
	u.mutex.Lock()
	u.memory[id] = map[string]interface{}{"id": id, "name": name}
	u.mutex.Unlock()
	return nil
}
func (u *UserInMemo) Delete(id int) error {
	if _, ok := u.memory[id]; !ok {
		return errors.New("this id is not exist")
	}
	u.mutex.Lock()
	delete(u.memory, id)
	u.mutex.Unlock()
	return nil
}
func (u *UserInMemo) UpdateName(id int, name string) error {
	if _, ok := u.memory[id]; !ok {
		return errors.New("this id is not exist")
	}
	if name == "" {
		return errors.New("name is empty")
	}
	u.mutex.Lock()
	u.memory[id] = map[string]interface{}{"id": id, "name": name}
	u.mutex.Unlock()
	return nil
}
