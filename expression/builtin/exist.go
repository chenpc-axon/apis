package builtin

import (
	"chenpc.com/axon/apis/datagram"
	expr "chenpc.com/axon/apis/expression"
	"fmt"
	"strings"
)

var _ expr.Executor = (*existExecutor)(nil)

// Exist 比较 Datagram 中是否存在指定 Key, 如果 Key 存在则会判定为 true, 即使值为 nil 也会判定为 true
const Exist expr.Operator = "Exist"

type existExecutor struct {
	scope expr.Scope
	key   string
}

func NewExistExecutor(expr expr.Expression) (expr.Executor, error) {
	key := strings.TrimSpace(expr.Key)
	if key == "" {
		return nil, fmt.Errorf("the key of Exist operator cannot be empty")
	}

	return &existExecutor{
		scope: expr.Scope,
		key:   key,
	}, nil
}

func (e existExecutor) Execute(datagram datagram.Datagram) (bool, error) {
	_, ok := expr.GetData(e.key, e.scope, datagram)
	return ok, nil
}
