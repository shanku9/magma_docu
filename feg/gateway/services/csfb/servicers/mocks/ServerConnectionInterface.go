// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// Run make gen at FeG to re-generate

package mocks

import mock "github.com/stretchr/testify/mock"

// ServerConnectionInterface is an autogenerated mock type for the ServerConnectionInterface type
type ServerConnectionInterface struct {
	mock.Mock
}

// AcceptConn provides a mock function with given fields:
func (_m *ServerConnectionInterface) AcceptConn() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseConn provides a mock function with given fields:
func (_m *ServerConnectionInterface) CloseConn() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseListener provides a mock function with given fields:
func (_m *ServerConnectionInterface) CloseListener() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectionEstablished provides a mock function with given fields:
func (_m *ServerConnectionInterface) ConnectionEstablished() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ReceiveThroughListener provides a mock function with given fields:
func (_m *ServerConnectionInterface) ReceiveThroughListener() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendFromServer provides a mock function with given fields: _a0
func (_m *ServerConnectionInterface) SendFromServer(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartListener provides a mock function with given fields: ipAddr, port
func (_m *ServerConnectionInterface) StartListener(ipAddr string, port int) (int, error) {
	ret := _m.Called(ipAddr, port)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, int) int); ok {
		r0 = rf(ipAddr, port)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(ipAddr, port)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
