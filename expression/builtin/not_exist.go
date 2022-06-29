package builtin

import (
	"chenpc.com/axon/apis/datagram"
	expr "chenpc.com/axon/apis/expression"
	"fmt"
	"strings"
)

var _ expr.Executor = (*notExistExecutor)(nil)

// NotExist 比较 Datagram 中是否存在指定 Key, 如果 Key 存在则判定为 false, 即使值为 nil 也会被判定为 false
const NotExist expr.Operator = "NotExist"

type notExistExecutor struct {
	scope expr.Scope
	key   string
	value string
}

func NewNotExistExecutor(expr expr.Expression) (expr.Executor, error) {
	key := strings.TrimSpace(expr.Key)
	if key == "" {
		return nil, fmt.Errorf("the key of NotExist operator cannot be empty")
	}

	return &notExistExecutor{
		scope: expr.Scope,
		key:   key,
	}, nil
}

func (e notExistExecutor) Execute(datagram datagram.Datagram) (bool, error) {
	_, ok := expr.GetData(e.key, e.scope, datagram)
	return ok, nil
}
