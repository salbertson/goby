package ast

import (
	"bytes"
	"github.com/goby-lang/goby/compiler/token"
)

type BaseNode struct {
	Token token.Token
	isExp bool
}

func (b *BaseNode) IsExp() bool {
	return b.isExp
}

func (b *BaseNode) MarkAsExp() {
	b.isExp = true
}

type node interface {
	TokenLiteral() string
	String() string
	IsExp() bool
	MarkAsExp()
}

type Statement interface {
	node
	statementNode()
}

type Expression interface {
	node
	expressionNode()
}

// Program is the root node of entire AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
