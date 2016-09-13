// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/CyCoreSystems/ari (interfaces: Application)

package mock

import (
	ari "github.com/CyCoreSystems/ari"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Application interface
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *_MockApplicationRecorder
}

// Recorder for MockApplication (not exported)
type _MockApplicationRecorder struct {
	mock *MockApplication
}

func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &_MockApplicationRecorder{mock}
	return mock
}

func (_m *MockApplication) EXPECT() *_MockApplicationRecorder {
	return _m.recorder
}

func (_m *MockApplication) Data(_param0 string) (ari.ApplicationData, error) {
	ret := _m.ctrl.Call(_m, "Data", _param0)
	ret0, _ := ret[0].(ari.ApplicationData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockApplicationRecorder) Data(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Data", arg0)
}

func (_m *MockApplication) Get(_param0 string) *ari.ApplicationHandle {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*ari.ApplicationHandle)
	return ret0
}

func (_mr *_MockApplicationRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockApplication) List() ([]*ari.ApplicationHandle, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]*ari.ApplicationHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockApplicationRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}

func (_m *MockApplication) Subscribe(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Subscribe", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockApplicationRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Subscribe", arg0, arg1)
}

func (_m *MockApplication) Unsubscribe(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Unsubscribe", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockApplicationRecorder) Unsubscribe(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unsubscribe", arg0, arg1)
}