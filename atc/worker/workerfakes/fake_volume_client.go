// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	sync "sync"

	lager "code.cloudfoundry.org/lager"
	db "github.com/concourse/concourse/atc/db"
	worker "github.com/concourse/concourse/atc/worker"
)

type FakeVolumeClient struct {
	CreateVolumeForTaskCacheStub        func(lager.Logger, worker.VolumeSpec, int, int, string, string) (worker.Volume, error)
	createVolumeForTaskCacheMutex       sync.RWMutex
	createVolumeForTaskCacheArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 int
		arg4 int
		arg5 string
		arg6 string
	}
	createVolumeForTaskCacheReturns struct {
		result1 worker.Volume
		result2 error
	}
	createVolumeForTaskCacheReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 error
	}
	FindOrCreateCOWVolumeForContainerStub        func(lager.Logger, worker.VolumeSpec, db.CreatingContainer, worker.Volume, int, string) (worker.Volume, error)
	findOrCreateCOWVolumeForContainerMutex       sync.RWMutex
	findOrCreateCOWVolumeForContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 db.CreatingContainer
		arg4 worker.Volume
		arg5 int
		arg6 string
	}
	findOrCreateCOWVolumeForContainerReturns struct {
		result1 worker.Volume
		result2 error
	}
	findOrCreateCOWVolumeForContainerReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 error
	}
	FindOrCreateVolumeForBaseResourceTypeStub        func(lager.Logger, worker.VolumeSpec, int, string) (worker.Volume, error)
	findOrCreateVolumeForBaseResourceTypeMutex       sync.RWMutex
	findOrCreateVolumeForBaseResourceTypeArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 int
		arg4 string
	}
	findOrCreateVolumeForBaseResourceTypeReturns struct {
		result1 worker.Volume
		result2 error
	}
	findOrCreateVolumeForBaseResourceTypeReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 error
	}
	FindOrCreateVolumeForContainerStub        func(lager.Logger, worker.VolumeSpec, db.CreatingContainer, int, string) (worker.Volume, error)
	findOrCreateVolumeForContainerMutex       sync.RWMutex
	findOrCreateVolumeForContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 db.CreatingContainer
		arg4 int
		arg5 string
	}
	findOrCreateVolumeForContainerReturns struct {
		result1 worker.Volume
		result2 error
	}
	findOrCreateVolumeForContainerReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 error
	}
	FindOrCreateVolumeForResourceCertsStub        func(lager.Logger) (worker.Volume, bool, error)
	findOrCreateVolumeForResourceCertsMutex       sync.RWMutex
	findOrCreateVolumeForResourceCertsArgsForCall []struct {
		arg1 lager.Logger
	}
	findOrCreateVolumeForResourceCertsReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findOrCreateVolumeForResourceCertsReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	FindVolumeForResourceCacheStub        func(lager.Logger, db.UsedResourceCache) (worker.Volume, bool, error)
	findVolumeForResourceCacheMutex       sync.RWMutex
	findVolumeForResourceCacheArgsForCall []struct {
		arg1 lager.Logger
		arg2 db.UsedResourceCache
	}
	findVolumeForResourceCacheReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findVolumeForResourceCacheReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	FindVolumeForTaskCacheStub        func(lager.Logger, int, int, string, string) (worker.Volume, bool, error)
	findVolumeForTaskCacheMutex       sync.RWMutex
	findVolumeForTaskCacheArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 int
		arg4 string
		arg5 string
	}
	findVolumeForTaskCacheReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findVolumeForTaskCacheReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	LookupVolumeStub        func(lager.Logger, string) (worker.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	lookupVolumeReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeClient) CreateVolumeForTaskCache(arg1 lager.Logger, arg2 worker.VolumeSpec, arg3 int, arg4 int, arg5 string, arg6 string) (worker.Volume, error) {
	fake.createVolumeForTaskCacheMutex.Lock()
	ret, specificReturn := fake.createVolumeForTaskCacheReturnsOnCall[len(fake.createVolumeForTaskCacheArgsForCall)]
	fake.createVolumeForTaskCacheArgsForCall = append(fake.createVolumeForTaskCacheArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 int
		arg4 int
		arg5 string
		arg6 string
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("CreateVolumeForTaskCache", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.createVolumeForTaskCacheMutex.Unlock()
	if fake.CreateVolumeForTaskCacheStub != nil {
		return fake.CreateVolumeForTaskCacheStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createVolumeForTaskCacheReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolumeClient) CreateVolumeForTaskCacheCallCount() int {
	fake.createVolumeForTaskCacheMutex.RLock()
	defer fake.createVolumeForTaskCacheMutex.RUnlock()
	return len(fake.createVolumeForTaskCacheArgsForCall)
}

func (fake *FakeVolumeClient) CreateVolumeForTaskCacheCalls(stub func(lager.Logger, worker.VolumeSpec, int, int, string, string) (worker.Volume, error)) {
	fake.createVolumeForTaskCacheMutex.Lock()
	defer fake.createVolumeForTaskCacheMutex.Unlock()
	fake.CreateVolumeForTaskCacheStub = stub
}

func (fake *FakeVolumeClient) CreateVolumeForTaskCacheArgsForCall(i int) (lager.Logger, worker.VolumeSpec, int, int, string, string) {
	fake.createVolumeForTaskCacheMutex.RLock()
	defer fake.createVolumeForTaskCacheMutex.RUnlock()
	argsForCall := fake.createVolumeForTaskCacheArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeVolumeClient) CreateVolumeForTaskCacheReturns(result1 worker.Volume, result2 error) {
	fake.createVolumeForTaskCacheMutex.Lock()
	defer fake.createVolumeForTaskCacheMutex.Unlock()
	fake.CreateVolumeForTaskCacheStub = nil
	fake.createVolumeForTaskCacheReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) CreateVolumeForTaskCacheReturnsOnCall(i int, result1 worker.Volume, result2 error) {
	fake.createVolumeForTaskCacheMutex.Lock()
	defer fake.createVolumeForTaskCacheMutex.Unlock()
	fake.CreateVolumeForTaskCacheStub = nil
	if fake.createVolumeForTaskCacheReturnsOnCall == nil {
		fake.createVolumeForTaskCacheReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 error
		})
	}
	fake.createVolumeForTaskCacheReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateCOWVolumeForContainer(arg1 lager.Logger, arg2 worker.VolumeSpec, arg3 db.CreatingContainer, arg4 worker.Volume, arg5 int, arg6 string) (worker.Volume, error) {
	fake.findOrCreateCOWVolumeForContainerMutex.Lock()
	ret, specificReturn := fake.findOrCreateCOWVolumeForContainerReturnsOnCall[len(fake.findOrCreateCOWVolumeForContainerArgsForCall)]
	fake.findOrCreateCOWVolumeForContainerArgsForCall = append(fake.findOrCreateCOWVolumeForContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 db.CreatingContainer
		arg4 worker.Volume
		arg5 int
		arg6 string
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("FindOrCreateCOWVolumeForContainer", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.findOrCreateCOWVolumeForContainerMutex.Unlock()
	if fake.FindOrCreateCOWVolumeForContainerStub != nil {
		return fake.FindOrCreateCOWVolumeForContainerStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOrCreateCOWVolumeForContainerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolumeClient) FindOrCreateCOWVolumeForContainerCallCount() int {
	fake.findOrCreateCOWVolumeForContainerMutex.RLock()
	defer fake.findOrCreateCOWVolumeForContainerMutex.RUnlock()
	return len(fake.findOrCreateCOWVolumeForContainerArgsForCall)
}

func (fake *FakeVolumeClient) FindOrCreateCOWVolumeForContainerCalls(stub func(lager.Logger, worker.VolumeSpec, db.CreatingContainer, worker.Volume, int, string) (worker.Volume, error)) {
	fake.findOrCreateCOWVolumeForContainerMutex.Lock()
	defer fake.findOrCreateCOWVolumeForContainerMutex.Unlock()
	fake.FindOrCreateCOWVolumeForContainerStub = stub
}

func (fake *FakeVolumeClient) FindOrCreateCOWVolumeForContainerArgsForCall(i int) (lager.Logger, worker.VolumeSpec, db.CreatingContainer, worker.Volume, int, string) {
	fake.findOrCreateCOWVolumeForContainerMutex.RLock()
	defer fake.findOrCreateCOWVolumeForContainerMutex.RUnlock()
	argsForCall := fake.findOrCreateCOWVolumeForContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeVolumeClient) FindOrCreateCOWVolumeForContainerReturns(result1 worker.Volume, result2 error) {
	fake.findOrCreateCOWVolumeForContainerMutex.Lock()
	defer fake.findOrCreateCOWVolumeForContainerMutex.Unlock()
	fake.FindOrCreateCOWVolumeForContainerStub = nil
	fake.findOrCreateCOWVolumeForContainerReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateCOWVolumeForContainerReturnsOnCall(i int, result1 worker.Volume, result2 error) {
	fake.findOrCreateCOWVolumeForContainerMutex.Lock()
	defer fake.findOrCreateCOWVolumeForContainerMutex.Unlock()
	fake.FindOrCreateCOWVolumeForContainerStub = nil
	if fake.findOrCreateCOWVolumeForContainerReturnsOnCall == nil {
		fake.findOrCreateCOWVolumeForContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 error
		})
	}
	fake.findOrCreateCOWVolumeForContainerReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceType(arg1 lager.Logger, arg2 worker.VolumeSpec, arg3 int, arg4 string) (worker.Volume, error) {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.Lock()
	ret, specificReturn := fake.findOrCreateVolumeForBaseResourceTypeReturnsOnCall[len(fake.findOrCreateVolumeForBaseResourceTypeArgsForCall)]
	fake.findOrCreateVolumeForBaseResourceTypeArgsForCall = append(fake.findOrCreateVolumeForBaseResourceTypeArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 int
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("FindOrCreateVolumeForBaseResourceType", []interface{}{arg1, arg2, arg3, arg4})
	fake.findOrCreateVolumeForBaseResourceTypeMutex.Unlock()
	if fake.FindOrCreateVolumeForBaseResourceTypeStub != nil {
		return fake.FindOrCreateVolumeForBaseResourceTypeStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOrCreateVolumeForBaseResourceTypeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceTypeCallCount() int {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.RLock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.RUnlock()
	return len(fake.findOrCreateVolumeForBaseResourceTypeArgsForCall)
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceTypeCalls(stub func(lager.Logger, worker.VolumeSpec, int, string) (worker.Volume, error)) {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.Lock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.Unlock()
	fake.FindOrCreateVolumeForBaseResourceTypeStub = stub
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceTypeArgsForCall(i int) (lager.Logger, worker.VolumeSpec, int, string) {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.RLock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.RUnlock()
	argsForCall := fake.findOrCreateVolumeForBaseResourceTypeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceTypeReturns(result1 worker.Volume, result2 error) {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.Lock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.Unlock()
	fake.FindOrCreateVolumeForBaseResourceTypeStub = nil
	fake.findOrCreateVolumeForBaseResourceTypeReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForBaseResourceTypeReturnsOnCall(i int, result1 worker.Volume, result2 error) {
	fake.findOrCreateVolumeForBaseResourceTypeMutex.Lock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.Unlock()
	fake.FindOrCreateVolumeForBaseResourceTypeStub = nil
	if fake.findOrCreateVolumeForBaseResourceTypeReturnsOnCall == nil {
		fake.findOrCreateVolumeForBaseResourceTypeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 error
		})
	}
	fake.findOrCreateVolumeForBaseResourceTypeReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainer(arg1 lager.Logger, arg2 worker.VolumeSpec, arg3 db.CreatingContainer, arg4 int, arg5 string) (worker.Volume, error) {
	fake.findOrCreateVolumeForContainerMutex.Lock()
	ret, specificReturn := fake.findOrCreateVolumeForContainerReturnsOnCall[len(fake.findOrCreateVolumeForContainerArgsForCall)]
	fake.findOrCreateVolumeForContainerArgsForCall = append(fake.findOrCreateVolumeForContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
		arg3 db.CreatingContainer
		arg4 int
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("FindOrCreateVolumeForContainer", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.findOrCreateVolumeForContainerMutex.Unlock()
	if fake.FindOrCreateVolumeForContainerStub != nil {
		return fake.FindOrCreateVolumeForContainerStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOrCreateVolumeForContainerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainerCallCount() int {
	fake.findOrCreateVolumeForContainerMutex.RLock()
	defer fake.findOrCreateVolumeForContainerMutex.RUnlock()
	return len(fake.findOrCreateVolumeForContainerArgsForCall)
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainerCalls(stub func(lager.Logger, worker.VolumeSpec, db.CreatingContainer, int, string) (worker.Volume, error)) {
	fake.findOrCreateVolumeForContainerMutex.Lock()
	defer fake.findOrCreateVolumeForContainerMutex.Unlock()
	fake.FindOrCreateVolumeForContainerStub = stub
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainerArgsForCall(i int) (lager.Logger, worker.VolumeSpec, db.CreatingContainer, int, string) {
	fake.findOrCreateVolumeForContainerMutex.RLock()
	defer fake.findOrCreateVolumeForContainerMutex.RUnlock()
	argsForCall := fake.findOrCreateVolumeForContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainerReturns(result1 worker.Volume, result2 error) {
	fake.findOrCreateVolumeForContainerMutex.Lock()
	defer fake.findOrCreateVolumeForContainerMutex.Unlock()
	fake.FindOrCreateVolumeForContainerStub = nil
	fake.findOrCreateVolumeForContainerReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForContainerReturnsOnCall(i int, result1 worker.Volume, result2 error) {
	fake.findOrCreateVolumeForContainerMutex.Lock()
	defer fake.findOrCreateVolumeForContainerMutex.Unlock()
	fake.FindOrCreateVolumeForContainerStub = nil
	if fake.findOrCreateVolumeForContainerReturnsOnCall == nil {
		fake.findOrCreateVolumeForContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 error
		})
	}
	fake.findOrCreateVolumeForContainerReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCerts(arg1 lager.Logger) (worker.Volume, bool, error) {
	fake.findOrCreateVolumeForResourceCertsMutex.Lock()
	ret, specificReturn := fake.findOrCreateVolumeForResourceCertsReturnsOnCall[len(fake.findOrCreateVolumeForResourceCertsArgsForCall)]
	fake.findOrCreateVolumeForResourceCertsArgsForCall = append(fake.findOrCreateVolumeForResourceCertsArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("FindOrCreateVolumeForResourceCerts", []interface{}{arg1})
	fake.findOrCreateVolumeForResourceCertsMutex.Unlock()
	if fake.FindOrCreateVolumeForResourceCertsStub != nil {
		return fake.FindOrCreateVolumeForResourceCertsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findOrCreateVolumeForResourceCertsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCertsCallCount() int {
	fake.findOrCreateVolumeForResourceCertsMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCertsMutex.RUnlock()
	return len(fake.findOrCreateVolumeForResourceCertsArgsForCall)
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCertsCalls(stub func(lager.Logger) (worker.Volume, bool, error)) {
	fake.findOrCreateVolumeForResourceCertsMutex.Lock()
	defer fake.findOrCreateVolumeForResourceCertsMutex.Unlock()
	fake.FindOrCreateVolumeForResourceCertsStub = stub
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCertsArgsForCall(i int) lager.Logger {
	fake.findOrCreateVolumeForResourceCertsMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCertsMutex.RUnlock()
	argsForCall := fake.findOrCreateVolumeForResourceCertsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCertsReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.findOrCreateVolumeForResourceCertsMutex.Lock()
	defer fake.findOrCreateVolumeForResourceCertsMutex.Unlock()
	fake.FindOrCreateVolumeForResourceCertsStub = nil
	fake.findOrCreateVolumeForResourceCertsReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) FindOrCreateVolumeForResourceCertsReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.findOrCreateVolumeForResourceCertsMutex.Lock()
	defer fake.findOrCreateVolumeForResourceCertsMutex.Unlock()
	fake.FindOrCreateVolumeForResourceCertsStub = nil
	if fake.findOrCreateVolumeForResourceCertsReturnsOnCall == nil {
		fake.findOrCreateVolumeForResourceCertsReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findOrCreateVolumeForResourceCertsReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) FindVolumeForResourceCache(arg1 lager.Logger, arg2 db.UsedResourceCache) (worker.Volume, bool, error) {
	fake.findVolumeForResourceCacheMutex.Lock()
	ret, specificReturn := fake.findVolumeForResourceCacheReturnsOnCall[len(fake.findVolumeForResourceCacheArgsForCall)]
	fake.findVolumeForResourceCacheArgsForCall = append(fake.findVolumeForResourceCacheArgsForCall, struct {
		arg1 lager.Logger
		arg2 db.UsedResourceCache
	}{arg1, arg2})
	fake.recordInvocation("FindVolumeForResourceCache", []interface{}{arg1, arg2})
	fake.findVolumeForResourceCacheMutex.Unlock()
	if fake.FindVolumeForResourceCacheStub != nil {
		return fake.FindVolumeForResourceCacheStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findVolumeForResourceCacheReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeVolumeClient) FindVolumeForResourceCacheCallCount() int {
	fake.findVolumeForResourceCacheMutex.RLock()
	defer fake.findVolumeForResourceCacheMutex.RUnlock()
	return len(fake.findVolumeForResourceCacheArgsForCall)
}

func (fake *FakeVolumeClient) FindVolumeForResourceCacheCalls(stub func(lager.Logger, db.UsedResourceCache) (worker.Volume, bool, error)) {
	fake.findVolumeForResourceCacheMutex.Lock()
	defer fake.findVolumeForResourceCacheMutex.Unlock()
	fake.FindVolumeForResourceCacheStub = stub
}

func (fake *FakeVolumeClient) FindVolumeForResourceCacheArgsForCall(i int) (lager.Logger, db.UsedResourceCache) {
	fake.findVolumeForResourceCacheMutex.RLock()
	defer fake.findVolumeForResourceCacheMutex.RUnlock()
	argsForCall := fake.findVolumeForResourceCacheArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolumeClient) FindVolumeForResourceCacheReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.findVolumeForResourceCacheMutex.Lock()
	defer fake.findVolumeForResourceCacheMutex.Unlock()
	fake.FindVolumeForResourceCacheStub = nil
	fake.findVolumeForResourceCacheReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) FindVolumeForResourceCacheReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.findVolumeForResourceCacheMutex.Lock()
	defer fake.findVolumeForResourceCacheMutex.Unlock()
	fake.FindVolumeForResourceCacheStub = nil
	if fake.findVolumeForResourceCacheReturnsOnCall == nil {
		fake.findVolumeForResourceCacheReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findVolumeForResourceCacheReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) FindVolumeForTaskCache(arg1 lager.Logger, arg2 int, arg3 int, arg4 string, arg5 string) (worker.Volume, bool, error) {
	fake.findVolumeForTaskCacheMutex.Lock()
	ret, specificReturn := fake.findVolumeForTaskCacheReturnsOnCall[len(fake.findVolumeForTaskCacheArgsForCall)]
	fake.findVolumeForTaskCacheArgsForCall = append(fake.findVolumeForTaskCacheArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 int
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("FindVolumeForTaskCache", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.findVolumeForTaskCacheMutex.Unlock()
	if fake.FindVolumeForTaskCacheStub != nil {
		return fake.FindVolumeForTaskCacheStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findVolumeForTaskCacheReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeVolumeClient) FindVolumeForTaskCacheCallCount() int {
	fake.findVolumeForTaskCacheMutex.RLock()
	defer fake.findVolumeForTaskCacheMutex.RUnlock()
	return len(fake.findVolumeForTaskCacheArgsForCall)
}

func (fake *FakeVolumeClient) FindVolumeForTaskCacheCalls(stub func(lager.Logger, int, int, string, string) (worker.Volume, bool, error)) {
	fake.findVolumeForTaskCacheMutex.Lock()
	defer fake.findVolumeForTaskCacheMutex.Unlock()
	fake.FindVolumeForTaskCacheStub = stub
}

func (fake *FakeVolumeClient) FindVolumeForTaskCacheArgsForCall(i int) (lager.Logger, int, int, string, string) {
	fake.findVolumeForTaskCacheMutex.RLock()
	defer fake.findVolumeForTaskCacheMutex.RUnlock()
	argsForCall := fake.findVolumeForTaskCacheArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeVolumeClient) FindVolumeForTaskCacheReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.findVolumeForTaskCacheMutex.Lock()
	defer fake.findVolumeForTaskCacheMutex.Unlock()
	fake.FindVolumeForTaskCacheStub = nil
	fake.findVolumeForTaskCacheReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) FindVolumeForTaskCacheReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.findVolumeForTaskCacheMutex.Lock()
	defer fake.findVolumeForTaskCacheMutex.Unlock()
	fake.FindVolumeForTaskCacheStub = nil
	if fake.findVolumeForTaskCacheReturnsOnCall == nil {
		fake.findVolumeForTaskCacheReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findVolumeForTaskCacheReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) LookupVolume(arg1 lager.Logger, arg2 string) (worker.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	ret, specificReturn := fake.lookupVolumeReturnsOnCall[len(fake.lookupVolumeArgsForCall)]
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupVolume", []interface{}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.lookupVolumeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeVolumeClient) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeVolumeClient) LookupVolumeCalls(stub func(lager.Logger, string) (worker.Volume, bool, error)) {
	fake.lookupVolumeMutex.Lock()
	defer fake.lookupVolumeMutex.Unlock()
	fake.LookupVolumeStub = stub
}

func (fake *FakeVolumeClient) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	argsForCall := fake.lookupVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolumeClient) LookupVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.lookupVolumeMutex.Lock()
	defer fake.lookupVolumeMutex.Unlock()
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) LookupVolumeReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.lookupVolumeMutex.Lock()
	defer fake.lookupVolumeMutex.Unlock()
	fake.LookupVolumeStub = nil
	if fake.lookupVolumeReturnsOnCall == nil {
		fake.lookupVolumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.lookupVolumeReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createVolumeForTaskCacheMutex.RLock()
	defer fake.createVolumeForTaskCacheMutex.RUnlock()
	fake.findOrCreateCOWVolumeForContainerMutex.RLock()
	defer fake.findOrCreateCOWVolumeForContainerMutex.RUnlock()
	fake.findOrCreateVolumeForBaseResourceTypeMutex.RLock()
	defer fake.findOrCreateVolumeForBaseResourceTypeMutex.RUnlock()
	fake.findOrCreateVolumeForContainerMutex.RLock()
	defer fake.findOrCreateVolumeForContainerMutex.RUnlock()
	fake.findOrCreateVolumeForResourceCertsMutex.RLock()
	defer fake.findOrCreateVolumeForResourceCertsMutex.RUnlock()
	fake.findVolumeForResourceCacheMutex.RLock()
	defer fake.findVolumeForResourceCacheMutex.RUnlock()
	fake.findVolumeForTaskCacheMutex.RLock()
	defer fake.findVolumeForTaskCacheMutex.RUnlock()
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolumeClient) recordInvocation(key string, args []interface{}) {
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

var _ worker.VolumeClient = new(FakeVolumeClient)
