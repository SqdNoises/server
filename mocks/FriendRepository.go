// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	model "github.com/aelpxy/xoniaapp/model"
	mock "github.com/stretchr/testify/mock"
)

// FriendRepository is an autogenerated mock type for the FriendRepository type
type FriendRepository struct {
	mock.Mock
}

// DeleteRequest provides a mock function with given fields: memberId, userId
func (_m *FriendRepository) DeleteRequest(memberId string, userId string) error {
	ret := _m.Called(memberId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(memberId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: id
func (_m *FriendRepository) FindByID(id string) (*model.User, error) {
	ret := _m.Called(id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FriendsList provides a mock function with given fields: id
func (_m *FriendRepository) FriendsList(id string) (*[]model.Friend, error) {
	ret := _m.Called(id)

	var r0 *[]model.Friend
	if rf, ok := ret.Get(0).(func(string) *[]model.Friend); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.Friend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveFriend provides a mock function with given fields: memberId, userId
func (_m *FriendRepository) RemoveFriend(memberId string, userId string) error {
	ret := _m.Called(memberId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(memberId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestList provides a mock function with given fields: id
func (_m *FriendRepository) RequestList(id string) (*[]model.FriendRequest, error) {
	ret := _m.Called(id)

	var r0 *[]model.FriendRequest
	if rf, ok := ret.Get(0).(func(string) *[]model.FriendRequest); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.FriendRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: user
func (_m *FriendRepository) Save(user *model.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}