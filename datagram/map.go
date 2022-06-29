package datagram

var _ Datagram = (*MapDatagram)(nil)

// MapDatagram 基于 map 结构的数据包实现
type MapDatagram struct {
	metadata map[string]interface{}
	data     map[string]interface{}
}

func NewMapDatagram() *MapDatagram {
	return &MapDatagram{
		metadata: make(map[string]interface{}, 4),
		data:     make(map[string]interface{}, 8),
	}
}

func (dg *MapDatagram) MPut(key string, value interface{}) interface{} {
	old := dg.metadata[key]
	dg.metadata[key] = value
	return old
}

func (dg *MapDatagram) MPuts(values map[string]interface{}) {
	if len(values) == 0 {
		return
	}

	for key, value := range values {
		dg.metadata[key] = value
	}
}

func (dg *MapDatagram) MRemove(key string) (interface{}, bool) {
	old, ok := dg.metadata[key]
	if ok {
		delete(dg.metadata, key)
	}
	return old, ok
}

func (dg *MapDatagram) MGet(key string) (interface{}, bool) {
	value, ok := dg.metadata[key]
	return value, ok
}

func (dg *MapDatagram) MContains(key string) bool {
	_, ok := dg.metadata[key]
	return ok
}

func (dg *MapDatagram) MKeys() []string {
	if len(dg.metadata) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(dg.metadata))
	for key := range dg.metadata {
		keys = append(keys, key)
	}

	return keys
}

func (dg *MapDatagram) MGets(keys ...string) map[string]interface{} {
	if len(dg.metadata) == 0 || len(keys) == 0 {
		return map[string]interface{}{}
	}

	cloneMetadata := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		cloneMetadata[key] = dg.metadata[key]
	}

	return cloneMetadata
}

func (dg *MapDatagram) MGetAll() map[string]interface{} {
	if len(dg.metadata) == 0 {
		return map[string]interface{}{}
	}

	cloneMetadata := make(map[string]interface{}, len(dg.metadata))
	for key, value := range dg.metadata {
		cloneMetadata[key] = value
	}

	return cloneMetadata
}

func (dg *MapDatagram) Put(key string, value interface{}) interface{} {
	old := dg.data[key]
	dg.data[key] = value
	return old
}

func (dg *MapDatagram) Puts(values map[string]interface{}) {
	if len(values) == 0 {
		return
	}

	for key, value := range values {
		dg.data[key] = value
	}
}

func (dg *MapDatagram) Remove(key string) (interface{}, bool) {
	old, ok := dg.data[key]
	if ok {
		delete(dg.data, key)
	}
	return old, ok
}

func (dg *MapDatagram) Get(key string) (interface{}, bool) {
	value, ok := dg.data[key]
	return value, ok
}

func (dg *MapDatagram) Contains(key string) bool {
	_, ok := dg.data[key]
	return ok
}

func (dg *MapDatagram) Keys() []string {
	if len(dg.data) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(dg.data))
	for key := range dg.data {
		keys = append(keys, key)
	}

	return keys
}

func (dg *MapDatagram) Gets(keys ...string) map[string]interface{} {
	if len(dg.data) == 0 || len(keys) == 0 {
		return map[string]interface{}{}
	}

	cloneData := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		cloneData[key] = dg.data[key]
	}

	return cloneData
}

func (dg *MapDatagram) GetAll() map[string]interface{} {
	if len(dg.data) == 0 {
		return map[string]interface{}{}
	}

	cloneData := make(map[string]interface{}, len(dg.data))
	for key, value := range dg.data {
		cloneData[key] = value
	}

	return cloneData
}
