package calculator

// PI 是一个导出的常量，用于测试全局符号
const PI = 3.14159

// Calculator 是一个导出的结构体，用于测试类型信息
type Calculator struct {
	Model string
	count int // 未导出的字段
}

// unexportedFunc 是一个包内私有函数，用于测试符号可见性
func unexportedFunc(c *Calculator) {
	c.count++
}

// Add 是一个简单的导出函数，用于测试常规函数调用
func Add(a, b int) int {
	return a + b
}

// Multiply 是 Calculator 结构体的一个方法，用于测试方法符号
func (c *Calculator) Multiply(a, b int) int {
	unexportedFunc(c) // 调用私有函数
	return a * b
}