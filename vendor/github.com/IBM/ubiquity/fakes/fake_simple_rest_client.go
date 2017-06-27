// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/IBM/ubiquity/local/scbe"
)

type FakeSimpleRestClient struct {
	LoginStub        func() error
	loginMutex       sync.RWMutex
	loginArgsForCall []struct{}
	loginReturns     struct {
		result1 error
	}
	loginReturnsOnCall map[int]struct {
		result1 error
	}
	PostStub        func(resource_url string, payload []byte, exitStatus int, v interface{}) error
	postMutex       sync.RWMutex
	postArgsForCall []struct {
		resource_url string
		payload      []byte
		exitStatus   int
		v            interface{}
	}
	postReturns struct {
		result1 error
	}
	postReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(resource_url string, params map[string]string, exitStatus int, v interface{}) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		resource_url string
		params       map[string]string
		exitStatus   int
		v            interface{}
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(resource_url string, payload []byte, exitStatus int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		resource_url string
		payload      []byte
		exitStatus   int
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSimpleRestClient) Login() error {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct{}{})
	fake.recordInvocation("Login", []interface{}{})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.loginReturns.result1
}

func (fake *FakeSimpleRestClient) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeSimpleRestClient) LoginReturns(result1 error) {
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimpleRestClient) LoginReturnsOnCall(i int, result1 error) {
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimpleRestClient) Post(resource_url string, payload []byte, exitStatus int, v interface{}) error {
	var payloadCopy []byte
	if payload != nil {
		payloadCopy = make([]byte, len(payload))
		copy(payloadCopy, payload)
	}
	fake.postMutex.Lock()
	ret, specificReturn := fake.postReturnsOnCall[len(fake.postArgsForCall)]
	fake.postArgsForCall = append(fake.postArgsForCall, struct {
		resource_url string
		payload      []byte
		exitStatus   int
		v            interface{}
	}{resource_url, payloadCopy, exitStatus, v})
	fake.recordInvocation("Post", []interface{}{resource_url, payloadCopy, exitStatus, v})
	fake.postMutex.Unlock()
	if fake.PostStub != nil {
		return fake.PostStub(resource_url, payload, exitStatus, v)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.postReturns.result1
}

func (fake *FakeSimpleRestClient) PostCallCount() int {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return len(fake.postArgsForCall)
}

func (fake *FakeSimpleRestClient) PostArgsForCall(i int) (string, []byte, int, interface{}) {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return fake.postArgsForCall[i].resource_url, fake.postArgsForCall[i].payload, fake.postArgsForCall[i].exitStatus, fake.postArgsForCall[i].v
}

func (fake *FakeSimpleRestClient) PostReturns(result1 error) {
	fake.PostStub = nil
	fake.postReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimpleRestClient) PostReturnsOnCall(i int, result1 error) {
	fake.PostStub = nil
	if fake.postReturnsOnCall == nil {
		fake.postReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.postReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimpleRestClient) Get(resource_url string, params map[string]string, exitStatus int, v interface{}) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		resource_url string
		params       map[string]string
		exitStatus   int
		v            interface{}
	}{resource_url, params, exitStatus, v})
	fake.recordInvocation("Get", []interface{}{resource_url, params, exitStatus, v})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(resource_url, params, exitStatus, v)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *FakeSimpleRestClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSimpleRestClient) GetArgsForCall(i int) (string, map[string]string, int, interface{}) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].resource_url, fake.getArgsForCall[i].params, fake.getArgsForCall[i].exitStatus, fake.getArgsForCall[i].v
}

func (fake *FakeSimpleRestClient) GetReturns(result1 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimpleRestClient) GetReturnsOnCall(i int, result1 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimpleRestClient) Delete(resource_url string, payload []byte, exitStatus int) error {
	var payloadCopy []byte
	if payload != nil {
		payloadCopy = make([]byte, len(payload))
		copy(payloadCopy, payload)
	}
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		resource_url string
		payload      []byte
		exitStatus   int
	}{resource_url, payloadCopy, exitStatus})
	fake.recordInvocation("Delete", []interface{}{resource_url, payloadCopy, exitStatus})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(resource_url, payload, exitStatus)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeSimpleRestClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSimpleRestClient) DeleteArgsForCall(i int) (string, []byte, int) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].resource_url, fake.deleteArgsForCall[i].payload, fake.deleteArgsForCall[i].exitStatus
}

func (fake *FakeSimpleRestClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimpleRestClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSimpleRestClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSimpleRestClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ scbe.SimpleRestClient = new(FakeSimpleRestClient)
