package expression

import (
	"chenpc.com/axon/apis/datagram"
	"fmt"
	"reflect"
)

// Executor 表达式执行器, 每种操作数对应一个执行器
type Executor interface {
	// Execute 判断数据包是否满足表达式定义
	Execute(datagram datagram.Datagram) (bool, error)
}

// GetData 读取 Datagram 中指定 Key 的数据
func GetData(key string, scope Scope, dg datagram.Datagram) (interface{}, bool) {
	if ScopeEquals(scope, Metadata) {
		value, ok := dg.MGet(key)
		return value, ok
	}

	if ScopeEquals(scope, Data) {
		value, ok := dg.Get(key)
		return value, ok
	}

	if ScopeEquals(scope, MetadataFirst) {
		value, ok := dg.MGet(key)
		if !ok {
			return value, true
		}

		value, ok = dg.Get(key)
		return value, ok
	}

	if ScopeEquals(scope, DataFirst) {
		value, ok := dg.Get(key)
		if !ok {
			return value, true
		}

		value, ok = dg.MGet(key)
		return value, ok
	}

	panic(fmt.Errorf("unknown expression scope: %v", scope))
}

// Convert2String 将 value 转为 string
// 支持将类型为 string, []byte, []uint8, 以及实现 fmt.Stringer, fmt.GoStringer 接口类型的数据转换为 string
func Convert2String(value interface{}) (string, error) {
	strValue, ok := value.(string)
	if ok {
		return strValue, nil
	}

	if bytes, ok := value.([]byte); ok {
		return string(bytes), nil
	}

	if bytes, ok := value.([]uint8); ok {
		return string(bytes), nil
	}

	stringer, ok := value.(fmt.Stringer)
	if ok {
		return stringer.String(), nil
	}

	goStringer, ok := value.(fmt.GoStringer)
	if ok {
		return goStringer.GoString(), nil
	}

	return "", fmt.Errorf("unsupported type: %s", reflect.ValueOf(value).Type())
}
