package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// Sound is an autogenerated mock type for the Sound type
type Sound struct {
	mock.Mock
}

// Data provides a mock function with given fields: key
func (_m *Sound) Data(key *ari.Key) (*ari.SoundData, error) {
	ret := _m.Called(key)

	var r0 *ari.SoundData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.SoundData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.SoundData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: key
func (_m *Sound) Get(key *ari.Key) *ari.SoundHandle {
	ret := _m.Called(key)

	var r0 *ari.SoundHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.SoundHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.SoundHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields: filters, keyFilter
func (_m *Sound) List(filters map[string]string, keyFilter *ari.Key) ([]*ari.Key, error) {
	ret := _m.Called(filters, keyFilter)

	var r0 []*ari.Key
	if rf, ok := ret.Get(0).(func(map[string]string, *ari.Key) []*ari.Key); ok {
		r0 = rf(filters, keyFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ari.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string, *ari.Key) error); ok {
		r1 = rf(filters, keyFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
