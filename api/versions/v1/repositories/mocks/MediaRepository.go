// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MediaRepository is an autogenerated mock type for the MediaRepository type
type MediaRepository struct {
	mock.Mock
}

// GetMediaByID provides a mock function with given fields: id
func (_m *MediaRepository) GetMediaByID(id string) (*interface{}, error) {
	ret := _m.Called(id)

	var r0 *interface{}
	if rf, ok := ret.Get(0).(func(string) *interface{}); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interface{})
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

// GetMediaByIDS provides a mock function with given fields: ids
func (_m *MediaRepository) GetMediaByIDS(ids []string) (*[]interface{}, error) {
	ret := _m.Called(ids)

	var r0 *[]interface{}
	if rf, ok := ret.Get(0).(func([]string) *[]interface{}); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
