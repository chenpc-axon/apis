package builtin

import (
	"chenpc.com/axon/apis/datagram"
	expr "chenpc.com/axon/apis/expression"
	"fmt"
	"strings"
)

var _ expr.Executor = (*notEqualsExecutor)(nil)

// NotEquals 比较 Datagram 中指定 Key 的值是否与设定的值相等
// 如果 Key 不存在或 Value 为空会判定为 false
const NotEquals expr.Operator = "NotEquals"

type notEqualsExecutor struct {
	scope expr.Scope
	key   string
	value string
}

func NewNotEqualsExecutor(expr expr.Expression) (expr.Executor, error) {
	if len(expr.Values) == 0 {
		return nil, fmt.Errorf("the values of NotEquals operator cannot be empty")
	}

	key := strings.TrimSpace(expr.Key)
	value := strings.TrimSpace(expr.Values[0])

	if key == "" {
		return nil, fmt.Errorf("the key of NotEquals operator cannot be empty")
	}

	if value == "" {
		return nil, fmt.Errorf("the first value of NotEquals operator cannot be empty")
	}

	return &notEqualsExecutor{
		scope: expr.Scope,
		key:   key,
		value: value,
	}, nil
}

func (e notEqualsExecutor) Execute(datagram datagram.Datagram) (bool, error) {
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

	return strings.Compare(e.value, strValue) != 0, nil
}
