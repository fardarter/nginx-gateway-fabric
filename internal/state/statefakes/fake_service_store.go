// Code generated by counterfeiter. DO NOT EDIT.
package statefakes

import (
	"sync"

	"github.com/nginxinc/nginx-kubernetes-gateway/internal/state"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type FakeServiceStore struct {
	DeleteStub        func(types.NamespacedName)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 types.NamespacedName
	}
	ResolveStub        func(types.NamespacedName) (string, error)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		arg1 types.NamespacedName
	}
	resolveReturns struct {
		result1 string
		result2 error
	}
	resolveReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UpsertStub        func(*v1.Service)
	upsertMutex       sync.RWMutex
	upsertArgsForCall []struct {
		arg1 *v1.Service
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceStore) Delete(arg1 types.NamespacedName) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 types.NamespacedName
	}{arg1})
	stub := fake.DeleteStub
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		fake.DeleteStub(arg1)
	}
}

func (fake *FakeServiceStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeServiceStore) DeleteCalls(stub func(types.NamespacedName)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeServiceStore) DeleteArgsForCall(i int) types.NamespacedName {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceStore) Resolve(arg1 types.NamespacedName) (string, error) {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		arg1 types.NamespacedName
	}{arg1})
	stub := fake.ResolveStub
	fakeReturns := fake.resolveReturns
	fake.recordInvocation("Resolve", []interface{}{arg1})
	fake.resolveMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceStore) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *FakeServiceStore) ResolveCalls(stub func(types.NamespacedName) (string, error)) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = stub
}

func (fake *FakeServiceStore) ResolveArgsForCall(i int) types.NamespacedName {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	argsForCall := fake.resolveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceStore) ResolveReturns(result1 string, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceStore) ResolveReturnsOnCall(i int, result1 string, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceStore) Upsert(arg1 *v1.Service) {
	fake.upsertMutex.Lock()
	fake.upsertArgsForCall = append(fake.upsertArgsForCall, struct {
		arg1 *v1.Service
	}{arg1})
	stub := fake.UpsertStub
	fake.recordInvocation("Upsert", []interface{}{arg1})
	fake.upsertMutex.Unlock()
	if stub != nil {
		fake.UpsertStub(arg1)
	}
}

func (fake *FakeServiceStore) UpsertCallCount() int {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	return len(fake.upsertArgsForCall)
}

func (fake *FakeServiceStore) UpsertCalls(stub func(*v1.Service)) {
	fake.upsertMutex.Lock()
	defer fake.upsertMutex.Unlock()
	fake.UpsertStub = stub
}

func (fake *FakeServiceStore) UpsertArgsForCall(i int) *v1.Service {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	argsForCall := fake.upsertArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceStore) recordInvocation(key string, args []interface{}) {
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

var _ state.ServiceStore = new(FakeServiceStore)