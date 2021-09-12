package gohttp_mock

import (
	"crypto/md5"
	"encoding/hex"
	"sync"

	"github.com/sebagalan/go-httpclient/core"
)

type mockServer struct {
	enable         bool
	concurrentLock sync.Mutex

	httpClient core.HttpClient

	mocks map[string]*Mock
}

var (
	MockServer mockServer = mockServer{
		mocks:      make(map[string]*Mock),
		httpClient: &httpClientMock{},
	}
)

func (m *mockServer) StartMockSever() {
	m.concurrentLock.Lock()
	defer m.concurrentLock.Unlock()
	m.enable = true
}

func (m *mockServer) StopMockSever() {
	m.concurrentLock.Lock()
	defer m.concurrentLock.Unlock()

	m.enable = false
}

func (m *mockServer) AddMock(mock Mock) {
	m.concurrentLock.Lock()
	defer m.concurrentLock.Unlock()

	key := m.GetMockKey(mock.Method, mock.Url, mock.RequestBody)

	m.mocks[key] = &mock
}

func (m *mockServer) DeleteMocks() {
	m.concurrentLock.Lock()
	defer m.concurrentLock.Unlock()

	m.mocks = make(map[string]*Mock)

}

func (m *mockServer) GetMockKey(method, url, requestBody string) string {

	hash := md5.New()
	hash.Write([]byte(method + url + requestBody))

	key := hex.EncodeToString(hash.Sum(nil))

	return key
}

func (m *mockServer) IsMockServerEnable() bool {
	return m.enable
}

func (m *mockServer) GetMockClient() core.HttpClient {
	return m.httpClient
}
