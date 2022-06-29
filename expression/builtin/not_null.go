package builtin

import (
	"chenpc.com/axon/apis/datagram"
	expr "chenpc.com/axon/apis/expression"
	"fmt"
	"strings"
)

var _ expr.Executor = (*nullExecutor)(nil)

// NotNull 比较 Datagram 中指定 Key 的值是否为 nil, 如果 Key 不存在或为 nil 则判定为 false
const NotNull expr.Operator = "NotNull"

type notNullExecutor struct {
	scope expr.Scope
	key   string
}

func NewNotNullExecutor(expr expr.Expression) (expr.Executor, error) {
	key := strings.TrimSpace(expr.Key)
	if key == "" {
		return nil, fmt.Errorf("the key of NotNull operator cannot be empty")
	}

	return &notNullExecutor{
		scope: expr.Scope,
		key:   key,
	}, nil
}

func (e notNullExecutor) Execute(datagram datagram.Datagram) (bool, error) {
	value, ok := expr.GetData(e.key, e.scope, datagram)
	return ok && value != nil, nil
}
