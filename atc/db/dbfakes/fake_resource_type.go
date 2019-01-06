// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	sync "sync"

	lager "code.cloudfoundry.org/lager"
	atc "github.com/concourse/concourse/atc"
	creds "github.com/concourse/concourse/atc/creds"
	db "github.com/concourse/concourse/atc/db"
)

type FakeResourceType struct {
	CheckErrorStub        func() error
	checkErrorMutex       sync.RWMutex
	checkErrorArgsForCall []struct {
	}
	checkErrorReturns struct {
		result1 error
	}
	checkErrorReturnsOnCall map[int]struct {
		result1 error
	}
	CheckEveryStub        func() string
	checkEveryMutex       sync.RWMutex
	checkEveryArgsForCall []struct {
	}
	checkEveryReturns struct {
		result1 string
	}
	checkEveryReturnsOnCall map[int]struct {
		result1 string
	}
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ParamsStub        func() atc.Params
	paramsMutex       sync.RWMutex
	paramsArgsForCall []struct {
	}
	paramsReturns struct {
		result1 atc.Params
	}
	paramsReturnsOnCall map[int]struct {
		result1 atc.Params
	}
	PrivilegedStub        func() bool
	privilegedMutex       sync.RWMutex
	privilegedArgsForCall []struct {
	}
	privilegedReturns struct {
		result1 bool
	}
	privilegedReturnsOnCall map[int]struct {
		result1 bool
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct {
	}
	reloadReturns struct {
		result1 bool
		result2 error
	}
	reloadReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ResourceConfigCheckErrorStub        func() error
	resourceConfigCheckErrorMutex       sync.RWMutex
	resourceConfigCheckErrorArgsForCall []struct {
	}
	resourceConfigCheckErrorReturns struct {
		result1 error
	}
	resourceConfigCheckErrorReturnsOnCall map[int]struct {
		result1 error
	}
	SetCheckErrorStub        func(error) error
	setCheckErrorMutex       sync.RWMutex
	setCheckErrorArgsForCall []struct {
		arg1 error
	}
	setCheckErrorReturns struct {
		result1 error
	}
	setCheckErrorReturnsOnCall map[int]struct {
		result1 error
	}
	SetResourceConfigStub        func(lager.Logger, atc.Source, creds.VersionedResourceTypes) (db.ResourceConfig, error)
	setResourceConfigMutex       sync.RWMutex
	setResourceConfigArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.Source
		arg3 creds.VersionedResourceTypes
	}
	setResourceConfigReturns struct {
		result1 db.ResourceConfig
		result2 error
	}
	setResourceConfigReturnsOnCall map[int]struct {
		result1 db.ResourceConfig
		result2 error
	}
	SourceStub        func() atc.Source
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct {
	}
	sourceReturns struct {
		result1 atc.Source
	}
	sourceReturnsOnCall map[int]struct {
		result1 atc.Source
	}
	TagsStub        func() atc.Tags
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct {
	}
	tagsReturns struct {
		result1 atc.Tags
	}
	tagsReturnsOnCall map[int]struct {
		result1 atc.Tags
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 atc.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceType) CheckError() error {
	fake.checkErrorMutex.Lock()
	ret, specificReturn := fake.checkErrorReturnsOnCall[len(fake.checkErrorArgsForCall)]
	fake.checkErrorArgsForCall = append(fake.checkErrorArgsForCall, struct {
	}{})
	fake.recordInvocation("CheckError", []interface{}{})
	fake.checkErrorMutex.Unlock()
	if fake.CheckErrorStub != nil {
		return fake.CheckErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkErrorReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) CheckErrorCallCount() int {
	fake.checkErrorMutex.RLock()
	defer fake.checkErrorMutex.RUnlock()
	return len(fake.checkErrorArgsForCall)
}

func (fake *FakeResourceType) CheckErrorCalls(stub func() error) {
	fake.checkErrorMutex.Lock()
	defer fake.checkErrorMutex.Unlock()
	fake.CheckErrorStub = stub
}

func (fake *FakeResourceType) CheckErrorReturns(result1 error) {
	fake.checkErrorMutex.Lock()
	defer fake.checkErrorMutex.Unlock()
	fake.CheckErrorStub = nil
	fake.checkErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceType) CheckErrorReturnsOnCall(i int, result1 error) {
	fake.checkErrorMutex.Lock()
	defer fake.checkErrorMutex.Unlock()
	fake.CheckErrorStub = nil
	if fake.checkErrorReturnsOnCall == nil {
		fake.checkErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceType) CheckEvery() string {
	fake.checkEveryMutex.Lock()
	ret, specificReturn := fake.checkEveryReturnsOnCall[len(fake.checkEveryArgsForCall)]
	fake.checkEveryArgsForCall = append(fake.checkEveryArgsForCall, struct {
	}{})
	fake.recordInvocation("CheckEvery", []interface{}{})
	fake.checkEveryMutex.Unlock()
	if fake.CheckEveryStub != nil {
		return fake.CheckEveryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkEveryReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) CheckEveryCallCount() int {
	fake.checkEveryMutex.RLock()
	defer fake.checkEveryMutex.RUnlock()
	return len(fake.checkEveryArgsForCall)
}

func (fake *FakeResourceType) CheckEveryCalls(stub func() string) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = stub
}

func (fake *FakeResourceType) CheckEveryReturns(result1 string) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = nil
	fake.checkEveryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) CheckEveryReturnsOnCall(i int, result1 string) {
	fake.checkEveryMutex.Lock()
	defer fake.checkEveryMutex.Unlock()
	fake.CheckEveryStub = nil
	if fake.checkEveryReturnsOnCall == nil {
		fake.checkEveryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.checkEveryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeResourceType) IDCalls(stub func() int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeResourceType) IDReturns(result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceType) IDReturnsOnCall(i int, result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceType) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeResourceType) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeResourceType) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) Params() atc.Params {
	fake.paramsMutex.Lock()
	ret, specificReturn := fake.paramsReturnsOnCall[len(fake.paramsArgsForCall)]
	fake.paramsArgsForCall = append(fake.paramsArgsForCall, struct {
	}{})
	fake.recordInvocation("Params", []interface{}{})
	fake.paramsMutex.Unlock()
	if fake.ParamsStub != nil {
		return fake.ParamsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.paramsReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) ParamsCallCount() int {
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	return len(fake.paramsArgsForCall)
}

func (fake *FakeResourceType) ParamsCalls(stub func() atc.Params) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = stub
}

func (fake *FakeResourceType) ParamsReturns(result1 atc.Params) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = nil
	fake.paramsReturns = struct {
		result1 atc.Params
	}{result1}
}

func (fake *FakeResourceType) ParamsReturnsOnCall(i int, result1 atc.Params) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = nil
	if fake.paramsReturnsOnCall == nil {
		fake.paramsReturnsOnCall = make(map[int]struct {
			result1 atc.Params
		})
	}
	fake.paramsReturnsOnCall[i] = struct {
		result1 atc.Params
	}{result1}
}

func (fake *FakeResourceType) Privileged() bool {
	fake.privilegedMutex.Lock()
	ret, specificReturn := fake.privilegedReturnsOnCall[len(fake.privilegedArgsForCall)]
	fake.privilegedArgsForCall = append(fake.privilegedArgsForCall, struct {
	}{})
	fake.recordInvocation("Privileged", []interface{}{})
	fake.privilegedMutex.Unlock()
	if fake.PrivilegedStub != nil {
		return fake.PrivilegedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.privilegedReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) PrivilegedCallCount() int {
	fake.privilegedMutex.RLock()
	defer fake.privilegedMutex.RUnlock()
	return len(fake.privilegedArgsForCall)
}

func (fake *FakeResourceType) PrivilegedCalls(stub func() bool) {
	fake.privilegedMutex.Lock()
	defer fake.privilegedMutex.Unlock()
	fake.PrivilegedStub = stub
}

func (fake *FakeResourceType) PrivilegedReturns(result1 bool) {
	fake.privilegedMutex.Lock()
	defer fake.privilegedMutex.Unlock()
	fake.PrivilegedStub = nil
	fake.privilegedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeResourceType) PrivilegedReturnsOnCall(i int, result1 bool) {
	fake.privilegedMutex.Lock()
	defer fake.privilegedMutex.Unlock()
	fake.PrivilegedStub = nil
	if fake.privilegedReturnsOnCall == nil {
		fake.privilegedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.privilegedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeResourceType) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	ret, specificReturn := fake.reloadReturnsOnCall[len(fake.reloadArgsForCall)]
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct {
	}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.reloadReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceType) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeResourceType) ReloadCalls(stub func() (bool, error)) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = stub
}

func (fake *FakeResourceType) ReloadReturns(result1 bool, result2 error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceType) ReloadReturnsOnCall(i int, result1 bool, result2 error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = nil
	if fake.reloadReturnsOnCall == nil {
		fake.reloadReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.reloadReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceType) ResourceConfigCheckError() error {
	fake.resourceConfigCheckErrorMutex.Lock()
	ret, specificReturn := fake.resourceConfigCheckErrorReturnsOnCall[len(fake.resourceConfigCheckErrorArgsForCall)]
	fake.resourceConfigCheckErrorArgsForCall = append(fake.resourceConfigCheckErrorArgsForCall, struct {
	}{})
	fake.recordInvocation("ResourceConfigCheckError", []interface{}{})
	fake.resourceConfigCheckErrorMutex.Unlock()
	if fake.ResourceConfigCheckErrorStub != nil {
		return fake.ResourceConfigCheckErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.resourceConfigCheckErrorReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) ResourceConfigCheckErrorCallCount() int {
	fake.resourceConfigCheckErrorMutex.RLock()
	defer fake.resourceConfigCheckErrorMutex.RUnlock()
	return len(fake.resourceConfigCheckErrorArgsForCall)
}

func (fake *FakeResourceType) ResourceConfigCheckErrorCalls(stub func() error) {
	fake.resourceConfigCheckErrorMutex.Lock()
	defer fake.resourceConfigCheckErrorMutex.Unlock()
	fake.ResourceConfigCheckErrorStub = stub
}

func (fake *FakeResourceType) ResourceConfigCheckErrorReturns(result1 error) {
	fake.resourceConfigCheckErrorMutex.Lock()
	defer fake.resourceConfigCheckErrorMutex.Unlock()
	fake.ResourceConfigCheckErrorStub = nil
	fake.resourceConfigCheckErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceType) ResourceConfigCheckErrorReturnsOnCall(i int, result1 error) {
	fake.resourceConfigCheckErrorMutex.Lock()
	defer fake.resourceConfigCheckErrorMutex.Unlock()
	fake.ResourceConfigCheckErrorStub = nil
	if fake.resourceConfigCheckErrorReturnsOnCall == nil {
		fake.resourceConfigCheckErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.resourceConfigCheckErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceType) SetCheckError(arg1 error) error {
	fake.setCheckErrorMutex.Lock()
	ret, specificReturn := fake.setCheckErrorReturnsOnCall[len(fake.setCheckErrorArgsForCall)]
	fake.setCheckErrorArgsForCall = append(fake.setCheckErrorArgsForCall, struct {
		arg1 error
	}{arg1})
	fake.recordInvocation("SetCheckError", []interface{}{arg1})
	fake.setCheckErrorMutex.Unlock()
	if fake.SetCheckErrorStub != nil {
		return fake.SetCheckErrorStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setCheckErrorReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) SetCheckErrorCallCount() int {
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	return len(fake.setCheckErrorArgsForCall)
}

func (fake *FakeResourceType) SetCheckErrorCalls(stub func(error) error) {
	fake.setCheckErrorMutex.Lock()
	defer fake.setCheckErrorMutex.Unlock()
	fake.SetCheckErrorStub = stub
}

func (fake *FakeResourceType) SetCheckErrorArgsForCall(i int) error {
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	argsForCall := fake.setCheckErrorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResourceType) SetCheckErrorReturns(result1 error) {
	fake.setCheckErrorMutex.Lock()
	defer fake.setCheckErrorMutex.Unlock()
	fake.SetCheckErrorStub = nil
	fake.setCheckErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceType) SetCheckErrorReturnsOnCall(i int, result1 error) {
	fake.setCheckErrorMutex.Lock()
	defer fake.setCheckErrorMutex.Unlock()
	fake.SetCheckErrorStub = nil
	if fake.setCheckErrorReturnsOnCall == nil {
		fake.setCheckErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setCheckErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceType) SetResourceConfig(arg1 lager.Logger, arg2 atc.Source, arg3 creds.VersionedResourceTypes) (db.ResourceConfig, error) {
	fake.setResourceConfigMutex.Lock()
	ret, specificReturn := fake.setResourceConfigReturnsOnCall[len(fake.setResourceConfigArgsForCall)]
	fake.setResourceConfigArgsForCall = append(fake.setResourceConfigArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.Source
		arg3 creds.VersionedResourceTypes
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetResourceConfig", []interface{}{arg1, arg2, arg3})
	fake.setResourceConfigMutex.Unlock()
	if fake.SetResourceConfigStub != nil {
		return fake.SetResourceConfigStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setResourceConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceType) SetResourceConfigCallCount() int {
	fake.setResourceConfigMutex.RLock()
	defer fake.setResourceConfigMutex.RUnlock()
	return len(fake.setResourceConfigArgsForCall)
}

func (fake *FakeResourceType) SetResourceConfigCalls(stub func(lager.Logger, atc.Source, creds.VersionedResourceTypes) (db.ResourceConfig, error)) {
	fake.setResourceConfigMutex.Lock()
	defer fake.setResourceConfigMutex.Unlock()
	fake.SetResourceConfigStub = stub
}

func (fake *FakeResourceType) SetResourceConfigArgsForCall(i int) (lager.Logger, atc.Source, creds.VersionedResourceTypes) {
	fake.setResourceConfigMutex.RLock()
	defer fake.setResourceConfigMutex.RUnlock()
	argsForCall := fake.setResourceConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeResourceType) SetResourceConfigReturns(result1 db.ResourceConfig, result2 error) {
	fake.setResourceConfigMutex.Lock()
	defer fake.setResourceConfigMutex.Unlock()
	fake.SetResourceConfigStub = nil
	fake.setResourceConfigReturns = struct {
		result1 db.ResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceType) SetResourceConfigReturnsOnCall(i int, result1 db.ResourceConfig, result2 error) {
	fake.setResourceConfigMutex.Lock()
	defer fake.setResourceConfigMutex.Unlock()
	fake.SetResourceConfigStub = nil
	if fake.setResourceConfigReturnsOnCall == nil {
		fake.setResourceConfigReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfig
			result2 error
		})
	}
	fake.setResourceConfigReturnsOnCall[i] = struct {
		result1 db.ResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceType) Source() atc.Source {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct {
	}{})
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeResourceType) SourceCalls(stub func() atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = stub
}

func (fake *FakeResourceType) SourceReturns(result1 atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResourceType) SourceReturnsOnCall(i int, result1 atc.Source) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 atc.Source
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResourceType) Tags() atc.Tags {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct {
	}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tagsReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeResourceType) TagsCalls(stub func() atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = stub
}

func (fake *FakeResourceType) TagsReturns(result1 atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeResourceType) TagsReturnsOnCall(i int, result1 atc.Tags) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 atc.Tags
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 atc.Tags
	}{result1}
}

func (fake *FakeResourceType) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.typeReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeResourceType) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeResourceType) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) TypeReturnsOnCall(i int, result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) Version() atc.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.versionReturns
	return fakeReturns.result1
}

func (fake *FakeResourceType) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeResourceType) VersionCalls(stub func() atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeResourceType) VersionReturns(result1 atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceType) VersionReturnsOnCall(i int, result1 atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceType) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkErrorMutex.RLock()
	defer fake.checkErrorMutex.RUnlock()
	fake.checkEveryMutex.RLock()
	defer fake.checkEveryMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	fake.privilegedMutex.RLock()
	defer fake.privilegedMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.resourceConfigCheckErrorMutex.RLock()
	defer fake.resourceConfigCheckErrorMutex.RUnlock()
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	fake.setResourceConfigMutex.RLock()
	defer fake.setResourceConfigMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceType) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceType = new(FakeResourceType)
