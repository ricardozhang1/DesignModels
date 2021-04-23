package factorymethod

// Operator 是被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口  将
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator接口实现的基类，封装公用方法
type OperatorBase struct {
	a int
	b int
}

// SetA 设置A
func (o *OperatorBase) SetA(a int)  {
	o.a = a
}

// SetB 设置B
func (o *OperatorBase) SetB(b int)  {
	o.b = b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct {}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (p *PlusOperator) Result() int {
	return p.a + p.b
}

// MinusOperatorFactor 是MinusOperator 的工厂类
type MinusOperatorFactor struct {}

func (MinusOperatorFactor) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperator Operator的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (m MinusOperator) Result() int {
	return m.a - m.b
}
