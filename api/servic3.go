package main

import "slices"

type UserInterface interface {
	Create(u *User) error
	Update(u *User) error
	Delete(id int) error
	Find(id int) (*User, error)
	List() ([]*User, error)
}

type userService struct {
	users []*User
}

func NewUserService() UserInterface {
	return &userService{
		users: make([]*User, 0),
	}
}

func (s *userService) Create(u *User) error {
	u.ID = len(s.users) + 1
	s.users = append(s.users, u)
	return nil
}

func (s *userService) Update(u *User) error {
	s.users[u.ID-1] = u
	return nil
}

func (s *userService) Delete(id int) error {
	s.users = slices.Delete(s.users, id-1, id)
	return nil
}

func (s *userService) Find(id int) (*User, error) {
	return s.users[id-1], nil
}

func (s *userService) List() ([]*User, error) {
	return s.users, nil
}
