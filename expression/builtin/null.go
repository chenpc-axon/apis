package builtin

import (
	"chenpc.com/axon/apis/datagram"
	expr "chenpc.com/axon/apis/expression"
	"fmt"
	"strings"
)

var _ expr.Executor = (*nullExecutor)(nil)

// Null 比较 Datagram 中指定 Key 的值是否为 nil, 如果 Key 不存在或不为 nil 则判定为 false
const Null expr.Operator = "Null"

type nullExecutor struct {
	scope expr.Scope
	key   string
}

func NewNullExecutor(expr expr.Expression) (expr.Executor, error) {
	key := strings.TrimSpace(expr.Key)
	if key == "" {
		return nil, fmt.Errorf("the key of Null operator cannot be empty")
	}

	return &nullExecutor{
		scope: expr.Scope,
		key:   key,
	}, nil
}

func (e nullExecutor) Execute(datagram datagram.Datagram) (bool, error) {
	value, ok := expr.GetData(e.key, e.scope, datagram)
	return ok && value == nil, nil
}
