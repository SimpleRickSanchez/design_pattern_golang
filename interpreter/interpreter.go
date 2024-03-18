package interpreter

import (
	"strconv"
	"strings"
)

// Context 提供了执行表达式时需要的环境
type Context struct {
	variables map[string]int
}

// NewContext 创建新的 Context 实例
func NewContext() *Context {
	return &Context{
		variables: make(map[string]int),
	}
}

// SetVariable 设置变量值
func (c *Context) SetVariable(name string, value int) {
	c.variables[name] = value
}

// GetVariable 获取变量值
func (c *Context) GetVariable(name string) int {
	if value, ok := c.variables[name]; ok {
		return value
	}
	return 0
}

// AbstractExpression 是所有表达式的接口
type AbstractExpression interface {
	Interpret(context *Context) int
}

// TerminalExpression 是终结符表达式，例如一个数字或变量
type TerminalExpression struct {
	value string
}

// Interpret 实现终结符表达式的解释
func (e *TerminalExpression) Interpret(context *Context) int {
	if strings.HasPrefix(e.value, "var(") && strings.HasSuffix(e.value, ")") {
		varName := e.value[4 : len(e.value)-1]
		return context.GetVariable(varName)
	}
	res, _ := strconv.Atoi(e.value)
	return res
}

// NonTerminalExpression 是非终结符表达式，例如加法或减法
type NonTerminalExpression struct {
	left, right AbstractExpression
	operator    string
}

// Interpret 实现非终结符表达式的解释
func (e *NonTerminalExpression) Interpret(context *Context) int {
	leftValue := e.left.Interpret(context)
	rightValue := e.right.Interpret(context)

	switch e.operator {
	case "+":
		return leftValue + rightValue
	case "-":
		return leftValue - rightValue
	}

	return 0
}

// ParseExpression 解析表达式字符串并返回对应的抽象表达式
func ParseExpression(expression string) AbstractExpression {
	tokens := strings.Fields(expression)
	if len(tokens) == 1 {
		// 假设 tokens[0] 是一个数字或者变量
		return &TerminalExpression{value: tokens[0]}
	}

	// 假设 tokens[1] 是操作符 "+" 或 "-"
	operator := tokens[1]
	left := ParseExpression(tokens[0])
	right := ParseExpression(strings.Join(tokens[2:], " "))

	return &NonTerminalExpression{left: left, right: right, operator: operator}
}
