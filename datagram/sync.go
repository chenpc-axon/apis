package datagram

import "sync"

var _ Datagram = (*SyncDatagram)(nil)

// SyncDatagram 将其它类型的 Datagram 包装为同步类型的 Datagram
type SyncDatagram struct {
	dg    Datagram
	mLock sync.RWMutex
	dLock sync.RWMutex
}

func NewSyncDatagram(dg Datagram) *SyncDatagram {
	return &SyncDatagram{
		dg: dg,
	}
}

func (s *SyncDatagram) MPut(key string, value interface{}) interface{} {
	s.mLock.Lock()
	defer s.mLock.Unlock()
	return s.dg.MPut(key, value)
}

func (s *SyncDatagram) MPuts(values map[string]interface{}) {
	s.mLock.Lock()
	defer s.mLock.Unlock()
	s.dg.MPuts(values)
}

func (s *SyncDatagram) MRemove(key string) (interface{}, bool) {
	s.mLock.Lock()
	defer s.mLock.Unlock()
	return s.dg.MRemove(key)
}

func (s *SyncDatagram) MGet(key string) (interface{}, bool) {
	s.mLock.RLock()
	defer s.mLock.RUnlock()
	return s.dg.MGet(key)
}

func (s *SyncDatagram) MContains(key string) bool {
	s.mLock.RLock()
	defer s.mLock.RUnlock()
	return s.dg.MContains(key)
}

func (s *SyncDatagram) MKeys() []string {
	s.mLock.RLock()
	defer s.mLock.RUnlock()
	return s.dg.MKeys()
}

func (s *SyncDatagram) MGets(keys ...string) map[string]interface{} {
	s.mLock.RLock()
	defer s.mLock.RUnlock()
	return s.dg.MGets(keys...)
}

func (s *SyncDatagram) MGetAll() map[string]interface{} {
	s.mLock.RLock()
	defer s.mLock.RUnlock()
	return s.dg.MGetAll()
}

func (s *SyncDatagram) Put(key string, value interface{}) interface{} {
	s.dLock.Lock()
	defer s.dLock.Unlock()
	return s.dg.Put(key, value)
}

func (s *SyncDatagram) Puts(values map[string]interface{}) {
	s.dLock.Lock()
	defer s.dLock.Unlock()
	s.dg.Puts(values)
}

func (s *SyncDatagram) Remove(key string) (interface{}, bool) {
	s.dLock.Lock()
	defer s.dLock.Unlock()
	return s.dg.Remove(key)
}

func (s *SyncDatagram) Get(key string) (interface{}, bool) {
	s.dLock.RLock()
	defer s.dLock.RUnlock()
	return s.dg.Get(key)
}

func (s *SyncDatagram) Contains(key string) bool {
	s.dLock.RLock()
	defer s.dLock.RUnlock()
	return s.dg.Contains(key)
}

func (s *SyncDatagram) Keys() []string {
	s.dLock.RLock()
	defer s.dLock.RUnlock()
	return s.dg.Keys()
}

func (s *SyncDatagram) Gets(keys ...string) map[string]interface{} {
	s.dLock.RLock()
	defer s.dLock.RUnlock()
	return s.dg.Gets(keys...)
}

func (s *SyncDatagram) GetAll() map[string]interface{} {
	s.dLock.RLock()
	defer s.dLock.RUnlock()
	return s.dg.GetAll()
}
