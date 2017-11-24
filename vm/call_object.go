package vm

import (
	"fmt"
	"github.com/goby-lang/goby/compiler/bytecode"
)

type callObject struct {
	method       *MethodObject
	receiverPtr  int
	argCount     int
	argSet       *bytecode.ArgSet
	argStackPtr  int
	lastArgIndex int
	callFrame    *normalCallFrame
}

func newCallObject(receiver Object, method *MethodObject, receiverPtr, argCount int, argSet *bytecode.ArgSet, blockFrame *normalCallFrame, sourceLine int) *callObject {
	cf := newNormalCallFrame(method.instructionSet, method.instructionSet.filename, sourceLine)
	cf.self = receiver
	cf.blockFrame = blockFrame

	return &callObject{
		method:      method,
		receiverPtr: receiverPtr,
		argCount:    argCount,
		argSet:      argSet,
		// This is only for normal/optioned arguments
		lastArgIndex: -1,
		callFrame:    cf,
	}
}

func (co *callObject) instructionSet() *instructionSet {
	return co.method.instructionSet
}

func (co *callObject) paramTypes() []int {
	return co.instructionSet().paramTypes.Types()
}

func (co *callObject) paramNames() []string {
	return co.instructionSet().paramTypes.Names()
}

func (co *callObject) methodName() string {
	return co.method.Name
}

func (co *callObject) argTypes() []int {
	if co.argSet == nil {
		return []int{}
	}

	return co.argSet.Types()
}

func (co *callObject) argPtr() int {
	return co.receiverPtr + 1
}

func (co *callObject) argPosition() int {
	return co.argPtr() + co.argStackPtr
}

func (co *callObject) assignNormalArguments(stack []*Pointer) {
	for i, paramType := range co.paramTypes() {
		if paramType == bytecode.NormalArg {
			index := co.argPosition()
			co.callFrame.insertLCL(i, 0, stack[index].Target)
			co.recordArgStackPtr(index)
		}
	}
}

func (co *callObject) assignNormalAndOptionedArguments(paramIndex int, stack []*Pointer) {
	/*
		Find first usable value as normal argument, for example:

		```ruby
		  def foo(x, y:); end

		  foo(y: 100, 10)
		```

		In the example we can see that 'x' is the first parameter,
		but in the method call it's the second argument.

		This loop is for skipping other types of arguments and get the correct argument index.
	*/
	for argIndex, at := range co.argTypes() {
		if co.lastArgIndex < argIndex && (at == bytecode.NormalArg || at == bytecode.OptionedArg) {
			index := co.argPtr() + argIndex
			co.callFrame.insertLCL(paramIndex, 0, stack[index].Target)
			co.recordArgStackPtr(index)
			// Store latest index value (and compare them to current argument index)
			// This is to make sure we won't get same argument's index twice.
			co.lastArgIndex = argIndex
			break
		} else if co.lastArgIndex < argIndex && at == bytecode.SplatArg {
			index := co.argPtr() + co.argStackPtr
			data := stack[index].Target
			fmt.Printf("Data: %s\n", data.toString())
			co.callFrame.insertLCL(paramIndex, 0, data)
			co.recordArgStackPtr(index)
		}
	}
}

func (co *callObject) assignKeywordArguments(stack []*Pointer) (err error) {
	for argIndex, argType := range co.argTypes() {
		if argType == bytecode.RequiredKeywordArg || argType == bytecode.OptionalKeywordArg {
			argName := co.argSet.Names()[argIndex]
			paramIndex, ok := co.hasKeywordParam(argName)

			if ok {
				index := co.argPtr() + argIndex
				co.callFrame.insertLCL(paramIndex, 0, stack[index].Target)
				co.recordArgStackPtr(index)
			} else {
				err = fmt.Errorf("unknown key %s for method %s", argName, co.methodName())
			}
		}
	}

	return
}

func (co *callObject) assignSplatArgument(stack []*Pointer, arr *ArrayObject) {
	index := len(co.paramTypes()) - 1

	for co.argStackPtr < co.argCount {
		arr.Elements = append(arr.Elements, stack[co.argPosition()].Target)
		co.recordArgStackPtr(co.argPosition())
	}

	co.callFrame.insertLCL(index, 0, arr)
}

func (co *callObject) hasKeywordParam(name string) (index int, result bool) {
	for paramIndex, paramType := range co.paramTypes() {
		paramName := co.paramNames()[paramIndex]

		if paramName == name && (paramType == bytecode.RequiredKeywordArg || paramType == bytecode.OptionalKeywordArg) {
			index = paramIndex
			result = true
			return
		}
	}

	return
}

func (co *callObject) recordArgStackPtr(value int) {
	if value > co.argStackPtr {
		co.argStackPtr++
	}
}

func (co *callObject) hasKeywordArgument(name string) (index int, result bool) {
	for argIndex, argType := range co.argTypes() {
		argName := co.argSet.Names()[argIndex]

		if argName == name && (argType == bytecode.RequiredKeywordArg || argType == bytecode.OptionalKeywordArg) {
			index = argIndex
			result = true
			return
		}
	}

	return
}

func (co *callObject) hasSplatArgument() bool {
	for _, at := range co.argTypes() {
		if at == bytecode.SplatArg {
			return true
		}
	}

	return false
}

func (co *callObject) countParams(paramType int) (n int) {
	for _, at := range co.paramTypes() {
		if at == paramType {
			n++
		}
	}

	return
}

func (co *callObject) countArgs(argType int) (n int) {
	for _, at := range co.argTypes() {
		if at == argType {
			n++
		}
	}

	return
}
