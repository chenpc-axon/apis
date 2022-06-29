package expression

import "strings"

// Operator 操作数, 每种操作数代表一种计算方式
type Operator string

// Scope 作用域, Datagram 内的数据分为元数据和普通数据, 可以通过 Scope 指定 Expression 作用的位置和顺序
type Scope string

const (
	// Metadata 只匹配元数据
	Metadata Scope = "Metadata"
	// Data 只匹配普通数据
	Data Scope = "Data"
	// MetadataFirst 优先匹配元数据
	MetadataFirst Scope = "MetadataFirst"
	// DataFirst 优先匹配普通数据
	DataFirst Scope = "DataFirst"
)

// Expression 表达式
type Expression struct {
	// Op 操作数
	Op Operator
	// Scope 作用域
	Scope Scope
	// Key 操作的目标数据项
	Key string
	// Values 匹配值, 有些操作数可以没有值
	Values []string
	// Metadata 自定义元数据, 每种操作数的扩展信息
	Metadata map[string]string
}

// OpEquals 比较两个操作数是否相等
func OpEquals(op1, op2 Operator) bool {
	return strings.Compare(string(op1), string(op2)) == 0
}

// ScopeEquals 比较两个 Scope 是否相等
func ScopeEquals(scope1, scope2 Scope) bool {
	return strings.Compare(string(scope1), string(scope2)) == 0
}
