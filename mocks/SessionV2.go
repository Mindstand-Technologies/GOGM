// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	go_cypherdsl "github.com/mindstand/go-cypherdsl"
	gogm "github.com/mindstand/gogm"

	mock "github.com/stretchr/testify/mock"

	neo4j "github.com/neo4j/neo4j-go-driver/neo4j"
)

// SessionV2 is an autogenerated mock type for the SessionV2 type
type SessionV2 struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *SessionV2) Begin() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *SessionV2) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *SessionV2) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: deleteObj
func (_m *SessionV2) Delete(deleteObj interface{}) error {
	ret := _m.Called(deleteObj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(deleteObj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUUID provides a mock function with given fields: uuid
func (_m *SessionV2) DeleteUUID(uuid string) error {
	ret := _m.Called(uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Load provides a mock function with given fields: respObj, id
func (_m *SessionV2) Load(respObj interface{}, id string) error {
	ret := _m.Called(respObj, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(respObj, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAll provides a mock function with given fields: respObj
func (_m *SessionV2) LoadAll(respObj interface{}) error {
	ret := _m.Called(respObj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(respObj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAllDepth provides a mock function with given fields: respObj, depth
func (_m *SessionV2) LoadAllDepth(respObj interface{}, depth int) error {
	ret := _m.Called(respObj, depth)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int) error); ok {
		r0 = rf(respObj, depth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAllDepthFilter provides a mock function with given fields: respObj, depth, filter, params
func (_m *SessionV2) LoadAllDepthFilter(respObj interface{}, depth int, filter go_cypherdsl.ConditionOperator, params map[string]interface{}) error {
	ret := _m.Called(respObj, depth, filter, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int, go_cypherdsl.ConditionOperator, map[string]interface{}) error); ok {
		r0 = rf(respObj, depth, filter, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAllDepthFilterPagination provides a mock function with given fields: respObj, depth, filter, params, pagination
func (_m *SessionV2) LoadAllDepthFilterPagination(respObj interface{}, depth int, filter go_cypherdsl.ConditionOperator, params map[string]interface{}, pagination *gogm.Pagination) error {
	ret := _m.Called(respObj, depth, filter, params, pagination)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int, go_cypherdsl.ConditionOperator, map[string]interface{}, *gogm.Pagination) error); ok {
		r0 = rf(respObj, depth, filter, params, pagination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadAllEdgeConstraint provides a mock function with given fields: respObj, endNodeType, endNodeField, edgeConstraint, minJumps, maxJumps, depth, filter
func (_m *SessionV2) LoadAllEdgeConstraint(respObj interface{}, endNodeType string, endNodeField string, edgeConstraint interface{}, minJumps int, maxJumps int, depth int, filter go_cypherdsl.ConditionOperator) error {
	ret := _m.Called(respObj, endNodeType, endNodeField, edgeConstraint, minJumps, maxJumps, depth, filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, string, interface{}, int, int, int, go_cypherdsl.ConditionOperator) error); ok {
		r0 = rf(respObj, endNodeType, endNodeField, edgeConstraint, minJumps, maxJumps, depth, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadDepth provides a mock function with given fields: respObj, id, depth
func (_m *SessionV2) LoadDepth(respObj interface{}, id string, depth int) error {
	ret := _m.Called(respObj, id, depth)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, int) error); ok {
		r0 = rf(respObj, id, depth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadDepthFilter provides a mock function with given fields: respObj, id, depth, filter, params
func (_m *SessionV2) LoadDepthFilter(respObj interface{}, id string, depth int, filter *go_cypherdsl.ConditionBuilder, params map[string]interface{}) error {
	ret := _m.Called(respObj, id, depth, filter, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, int, *go_cypherdsl.ConditionBuilder, map[string]interface{}) error); ok {
		r0 = rf(respObj, id, depth, filter, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadDepthFilterPagination provides a mock function with given fields: respObj, id, depth, filter, params, pagination
func (_m *SessionV2) LoadDepthFilterPagination(respObj interface{}, id string, depth int, filter go_cypherdsl.ConditionOperator, params map[string]interface{}, pagination *gogm.Pagination) error {
	ret := _m.Called(respObj, id, depth, filter, params, pagination)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, int, go_cypherdsl.ConditionOperator, map[string]interface{}, *gogm.Pagination) error); ok {
		r0 = rf(respObj, id, depth, filter, params, pagination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ManagedTransaction provides a mock function with given fields: work
func (_m *SessionV2) ManagedTransaction(work gogm.TransactionWork) error {
	ret := _m.Called(work)

	var r0 error
	if rf, ok := ret.Get(0).(func(gogm.TransactionWork) error); ok {
		r0 = rf(work)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PurgeDatabase provides a mock function with given fields:
func (_m *SessionV2) PurgeDatabase() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query provides a mock function with given fields: query, properties, respObj
func (_m *SessionV2) Query(query string, properties map[string]interface{}, respObj interface{}) error {
	ret := _m.Called(query, properties, respObj)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}, interface{}) error); ok {
		r0 = rf(query, properties, respObj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryRaw provides a mock function with given fields: query, properties
func (_m *SessionV2) QueryRaw(query string, properties map[string]interface{}) ([][]interface{}, neo4j.ResultSummary, error) {
	ret := _m.Called(query, properties)

	var r0 [][]interface{}
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) [][]interface{}); ok {
		r0 = rf(query, properties)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]interface{})
		}
	}

	var r1 neo4j.ResultSummary
	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) neo4j.ResultSummary); ok {
		r1 = rf(query, properties)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(neo4j.ResultSummary)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, map[string]interface{}) error); ok {
		r2 = rf(query, properties)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Rollback provides a mock function with given fields:
func (_m *SessionV2) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RollbackWithError provides a mock function with given fields: err
func (_m *SessionV2) RollbackWithError(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: saveObj
func (_m *SessionV2) Save(saveObj interface{}) error {
	ret := _m.Called(saveObj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(saveObj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDepth provides a mock function with given fields: saveObj, depth
func (_m *SessionV2) SaveDepth(saveObj interface{}, depth int) error {
	ret := _m.Called(saveObj, depth)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int) error); ok {
		r0 = rf(saveObj, depth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
