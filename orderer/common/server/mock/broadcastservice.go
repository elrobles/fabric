// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric/orderer/common/server"
)

type BroadcastService struct {
	BroadcastStub        func(orderer.AtomicBroadcast_BroadcastServer) error
	broadcastMutex       sync.RWMutex
	broadcastArgsForCall []struct {
		arg1 orderer.AtomicBroadcast_BroadcastServer
	}
	broadcastReturns struct {
		result1 error
	}
	broadcastReturnsOnCall map[int]struct {
		result1 error
	}
	DeliverStub        func(orderer.AtomicBroadcast_DeliverServer) error
	deliverMutex       sync.RWMutex
	deliverArgsForCall []struct {
		arg1 orderer.AtomicBroadcast_DeliverServer
	}
	deliverReturns struct {
		result1 error
	}
	deliverReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BroadcastService) Broadcast(arg1 orderer.AtomicBroadcast_BroadcastServer) error {
	fake.broadcastMutex.Lock()
	ret, specificReturn := fake.broadcastReturnsOnCall[len(fake.broadcastArgsForCall)]
	fake.broadcastArgsForCall = append(fake.broadcastArgsForCall, struct {
		arg1 orderer.AtomicBroadcast_BroadcastServer
	}{arg1})
	fake.recordInvocation("Broadcast", []interface{}{arg1})
	fake.broadcastMutex.Unlock()
	if fake.BroadcastStub != nil {
		return fake.BroadcastStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.broadcastReturns
	return fakeReturns.result1
}

func (fake *BroadcastService) BroadcastCallCount() int {
	fake.broadcastMutex.RLock()
	defer fake.broadcastMutex.RUnlock()
	return len(fake.broadcastArgsForCall)
}

func (fake *BroadcastService) BroadcastCalls(stub func(orderer.AtomicBroadcast_BroadcastServer) error) {
	fake.broadcastMutex.Lock()
	defer fake.broadcastMutex.Unlock()
	fake.BroadcastStub = stub
}

func (fake *BroadcastService) BroadcastArgsForCall(i int) orderer.AtomicBroadcast_BroadcastServer {
	fake.broadcastMutex.RLock()
	defer fake.broadcastMutex.RUnlock()
	argsForCall := fake.broadcastArgsForCall[i]
	return argsForCall.arg1
}

func (fake *BroadcastService) BroadcastReturns(result1 error) {
	fake.broadcastMutex.Lock()
	defer fake.broadcastMutex.Unlock()
	fake.BroadcastStub = nil
	fake.broadcastReturns = struct {
		result1 error
	}{result1}
}

func (fake *BroadcastService) BroadcastReturnsOnCall(i int, result1 error) {
	fake.broadcastMutex.Lock()
	defer fake.broadcastMutex.Unlock()
	fake.BroadcastStub = nil
	if fake.broadcastReturnsOnCall == nil {
		fake.broadcastReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.broadcastReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *BroadcastService) Deliver(arg1 orderer.AtomicBroadcast_DeliverServer) error {
	fake.deliverMutex.Lock()
	ret, specificReturn := fake.deliverReturnsOnCall[len(fake.deliverArgsForCall)]
	fake.deliverArgsForCall = append(fake.deliverArgsForCall, struct {
		arg1 orderer.AtomicBroadcast_DeliverServer
	}{arg1})
	fake.recordInvocation("Deliver", []interface{}{arg1})
	fake.deliverMutex.Unlock()
	if fake.DeliverStub != nil {
		return fake.DeliverStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deliverReturns
	return fakeReturns.result1
}

func (fake *BroadcastService) DeliverCallCount() int {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return len(fake.deliverArgsForCall)
}

func (fake *BroadcastService) DeliverCalls(stub func(orderer.AtomicBroadcast_DeliverServer) error) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = stub
}

func (fake *BroadcastService) DeliverArgsForCall(i int) orderer.AtomicBroadcast_DeliverServer {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	argsForCall := fake.deliverArgsForCall[i]
	return argsForCall.arg1
}

func (fake *BroadcastService) DeliverReturns(result1 error) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = nil
	fake.deliverReturns = struct {
		result1 error
	}{result1}
}

func (fake *BroadcastService) DeliverReturnsOnCall(i int, result1 error) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = nil
	if fake.deliverReturnsOnCall == nil {
		fake.deliverReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deliverReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *BroadcastService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.broadcastMutex.RLock()
	defer fake.broadcastMutex.RUnlock()
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BroadcastService) recordInvocation(key string, args []interface{}) {
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

var _ server.BroadcastService = new(BroadcastService)
