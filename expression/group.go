package expression

import (
	"chenpc.com/axon/apis/datagram"
	"fmt"
	"strings"
)

var _ Executor = (*andGroup)(nil)
var _ Executor = (*orGroup)(nil)

// LogicalOperator 逻辑运算符
type LogicalOperator string

const (
	// And 组内所有的表达式都判定为 true 则返回 true
	And LogicalOperator = "And"
	// Or 组内任一表达式判定为 true 则返回 true
	Or LogicalOperator = "Or"
)

// LogicalOpEquals 判断两个逻辑运算符是否相等
func LogicalOpEquals(op1, op2 LogicalOperator) bool {
	return strings.Compare(string(op1), string(op2)) == 0
}

type andGroup struct {
	// 表达式执行器列表
	executors []Executor
}

type orGroup struct {
	// 表达式执行器列表
	executors []Executor
}

func NewGroup(op LogicalOperator, executors []Executor) (Executor, error) {
	if len(executors) == 0 {
		return nil, fmt.Errorf("the executors of group is empty")
	}

	if LogicalOpEquals(op, And) {
		return &andGroup{
			executors: executors,
		}, nil
	}

	if LogicalOpEquals(op, Or) {
		return &orGroup{
			executors: executors,
		}, nil
	}

	return nil, fmt.Errorf("unknown group logical operator: %s", op)
}

func (g andGroup) Execute(datagram datagram.Datagram) (bool, error) {
	for _, executor := range g.executors {
		ok, err := executor.Execute(datagram)
		if !ok {
			return ok, err
		}
	}
	return true, nil
}

func (g orGroup) Execute(datagram datagram.Datagram) (bool, error) {
	for _, executor := range g.executors {
		ok, err := executor.Execute(datagram)
		if ok {
			return ok, err
		}
	}
	return false, nil
}
