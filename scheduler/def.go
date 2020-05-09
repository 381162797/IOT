package scheduler

import "sync"

import "net"

import "mserver/models"

var (
	//ID2ConnMap 通过ID得到连接
	ID2ConnMap = connPool{item: make(map[string]net.Conn)}
	//TaskCH 任务传递通道
	TaskCH = make(chan *models.DeviceTask, 5)
	//QueryCH 发送指令通道 pointID为0表示不存数据
	QueryCH = make(chan *models.DeviceTask, 20)
	//StopMap 通过ID停止定时任务
	StopMap = stopCH{item: make(map[string]chan bool)}
)

type stopCH struct {
	mu   sync.RWMutex
	item map[string]chan bool
}
type connPool struct {
	mu   sync.RWMutex
	item map[string]net.Conn
}

type devicePool struct {
	mu   sync.RWMutex
	item map[string]device
}
type device struct {
}

func (m *connPool) Set(k string, v net.Conn) {
	m.mu.Lock()
	m.item[k] = v
	m.mu.Unlock()
}

func (m *connPool) Get(k string) net.Conn {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.item[k]
}
func (m *connPool) Del(k string) {
	m.mu.Lock()
	delete(m.item, k)
	m.mu.Unlock()
}

func (s *stopCH) Set(k string, v chan bool) {
	s.mu.Lock()
	s.item[k] = v
	s.mu.Unlock()
}

func (s *stopCH) Get(k string) chan bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.item[k]
}
func (s *stopCH) Del(k string) {
	s.mu.Lock()
	delete(s.item, k)
	s.mu.Unlock()
}
