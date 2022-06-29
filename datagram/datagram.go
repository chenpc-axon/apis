package datagram

// Datagram 数据包
type Datagram interface {
	// MPut 向元数据中写入数据, 如果 Key 已存在则返回旧数据
	MPut(key string, value interface{}) interface{}
	// MPuts 批量写入元数据
	MPuts(values map[string]interface{})
	// MRemove 删除元数据中指定 Key, 如果 Key 已存在则返回旧数据和 true, 否则返回 nil 和 false
	MRemove(key string) (interface{}, bool)
	// MGet 获取指定 Key 的元数据
	MGet(key string) (interface{}, bool)
	// MContains 判断元数据包中是否包含指定的 Key
	MContains(key string) bool
	// MKeys 返回元数据包中所有 Key
	MKeys() []string
	// MGets 获取多个 Key 的元数据
	MGets(keys ...string) map[string]interface{}
	// MGetAll 获取所有的元数据
	MGetAll() map[string]interface{}

	// Put 向数据包中写入数据, 如果 Key 已存在则返回旧数据
	Put(key string, value interface{}) interface{}
	// Puts 批量写入数据
	Puts(values map[string]interface{})
	// Remove 删除数据包中指定 Key, 如果 Key 已存在则返回旧数据和 true, 否则返回 nil 和 false
	Remove(key string) (interface{}, bool)
	// Get 获取指定 Key 的数据
	Get(key string) (interface{}, bool)
	// Contains 判断数据包中是否包含指定的 Key
	Contains(key string) bool
	// Keys 返回数据包中所有 Key
	Keys() []string
	// Gets 获取多个 Key 的数据
	Gets(keys ...string) map[string]interface{}
	// GetAll 获取所有的数据
	GetAll() map[string]interface{}
}
