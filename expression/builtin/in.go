package builtin

import (
	"chenpc.com/axon/apis/datagram"
	expr "chenpc.com/axon/apis/expression"
	"fmt"
	"strings"
)

var _ expr.Executor = (*inExecutor)(nil)

// In 比较 Datagram 中指定 Key 的值是否在指定的列表中
// 如果 Key 不存在或 Value 为空会判定为 false
const In expr.Operator = "In"

type inExecutor struct {
	scope  expr.Scope
	key    string
	values map[string]struct{}
}

func NewInExecutor(expr expr.Expression) (expr.Executor, error) {
	if len(expr.Values) == 0 {
		return nil, fmt.Errorf("the values of In operator cannot be empty")
	}

	key := strings.TrimSpace(expr.Key)
	if key == "" {
		return nil, fmt.Errorf("the key of In operator cannot be empty")
	}

	values := make(map[string]struct{}, len(expr.Values))
	for _, value := range expr.Values {
		v := strings.TrimSpace(value)
		if _, ok := values[v]; ok {
			return nil, fmt.Errorf("duplicate values")
		}
		values[v] = struct{}{}
	}

	return &inExecutor{
		scope:  expr.Scope,
		key:    key,
		values: values,
	}, nil
}

func (e inExecutor) Execute(datagram datagram.Datagram) (bool, error) {
	value, ok := expr.GetData(e.key, e.scope, datagram)
	if !ok {
		return false, fmt.Errorf("the key %s is not exist", e.key)
	}

	if value == nil {
		return false, fmt.Errorf("the value of key %s is nil", e.key)
	}

	strValue, err := expr.Convert2String(value)
	if err != nil {
		return false, fmt.Errorf("convert the value [%v] of key %s to string: %+v", value, e.key, err)
	}

	_, ok = e.values[strValue]
	return ok, nil
}
