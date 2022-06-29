package expression

import (
	"fmt"
	"sync"
)

var defaultRegistry = new(registry)

// ExecutorCreator 根据 expression 创建对应的 Executor
type ExecutorCreator func(expression Expression) (Executor, error)

// 表达式定义
type definition struct {
	// 操作数
	op Operator
	// Executor 构造器
	creator ExecutorCreator
}

type registry struct {
	defs map[Operator]definition
	lock sync.RWMutex
}

func (r *registry) Register(op Operator, creator ExecutorCreator) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.defs[op]; ok {
		return fmt.Errorf("the operator %s has bean registered", op)
	}

	r.defs[op] = definition{
		op:      op,
		creator: creator,
	}

	return nil
}

func (r *registry) Instantiate(expr Expression) (Executor, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	op := expr.Op
	def, ok := r.defs[op]
	if !ok {
		return nil, fmt.Errorf("no operator %s was registered", op)
	}

	return def.creator(expr)
}

// Register 注册
func Register(op Operator, creator ExecutorCreator) error {
	return defaultRegistry.Register(op, creator)
}

// Instantiate 根据 Expression 创建对应的 Executor
func Instantiate(expr Expression) (Executor, error) {
	return defaultRegistry.Instantiate(expr)
}

// Instantiates 批量创建 Executor
func Instantiates(expressions []Expression) ([]Executor, error) {
	if len(expressions) == 0 {
		return []Executor{}, nil
	}

	executors := make([]Executor, len(expressions))
	for i, expr := range expressions {
		e, err := Instantiate(expr)
		if err != nil {
			return []Executor{}, fmt.Errorf("create executor of %+v: %+v", expr, err)
		}
		executors[i] = e
	}

	return executors, nil
}
