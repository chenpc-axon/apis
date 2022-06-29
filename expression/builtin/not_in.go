package builtin

import (
	"chenpc.com/axon/apis/datagram"
	expr "chenpc.com/axon/apis/expression"
	"fmt"
	"strings"
)

var _ expr.Executor = (*notInExecutor)(nil)

// NotIn 比较 Datagram 中指定 Key 的值不在指定的列表中
// 如果 Key 不存在或 Value 为空会判定为 false
const NotIn expr.Operator = "NotIn"

type notInExecutor struct {
	scope  expr.Scope
	key    string
	values map[string]struct{}
}

func NewNotInExecutor(expr expr.Expression) (expr.Executor, error) {
	if len(expr.Values) == 0 {
		return nil, fmt.Errorf("the values of NotIn operator cannot be empty")
	}

	key := strings.TrimSpace(expr.Key)
	if key == "" {
		return nil, fmt.Errorf("the key of NotIn operator cannot be empty")
	}

	values := make(map[string]struct{}, len(expr.Values))
	for _, value := range expr.Values {
		v := strings.TrimSpace(value)
		if _, ok := values[v]; ok {
			return nil, fmt.Errorf("duplicate values")
		}
		values[v] = struct{}{}
	}

	return &notInExecutor{
		scope:  expr.Scope,
		key:    key,
		values: values,
	}, nil
}

func (e notInExecutor) Execute(datagram datagram.Datagram) (bool, error) {
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
	return !ok, nil
}
