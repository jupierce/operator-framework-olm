// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/operator-framework/operator-registry/pkg/api"
	"google.golang.org/grpc/metadata"
)

type FakeRegistry_ListBundlesClient struct {
	CloseSendStub        func() error
	closeSendMutex       sync.RWMutex
	closeSendArgsForCall []struct {
	}
	closeSendReturns struct {
		result1 error
	}
	closeSendReturnsOnCall map[int]struct {
		result1 error
	}
	ContextStub        func() context.Context
	contextMutex       sync.RWMutex
	contextArgsForCall []struct {
	}
	contextReturns struct {
		result1 context.Context
	}
	contextReturnsOnCall map[int]struct {
		result1 context.Context
	}
	HeaderStub        func() (metadata.MD, error)
	headerMutex       sync.RWMutex
	headerArgsForCall []struct {
	}
	headerReturns struct {
		result1 metadata.MD
		result2 error
	}
	headerReturnsOnCall map[int]struct {
		result1 metadata.MD
		result2 error
	}
	RecvStub        func() (*api.Bundle, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct {
	}
	recvReturns struct {
		result1 *api.Bundle
		result2 error
	}
	recvReturnsOnCall map[int]struct {
		result1 *api.Bundle
		result2 error
	}
	RecvMsgStub        func(interface{}) error
	recvMsgMutex       sync.RWMutex
	recvMsgArgsForCall []struct {
		arg1 interface{}
	}
	recvMsgReturns struct {
		result1 error
	}
	recvMsgReturnsOnCall map[int]struct {
		result1 error
	}
	SendMsgStub        func(interface{}) error
	sendMsgMutex       sync.RWMutex
	sendMsgArgsForCall []struct {
		arg1 interface{}
	}
	sendMsgReturns struct {
		result1 error
	}
	sendMsgReturnsOnCall map[int]struct {
		result1 error
	}
	TrailerStub        func() metadata.MD
	trailerMutex       sync.RWMutex
	trailerArgsForCall []struct {
	}
	trailerReturns struct {
		result1 metadata.MD
	}
	trailerReturnsOnCall map[int]struct {
		result1 metadata.MD
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistry_ListBundlesClient) CloseSend() error {
	fake.closeSendMutex.Lock()
	ret, specificReturn := fake.closeSendReturnsOnCall[len(fake.closeSendArgsForCall)]
	fake.closeSendArgsForCall = append(fake.closeSendArgsForCall, struct {
	}{})
	fake.recordInvocation("CloseSend", []interface{}{})
	fake.closeSendMutex.Unlock()
	if fake.CloseSendStub != nil {
		return fake.CloseSendStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeSendReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry_ListBundlesClient) CloseSendCallCount() int {
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	return len(fake.closeSendArgsForCall)
}

func (fake *FakeRegistry_ListBundlesClient) CloseSendCalls(stub func() error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = stub
}

func (fake *FakeRegistry_ListBundlesClient) CloseSendReturns(result1 error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = nil
	fake.closeSendReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) CloseSendReturnsOnCall(i int, result1 error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = nil
	if fake.closeSendReturnsOnCall == nil {
		fake.closeSendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeSendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) Context() context.Context {
	fake.contextMutex.Lock()
	ret, specificReturn := fake.contextReturnsOnCall[len(fake.contextArgsForCall)]
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct {
	}{})
	fake.recordInvocation("Context", []interface{}{})
	fake.contextMutex.Unlock()
	if fake.ContextStub != nil {
		return fake.ContextStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.contextReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry_ListBundlesClient) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *FakeRegistry_ListBundlesClient) ContextCalls(stub func() context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = stub
}

func (fake *FakeRegistry_ListBundlesClient) ContextReturns(result1 context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 context.Context
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) ContextReturnsOnCall(i int, result1 context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = nil
	if fake.contextReturnsOnCall == nil {
		fake.contextReturnsOnCall = make(map[int]struct {
			result1 context.Context
		})
	}
	fake.contextReturnsOnCall[i] = struct {
		result1 context.Context
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) Header() (metadata.MD, error) {
	fake.headerMutex.Lock()
	ret, specificReturn := fake.headerReturnsOnCall[len(fake.headerArgsForCall)]
	fake.headerArgsForCall = append(fake.headerArgsForCall, struct {
	}{})
	fake.recordInvocation("Header", []interface{}{})
	fake.headerMutex.Unlock()
	if fake.HeaderStub != nil {
		return fake.HeaderStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.headerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegistry_ListBundlesClient) HeaderCallCount() int {
	fake.headerMutex.RLock()
	defer fake.headerMutex.RUnlock()
	return len(fake.headerArgsForCall)
}

func (fake *FakeRegistry_ListBundlesClient) HeaderCalls(stub func() (metadata.MD, error)) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = stub
}

func (fake *FakeRegistry_ListBundlesClient) HeaderReturns(result1 metadata.MD, result2 error) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = nil
	fake.headerReturns = struct {
		result1 metadata.MD
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry_ListBundlesClient) HeaderReturnsOnCall(i int, result1 metadata.MD, result2 error) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = nil
	if fake.headerReturnsOnCall == nil {
		fake.headerReturnsOnCall = make(map[int]struct {
			result1 metadata.MD
			result2 error
		})
	}
	fake.headerReturnsOnCall[i] = struct {
		result1 metadata.MD
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry_ListBundlesClient) Recv() (*api.Bundle, error) {
	fake.recvMutex.Lock()
	ret, specificReturn := fake.recvReturnsOnCall[len(fake.recvArgsForCall)]
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct {
	}{})
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if fake.RecvStub != nil {
		return fake.RecvStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.recvReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegistry_ListBundlesClient) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *FakeRegistry_ListBundlesClient) RecvCalls(stub func() (*api.Bundle, error)) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = stub
}

func (fake *FakeRegistry_ListBundlesClient) RecvReturns(result1 *api.Bundle, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry_ListBundlesClient) RecvReturnsOnCall(i int, result1 *api.Bundle, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	if fake.recvReturnsOnCall == nil {
		fake.recvReturnsOnCall = make(map[int]struct {
			result1 *api.Bundle
			result2 error
		})
	}
	fake.recvReturnsOnCall[i] = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry_ListBundlesClient) RecvMsg(arg1 interface{}) error {
	fake.recvMsgMutex.Lock()
	ret, specificReturn := fake.recvMsgReturnsOnCall[len(fake.recvMsgArgsForCall)]
	fake.recvMsgArgsForCall = append(fake.recvMsgArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("RecvMsg", []interface{}{arg1})
	fake.recvMsgMutex.Unlock()
	if fake.RecvMsgStub != nil {
		return fake.RecvMsgStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.recvMsgReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry_ListBundlesClient) RecvMsgCallCount() int {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return len(fake.recvMsgArgsForCall)
}

func (fake *FakeRegistry_ListBundlesClient) RecvMsgCalls(stub func(interface{}) error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = stub
}

func (fake *FakeRegistry_ListBundlesClient) RecvMsgArgsForCall(i int) interface{} {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	argsForCall := fake.recvMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRegistry_ListBundlesClient) RecvMsgReturns(result1 error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = nil
	fake.recvMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) RecvMsgReturnsOnCall(i int, result1 error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = nil
	if fake.recvMsgReturnsOnCall == nil {
		fake.recvMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.recvMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) SendMsg(arg1 interface{}) error {
	fake.sendMsgMutex.Lock()
	ret, specificReturn := fake.sendMsgReturnsOnCall[len(fake.sendMsgArgsForCall)]
	fake.sendMsgArgsForCall = append(fake.sendMsgArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("SendMsg", []interface{}{arg1})
	fake.sendMsgMutex.Unlock()
	if fake.SendMsgStub != nil {
		return fake.SendMsgStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendMsgReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry_ListBundlesClient) SendMsgCallCount() int {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	return len(fake.sendMsgArgsForCall)
}

func (fake *FakeRegistry_ListBundlesClient) SendMsgCalls(stub func(interface{}) error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = stub
}

func (fake *FakeRegistry_ListBundlesClient) SendMsgArgsForCall(i int) interface{} {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	argsForCall := fake.sendMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRegistry_ListBundlesClient) SendMsgReturns(result1 error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = nil
	fake.sendMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) SendMsgReturnsOnCall(i int, result1 error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = nil
	if fake.sendMsgReturnsOnCall == nil {
		fake.sendMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) Trailer() metadata.MD {
	fake.trailerMutex.Lock()
	ret, specificReturn := fake.trailerReturnsOnCall[len(fake.trailerArgsForCall)]
	fake.trailerArgsForCall = append(fake.trailerArgsForCall, struct {
	}{})
	fake.recordInvocation("Trailer", []interface{}{})
	fake.trailerMutex.Unlock()
	if fake.TrailerStub != nil {
		return fake.TrailerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.trailerReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry_ListBundlesClient) TrailerCallCount() int {
	fake.trailerMutex.RLock()
	defer fake.trailerMutex.RUnlock()
	return len(fake.trailerArgsForCall)
}

func (fake *FakeRegistry_ListBundlesClient) TrailerCalls(stub func() metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = stub
}

func (fake *FakeRegistry_ListBundlesClient) TrailerReturns(result1 metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = nil
	fake.trailerReturns = struct {
		result1 metadata.MD
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) TrailerReturnsOnCall(i int, result1 metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = nil
	if fake.trailerReturnsOnCall == nil {
		fake.trailerReturnsOnCall = make(map[int]struct {
			result1 metadata.MD
		})
	}
	fake.trailerReturnsOnCall[i] = struct {
		result1 metadata.MD
	}{result1}
}

func (fake *FakeRegistry_ListBundlesClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	fake.headerMutex.RLock()
	defer fake.headerMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	fake.trailerMutex.RLock()
	defer fake.trailerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistry_ListBundlesClient) recordInvocation(key string, args []interface{}) {
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

var _ api.Registry_ListBundlesClient = new(FakeRegistry_ListBundlesClient)
