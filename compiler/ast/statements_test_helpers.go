//+build !release

package ast

import (
	"testing"
)

/*
 BaseNode
*/

// IsBeginStmt fails the test and returns nil by default
func (b *BaseNode) IsBeginStmt(t *testing.T) (tbs *TestableBeginStatement) {
	t.Helper()
	t.Fatalf(nodeFailureMsgFormat, "begin statement", b)
	return nil
}

// IsClassStmt fails the test and returns nil by default
func (b *BaseNode) IsClassStmt(t *testing.T) *TestableClassStatement {
	t.Helper()
	t.Fatalf(nodeFailureMsgFormat, "class statement", b)
	return nil
}

// IsModuleStmt fails the test and returns nil by default
func (b *BaseNode) IsModuleStmt(t *testing.T) *TestableModuleStatement {
	t.Helper()
	t.Fatalf(nodeFailureMsgFormat, "module statement", b)
	return nil
}

// IsReturnStmt fails the test and returns nil by default
func (b *BaseNode) IsReturnStmt(t *testing.T) *TestableReturnStatement {
	t.Helper()
	t.Fatalf(nodeFailureMsgFormat, "return statement", b)
	return nil
}

// IsRescueStmt fails the test and returns nil by default
func (b *BaseNode) IsRescueStmt(t *testing.T) *TestableRescueStatement {
	t.Helper()
	t.Fatalf(nodeFailureMsgFormat, "rescue statement", b)
	return nil
}

// IsDefStmt fails the test and returns nil by default
func (b *BaseNode) IsDefStmt(t *testing.T) *TestableDefStatement {
	t.Helper()
	t.Fatalf(nodeFailureMsgFormat, "method definition", b)
	return nil
}

// IsWhileStmt fails the test and returns nil by default
func (b *BaseNode) IsWhileStmt(t *testing.T) (ws *TestableWhileStatement) {
	t.Helper()
	t.Fatalf(nodeFailureMsgFormat, "while statement", b)
	return nil
}

// IsBeginStmt returns a pointer of the begin statement
func (bs *BeginStatement) IsBeginStmt(t *testing.T) *TestableBeginStatement {
	return &TestableBeginStatement{t: t, BeginStatement: bs}
}

// IsClassStmt returns a pointer of the class statement
func (cs *ClassStatement) IsClassStmt(t *testing.T) *TestableClassStatement {
	return &TestableClassStatement{t: t, ClassStatement: cs}
}

// IsModuleStmt returns a pointer of the module statement
func (ms *ModuleStatement) IsModuleStmt(t *testing.T) *TestableModuleStatement {
	return &TestableModuleStatement{ModuleStatement: ms, t: t}
}

// IsDefStmt returns a pointer of the DefStatement
func (ds *DefStatement) IsDefStmt(t *testing.T) *TestableDefStatement {
	return &TestableDefStatement{DefStatement: ds, t: t}
}

// IsRescueStmt returns a pointer of the RescueStatement
func (rs *RescueStatement) IsRescueStmt(t *testing.T) (trs *TestableRescueStatement) {
	return &TestableRescueStatement{t: t, RescueStatement: rs}
}

// IsDefStmt returns a pointer of the ReturnStatement
func (rs *ReturnStatement) IsReturnStmt(t *testing.T) (trs *TestableReturnStatement) {
	return &TestableReturnStatement{t: t, ReturnStatement: rs}
}

// IsExpressionStmt returns ExpressionStatement itself
func (ts *ExpressionStatement) IsExpression(t *testing.T) TestableExpression {
	return ts.Expression.(TestableExpression)
}

// IsWhileStmt returns the pointer of current while statement
func (ws *WhileStatement) IsWhileStmt(t *testing.T) *TestableWhileStatement {
	return &TestableWhileStatement{WhileStatement: ws, t: t}
}
