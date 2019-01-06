// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	context "context"
	sync "sync"

	lager "code.cloudfoundry.org/lager"
	db "github.com/concourse/concourse/atc/db"
	worker "github.com/concourse/concourse/atc/worker"
)

type FakeImage struct {
	FetchForContainerStub        func(context.Context, lager.Logger, db.CreatingContainer) (worker.FetchedImage, error)
	fetchForContainerMutex       sync.RWMutex
	fetchForContainerArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.CreatingContainer
	}
	fetchForContainerReturns struct {
		result1 worker.FetchedImage
		result2 error
	}
	fetchForContainerReturnsOnCall map[int]struct {
		result1 worker.FetchedImage
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImage) FetchForContainer(arg1 context.Context, arg2 lager.Logger, arg3 db.CreatingContainer) (worker.FetchedImage, error) {
	fake.fetchForContainerMutex.Lock()
	ret, specificReturn := fake.fetchForContainerReturnsOnCall[len(fake.fetchForContainerArgsForCall)]
	fake.fetchForContainerArgsForCall = append(fake.fetchForContainerArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.CreatingContainer
	}{arg1, arg2, arg3})
	fake.recordInvocation("FetchForContainer", []interface{}{arg1, arg2, arg3})
	fake.fetchForContainerMutex.Unlock()
	if fake.FetchForContainerStub != nil {
		return fake.FetchForContainerStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchForContainerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) FetchForContainerCallCount() int {
	fake.fetchForContainerMutex.RLock()
	defer fake.fetchForContainerMutex.RUnlock()
	return len(fake.fetchForContainerArgsForCall)
}

func (fake *FakeImage) FetchForContainerCalls(stub func(context.Context, lager.Logger, db.CreatingContainer) (worker.FetchedImage, error)) {
	fake.fetchForContainerMutex.Lock()
	defer fake.fetchForContainerMutex.Unlock()
	fake.FetchForContainerStub = stub
}

func (fake *FakeImage) FetchForContainerArgsForCall(i int) (context.Context, lager.Logger, db.CreatingContainer) {
	fake.fetchForContainerMutex.RLock()
	defer fake.fetchForContainerMutex.RUnlock()
	argsForCall := fake.fetchForContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImage) FetchForContainerReturns(result1 worker.FetchedImage, result2 error) {
	fake.fetchForContainerMutex.Lock()
	defer fake.fetchForContainerMutex.Unlock()
	fake.FetchForContainerStub = nil
	fake.fetchForContainerReturns = struct {
		result1 worker.FetchedImage
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) FetchForContainerReturnsOnCall(i int, result1 worker.FetchedImage, result2 error) {
	fake.fetchForContainerMutex.Lock()
	defer fake.fetchForContainerMutex.Unlock()
	fake.FetchForContainerStub = nil
	if fake.fetchForContainerReturnsOnCall == nil {
		fake.fetchForContainerReturnsOnCall = make(map[int]struct {
			result1 worker.FetchedImage
			result2 error
		})
	}
	fake.fetchForContainerReturnsOnCall[i] = struct {
		result1 worker.FetchedImage
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchForContainerMutex.RLock()
	defer fake.fetchForContainerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImage) recordInvocation(key string, args []interface{}) {
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

var _ worker.Image = new(FakeImage)
