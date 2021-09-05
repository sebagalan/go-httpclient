package gohttp

import (
	"errors"
	"fmt"
	"sync"
)

type MockServer struct {
	enable         bool
	concurrentLock sync.Mutex

	mocks map[string]*Mock
}

var (
	mockServer MockServer = MockServer{
		mocks: make(map[string]*Mock),
	}
)

func StartMockSever() {
	mockServer.concurrentLock.Lock()
	defer mockServer.concurrentLock.Unlock()
	mockServer.enable = true
}

func StopMockSever() {
	mockServer.concurrentLock.Lock()
	defer mockServer.concurrentLock.Unlock()

	mockServer.enable = false
}

func AddMock(mock Mock) {
	mockServer.concurrentLock.Lock()
	defer mockServer.concurrentLock.Unlock()

	key := mockServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)

	mockServer.mocks[key] = &mock
}

func (m *MockServer) getMockKey(method, url, requestBody string) string {
	key := method + url + requestBody
	return key
}

func (m *MockServer) getMock(method, url, requestBody string) *Mock {
	if mockServer.enable {
		key := m.getMockKey(method, url, requestBody)

		mock := mockServer.mocks[key]

		if mock != nil {
			return mock
		}

		return &Mock{
			Error: errors.New(fmt.Sprintf("no mock for %s %s", method, url)),
		}
	}

	return nil
}
