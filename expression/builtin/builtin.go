package builtin

import (
	registry "chenpc.com/axon/apis/expression"
	"fmt"
)

func init() {
	if err := registry.Register(Equals, NewEqualsExecutor); err != nil {
		panic(fmt.Errorf("register operator %s: %+v", Equals, err))
	}

	if err := registry.Register(NotEquals, NewNotEqualsExecutor); err != nil {
		panic(fmt.Errorf("register operator %s: %+v", NotEquals, err))
	}

	if err := registry.Register(Exist, NewExistExecutor); err != nil {
		panic(fmt.Errorf("register operator %s: %+v", Exist, err))
	}

	if err := registry.Register(NotExist, NewNotExistExecutor); err != nil {
		panic(fmt.Errorf("register operator %s: %+v", NotExist, err))
	}

	if err := registry.Register(In, NewInExecutor); err != nil {
		panic(fmt.Errorf("register operator %s: %+v", In, err))
	}

	if err := registry.Register(NotIn, NewNotInExecutor); err != nil {
		panic(fmt.Errorf("register operator %s: %+v", NotIn, err))
	}

	if err := registry.Register(Null, NewNullExecutor); err != nil {
		panic(fmt.Errorf("register operator %s: %+v", Null, err))
	}

	if err := registry.Register(NotNull, NewNotNullExecutor); err != nil {
		panic(fmt.Errorf("register operator %s: %+v", NotNull, err))
	}
}
