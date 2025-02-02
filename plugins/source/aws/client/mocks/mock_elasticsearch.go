// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: ElasticSearch)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	elasticsearchservice "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	gomock "github.com/golang/mock/gomock"
)

// MockElasticSearch is a mock of ElasticSearch interface.
type MockElasticSearch struct {
	ctrl     *gomock.Controller
	recorder *MockElasticSearchMockRecorder
}

// MockElasticSearchMockRecorder is the mock recorder for MockElasticSearch.
type MockElasticSearchMockRecorder struct {
	mock *MockElasticSearch
}

// NewMockElasticSearch creates a new mock instance.
func NewMockElasticSearch(ctrl *gomock.Controller) *MockElasticSearch {
	mock := &MockElasticSearch{ctrl: ctrl}
	mock.recorder = &MockElasticSearchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElasticSearch) EXPECT() *MockElasticSearchMockRecorder {
	return m.recorder
}

// DescribeElasticsearchDomain mocks base method.
func (m *MockElasticSearch) DescribeElasticsearchDomain(arg0 context.Context, arg1 *elasticsearchservice.DescribeElasticsearchDomainInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeElasticsearchDomain", varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeElasticsearchDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeElasticsearchDomain indicates an expected call of DescribeElasticsearchDomain.
func (mr *MockElasticSearchMockRecorder) DescribeElasticsearchDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeElasticsearchDomain", reflect.TypeOf((*MockElasticSearch)(nil).DescribeElasticsearchDomain), varargs...)
}

// ListDomainNames mocks base method.
func (m *MockElasticSearch) ListDomainNames(arg0 context.Context, arg1 *elasticsearchservice.ListDomainNamesInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListDomainNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDomainNames", varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDomainNames indicates an expected call of ListDomainNames.
func (mr *MockElasticSearchMockRecorder) ListDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomainNames", reflect.TypeOf((*MockElasticSearch)(nil).ListDomainNames), varargs...)
}

// ListTags mocks base method.
func (m *MockElasticSearch) ListTags(arg0 context.Context, arg1 *elasticsearchservice.ListTagsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockElasticSearchMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockElasticSearch)(nil).ListTags), varargs...)
}
