// Code generated - DO NOT EDIT.

package runtime

import (
	"testing"

	"github.com/dexon-foundation/decimal"

	"github.com/dexon-foundation/dexon/core/vm/sqlvm/ast"
	"github.com/dexon-foundation/dexon/core/vm/sqlvm/errors"
)

var testcasesOpAdd = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: ADD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-2)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}},
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(3)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(4)}},
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(20)}},
				{&Raw{Value: decimal.NewFromFloat(-20)}, &Raw{Value: decimal.NewFromFloat(13)}},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: ADD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(11)}, &Raw{Value: decimal.NewFromFloat(8)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(-9)}, &Raw{Value: decimal.NewFromFloat(-12)}, &Raw{Value: decimal.NewFromFloat(-20)}},
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-10)}},
			},
		),
		nil,
	},
	{
		"Immediate 2",
		Instruction{
			Op: ADD,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(11)}, &Raw{Value: decimal.NewFromFloat(8)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(-9)}, &Raw{Value: decimal.NewFromFloat(-12)}, &Raw{Value: decimal.NewFromFloat(-20)}},
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-10)}},
			},
		),
		nil,
	},
	{
		"Overflow - Immediate",
		Instruction{
			Op: ADD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(127)}},
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow None Immediate",
		Instruction{
			Op: ADD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(126)}},
						{&Raw{Value: decimal.NewFromFloat(126)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow - Immediate",
		Instruction{
			Op: ADD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-128)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow None Immediate",
		Instruction{
			Op: ADD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-127)}},
						{&Raw{Value: decimal.NewFromFloat(-127)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(-2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
}

func (s *instructionSuite) TestOpAdd() {
	s.run(testcasesOpAdd, opAdd)
}

func BenchmarkOpAdd(b *testing.B) {
	runBench(b, testcasesOpAdd, opAdd)
}

var testcasesOpSub = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: SUB,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-2)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}},
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(3)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-4)}},
				{&Raw{Value: decimal.NewFromFloat(20)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(-20)}, &Raw{Value: decimal.NewFromFloat(7)}},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: SUB,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(9)}, &Raw{Value: decimal.NewFromFloat(12)}, &Raw{Value: decimal.NewFromFloat(20)}},
				{&Raw{Value: decimal.NewFromFloat(-11)}, &Raw{Value: decimal.NewFromFloat(-8)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(4)}, &Raw{Value: decimal.NewFromFloat(10)}},
			},
		),
		nil,
	},
	{
		"Immediate 2",
		Instruction{
			Op: SUB,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(-9)}, &Raw{Value: decimal.NewFromFloat(-12)}, &Raw{Value: decimal.NewFromFloat(-20)}},
				{&Raw{Value: decimal.NewFromFloat(11)}, &Raw{Value: decimal.NewFromFloat(8)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-4)}, &Raw{Value: decimal.NewFromFloat(-10)}},
			},
		),
		nil,
	},
	{
		"Overflow - Immediate",
		Instruction{
			Op: SUB,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(127)}},
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow None Immediate",
		Instruction{
			Op: SUB,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(126)}},
						{&Raw{Value: decimal.NewFromFloat(126)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(-2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow - Immediate",
		Instruction{
			Op: SUB,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-128)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow None Immediate",
		Instruction{
			Op: SUB,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-127)}},
						{&Raw{Value: decimal.NewFromFloat(-127)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
}

func (s *instructionSuite) TestOpSub() {
	s.run(testcasesOpSub, opSub)
}

func BenchmarkOpSub(b *testing.B) {
	runBench(b, testcasesOpSub, opSub)
}

var testcasesOpMul = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: MUL,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}},
				{&Raw{Value: decimal.NewFromFloat(4)}, &Raw{Value: decimal.NewFromFloat(-1)}},
				{&Raw{Value: decimal.NewFromFloat(-4)}, &Raw{Value: decimal.NewFromFloat(-100)}},
				{&Raw{Value: decimal.NewFromFloat(100)}, &Raw{Value: decimal.NewFromFloat(100)}},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: MUL,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(-20)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(20)}, &Raw{Value: decimal.NewFromFloat(0)}},
			},
		),
		nil,
	},
	{
		"Immediate - 2",
		Instruction{
			Op: MUL,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(-20)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(20)}, &Raw{Value: decimal.NewFromFloat(0)}},
			},
		),
		nil,
	},
	{
		"Overflow - Immediate",
		Instruction{
			Op: MUL,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(127)}},
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow None Immediate",
		Instruction{
			Op: MUL,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(126)}},
						{&Raw{Value: decimal.NewFromFloat(126)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow - Immediate",
		Instruction{
			Op: MUL,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-128)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow None Immediate",
		Instruction{
			Op: MUL,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-127)}},
						{&Raw{Value: decimal.NewFromFloat(-127)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
}

func (s *instructionSuite) TestOpMul() {
	s.run(testcasesOpMul, opMul)
}

func BenchmarkOpMul(b *testing.B) {
	runBench(b, testcasesOpMul, opMul)
}

var testcasesOpDiv = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: DIV,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}},
				{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
				{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
				{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: DIV,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(13)}, &Raw{Value: decimal.NewFromFloat(13)}},
						{&Raw{Value: decimal.NewFromFloat(-13)}, &Raw{Value: decimal.NewFromFloat(-13)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(-5)}},
				{&Raw{Value: decimal.NewFromFloat(-5)}, &Raw{Value: decimal.NewFromFloat(5)}},
				{&Raw{Value: decimal.NewFromFloat(6)}, &Raw{Value: decimal.NewFromFloat(-6)}},
				{&Raw{Value: decimal.NewFromFloat(-6)}, &Raw{Value: decimal.NewFromFloat(6)}},
			},
		),
		nil,
	},
	{
		"Immediate 2",
		Instruction{
			Op: DIV,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(50)}, &Raw{Value: decimal.NewFromFloat(-50)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(9)}, &Raw{Value: decimal.NewFromFloat(9)}},
						{&Raw{Value: decimal.NewFromFloat(-9)}, &Raw{Value: decimal.NewFromFloat(-9)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(-5)}},
				{&Raw{Value: decimal.NewFromFloat(-5)}, &Raw{Value: decimal.NewFromFloat(5)}},
				{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(-5)}},
				{&Raw{Value: decimal.NewFromFloat(-5)}, &Raw{Value: decimal.NewFromFloat(5)}},
			},
		),
		nil,
	},
	{
		"DivideByZero Immediate",
		Instruction{
			Op: DIV,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeDividedByZero,
	},
	{
		"DivideByZero None Immediate",
		Instruction{
			Op: DIV,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeDividedByZero,
	},
	{
		"Overflow - Immediate",
		Instruction{
			Op: DIV,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(-128)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Overflow None Immediate",
		Instruction{
			Op: DIV,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-128)}},
						{&Raw{Value: decimal.NewFromFloat(-128)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(-2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
}

func (s *instructionSuite) TestOpDiv() {
	s.run(testcasesOpDiv, opDiv)
}

func BenchmarkOpDiv(b *testing.B) {
	runBench(b, testcasesOpDiv, opDiv)
}

var testcasesOpMod = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: MOD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-2)}},
						{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-2)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(3)}, &Raw{Value: decimal.NewFromFloat(3)}},
						{&Raw{Value: decimal.NewFromFloat(-3)}, &Raw{Value: decimal.NewFromFloat(-3)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-2)}},
				{&Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-2)}},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: MOD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(13)}, &Raw{Value: decimal.NewFromFloat(13)}},
						{&Raw{Value: decimal.NewFromFloat(-13)}, &Raw{Value: decimal.NewFromFloat(-13)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(3)}, &Raw{Value: decimal.NewFromFloat(-3)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
				{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
				{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
				{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
				{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
			},
		),
		nil,
	},
	{
		"Immediate - 2",
		Instruction{
			Op: MOD,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(31)}, &Raw{Value: decimal.NewFromFloat(-31)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}, &Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(-10)}, &Raw{Value: decimal.NewFromFloat(-10)}},
						{&Raw{Value: decimal.NewFromFloat(13)}, &Raw{Value: decimal.NewFromFloat(13)}},
						{&Raw{Value: decimal.NewFromFloat(-13)}, &Raw{Value: decimal.NewFromFloat(-13)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
				{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
				{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(-5)}},
				{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(-5)}},
			},
		),
		nil,
	},
	{
		"ModideByZero Immediate",
		Instruction{
			Op: MOD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeDividedByZero,
	},
	{
		"ModideByZero None Immediate",
		Instruction{
			Op: MOD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}},
						{&Raw{Value: decimal.NewFromFloat(10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeDividedByZero,
	},
}

func (s *instructionSuite) TestOpMod() {
	s.run(testcasesOpMod, opMod)
}

func BenchmarkOpMod(b *testing.B) {
	runBench(b, testcasesOpMod, opMod)
}

var testcasesOpConcat = []opTestcase{
	{
		"Concat bytes",
		Instruction{
			Op: CONCAT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abc-1")}, &Raw{Bytes: []byte("xyz-1")}},
						{&Raw{Bytes: []byte("abc-2")}, &Raw{Bytes: []byte("xyz-2")}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("ABC-1")}, &Raw{Bytes: []byte("XYZ-1")}},
						{&Raw{Bytes: []byte("ABC-2")}, &Raw{Bytes: []byte("XYZ-2")}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("abc-1ABC-1")}, &Raw{Bytes: []byte("xyz-1XYZ-1")}},
				{&Raw{Bytes: []byte("abc-2ABC-2")}, &Raw{Bytes: []byte("xyz-2XYZ-2")}},
			},
		),
		nil,
	},
	{
		"Invalid concat",
		Instruction{
			Op: CONCAT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abc-1")}, rawTrue},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("ABC-1")}, rawFalse},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeInvalidDataType,
	},
}

func (s *instructionSuite) TestOpConcat() {
	s.run(testcasesOpConcat, opConcat)
}

func BenchmarkOpConcat(b *testing.B) {
	runBench(b, testcasesOpConcat, opConcat)
}

var testcasesOpNeg = []opTestcase{
	{
		"Neg unary",
		Instruction{
			Op: NEG,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}},
			},
		),
		nil,
	},
	{
		"Overflow Neg",
		Instruction{
			Op: NEG,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-128)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeOverflow,
	},
	{
		"Invalid Neg",
		Instruction{
			Op: NEG,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abc-1")}, rawTrue},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeInvalidDataType,
	},
	{
		"Invalid Neg",
		Instruction{
			Op: NEG,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abc-1")}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeDataLengthNotMatch,
	},
}

func (s *instructionSuite) TestOpNeg() {
	s.run(testcasesOpNeg, opNeg)
}

func BenchmarkOpNeg(b *testing.B) {
	runBench(b, testcasesOpNeg, opNeg)
}

var testcasesOpLt = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: LT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawTrue, rawTrue},
				{rawFalse, rawFalse, rawTrue},
				{rawFalse, rawFalse, rawFalse},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: LT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawFalse, rawTrue},
			},
		),
		nil,
	},
	{
		"Immediate - 2",
		Instruction{
			Op: LT,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawTrue, rawFalse},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpLt() {
	s.run(testcasesOpLt, opLt)
}

func BenchmarkOpLt(b *testing.B) {
	runBench(b, testcasesOpLt, opLt)
}

var testcasesOpGt = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: GT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
						{&Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawFalse, rawFalse},
				{rawTrue, rawFalse, rawFalse},
				{rawTrue, rawTrue, rawFalse},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: GT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawTrue, rawFalse},
			},
		),
		nil,
	},
	{
		"Immediate - 2",
		Instruction{
			Op: GT,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawFalse, rawTrue},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpGt() {
	s.run(testcasesOpGt, opGt)
}

func BenchmarkOpGt(b *testing.B) {
	runBench(b, testcasesOpGt, opGt)
}

var testcasesOpEq = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: EQ,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}},
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawTrue, rawTrue},
				{rawTrue, rawFalse, rawFalse},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: EQ,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(-1)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawTrue, rawTrue},
				{rawTrue, rawFalse, rawFalse},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpEq() {
	s.run(testcasesOpEq, opEq)
}

func BenchmarkOpEq(b *testing.B) {
	runBench(b, testcasesOpEq, opEq)
}

var testcasesOpAnd = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: AND,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawTrue},
						{rawFalse, rawFalse},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawFalse},
				{rawFalse, rawFalse},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: AND,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawTrue},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawFalse},
				{rawFalse, rawTrue},
			},
		),
		nil,
	},
	{
		"Immediate - 2",
		Instruction{
			Op: AND,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawTrue},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawFalse},
				{rawFalse, rawTrue},
			},
		),
		nil,
	},
	{
		"Invalid Data Type",
		Instruction{
			Op: AND,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeInvalidDataType,
	},
}

func (s *instructionSuite) TestOpAnd() {
	s.run(testcasesOpAnd, opAnd)
}

func BenchmarkOpAnd(b *testing.B) {
	runBench(b, testcasesOpAnd, opAnd)
}

var testcasesOpOr = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: OR,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawTrue},
						{rawFalse, rawFalse},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawTrue},
				{rawFalse, rawTrue},
			},
		),
		nil,
	},
	{
		"Immediate",
		Instruction{
			Op: OR,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawTrue},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawTrue},
				{rawTrue, rawTrue},
			},
		),
		nil,
	},
	{
		"Immediate - 2",
		Instruction{
			Op: OR,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawTrue},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawTrue},
				{rawTrue, rawTrue},
			},
		),
		nil,
	},
	{
		"Invalid Data Type",
		Instruction{
			Op: OR,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeInvalidDataType,
	},
}

func (s *instructionSuite) TestOpOr() {
	s.run(testcasesOpOr, opOr)
}

func BenchmarkOpOr(b *testing.B) {
	runBench(b, testcasesOpOr, opOr)
}

var testcasesOpNot = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: NOT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawTrue},
				{rawTrue, rawFalse},
			},
		),
		nil,
	},
	{
		"Errors Invalid Data Type",
		Instruction{
			Op: NOT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeInvalidDataType,
	},
}

func (s *instructionSuite) TestOpNot() {
	s.run(testcasesOpNot, opNot)
}

func BenchmarkOpNot(b *testing.B) {
	runBench(b, testcasesOpNot, opNot)
}

var testcasesOpUnion = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: UNION,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawTrue},
						{rawFalse, rawFalse},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawFalse},
				{rawFalse, rawTrue},
				{rawTrue, rawFalse},
				{rawTrue, rawTrue},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpUnion() {
	s.run(testcasesOpUnion, opUnion)
}

func BenchmarkOpUnion(b *testing.B) {
	runBench(b, testcasesOpUnion, opUnion)
}

var testcasesOpIntxn = []opTestcase{
	{
		"None Immediate",
		Instruction{
			Op: INTXN,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
						{rawTrue, rawTrue},
						{rawFalse, rawFalse},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawTrue},
						{rawFalse, rawFalse},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawFalse, rawFalse},
				{rawTrue, rawTrue},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpIntxn() {
	s.run(testcasesOpIntxn, opIntxn)
}

func BenchmarkOpIntxn(b *testing.B) {
	runBench(b, testcasesOpIntxn, opIntxn)
}

var testcasesOpLike = []opTestcase{
	{
		"Like %\\%b% escape \\",
		Instruction{
			Op: LIKE,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("a%bcdefg")}, &Raw{Bytes: []byte("gfedcba")}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("%\\%b%")}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("\\")}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawFalse},
			},
		),
		nil,
	},
	{
		"Like t1 escape t2",
		Instruction{
			Op: LIKE,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("a%bcdefg")}},
						{&Raw{Bytes: []byte("gfedcba")}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("%\\%b%")}},
						{&Raw{Bytes: []byte("_fed%")}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("\\")}},
						{&Raw{Bytes: []byte("")}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue},
				{rawTrue},
			},
		),
		nil,
	},
	{
		"Like with valid and invalid UTF8",
		Instruction{
			Op: LIKE,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte{226, 40, 161, 228, 189, 160, 229, 165, 189}}, &Raw{Bytes: []byte("gfedcba")}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte{37, 228, 189, 160, 37}}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawFalse},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpLike() {
	s.run(testcasesOpLike, opLike)
}

func BenchmarkOpLike(b *testing.B) {
	runBench(b, testcasesOpLike, opLike)
}

var testcasesOpZip = []opTestcase{
	{
		"Zip two array",
		Instruction{
			Op: ZIP,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
				{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
			},
		),
		nil,
	},
	{
		"Zip immediate",
		Instruction{
			Op: ZIP,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
				{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpZip() {
	s.run(testcasesOpZip, opZip)
}

func BenchmarkOpZip(b *testing.B) {
	runBench(b, testcasesOpZip, opZip)
}

var testcasesOpField = []opTestcase{
	{
		"Retrieve 2nd,3rd column",
		Instruction{
			Op: FIELD,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}},
				{&Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpField() {
	s.run(testcasesOpField, opField)
}

func BenchmarkOpField(b *testing.B) {
	runBench(b, testcasesOpField, opField)
}

var testcasesOpPrune = []opTestcase{
	{
		"Prune 2nd,4th,5th column",
		Instruction{
			Op: PRUNE,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawFalse, rawTrue},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawTrue, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(3)}, &Raw{Value: decimal.NewFromFloat(4)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Value: decimal.NewFromFloat(1)}},
				{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Value: decimal.NewFromFloat(2)}},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpPrune() {
	s.run(testcasesOpPrune, opPrune)
}

func BenchmarkOpPrune(b *testing.B) {
	runBench(b, testcasesOpPrune, opPrune)
}

var testcasesOpCut = []opTestcase{
	{
		"Cut 2nd to 4th columns",
		Instruction{
			Op: CUT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawFalse, rawTrue},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawTrue, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(3)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("abcdefg-1")}, rawTrue},
				{&Raw{Bytes: []byte("abcdefg-2")}, rawFalse},
			},
		),
		nil,
	},
	{
		"Cut 1st column",
		Instruction{
			Op: CUT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawFalse, rawTrue},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawTrue, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawFalse, rawTrue},
				{&Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawTrue, rawFalse},
			},
		),
		nil,
	},
	{
		"Cut since 2nd column",
		Instruction{
			Op: CUT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("abcdefg-1")}},
				{&Raw{Bytes: []byte("abcdefg-2")}},
			},
		),
		nil,
	},
	{
		"Cut all columns",
		Instruction{
			Op: CUT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{
				{},
				{},
			},
		),
		nil,
	},
	{
		"Cut error range - 1",
		Instruction{
			Op: CUT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(5)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeIndexOutOfRange,
	},
	{
		"Cut error range - 2",
		Instruction{
			Op: CUT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("abcdefg-1")}, &Raw{Bytes: []byte("gfedcba-1")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Bytes: []byte("abcdefg-2")}, &Raw{Bytes: []byte("gfedcba-2")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(15)}, &Raw{Value: decimal.NewFromFloat(17)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeIndexOutOfRange,
	},
}

func (s *instructionSuite) TestOpCut() {
	s.run(testcasesOpCut, opCut)
}

func BenchmarkOpCut(b *testing.B) {
	runBench(b, testcasesOpCut, opCut)
}

var testcasesOpFilter = []opTestcase{
	{
		"Filter first 2 rows",
		Instruction{
			Op: FILTER,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue, rawFalse},
						{rawFalse, rawTrue},
						{rawTrue, rawTrue},
						{rawFalse, rawFalse},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue},
						{rawTrue},
						{rawFalse},
						{rawFalse},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawFalse},
				{rawFalse, rawTrue},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpFilter() {
	s.run(testcasesOpFilter, opFilter)
}

func BenchmarkOpFilter(b *testing.B) {
	runBench(b, testcasesOpFilter, opFilter)
}

var testcasesOpCast = []opTestcase{
	{
		"None Immediate - int",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 1), ast.ComposeDataType(ast.DataTypeMajorInt, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(127)}, &Raw{Value: decimal.NewFromFloat(127)}},
						{&Raw{Value: decimal.NewFromFloat(-128)}, &Raw{Value: decimal.NewFromFloat(-128)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 2),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 2),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(127)}, &Raw{Value: decimal.NewFromFloat(127)}},
				{&Raw{Value: decimal.NewFromFloat(-128)}, &Raw{Value: decimal.NewFromFloat(-128)}},
			},
		),
		nil,
	},
	{
		"None Immediate - int2",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 1), ast.ComposeDataType(ast.DataTypeMajorInt, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Value: decimal.NewFromFloat(-32768)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorUint, 1), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Value: decimal.NewFromFloat(32768)}},
			},
		),
		nil,
	},
	{
		"None Immediate - int3",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 1), ast.ComposeDataType(ast.DataTypeMajorInt, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Value: decimal.NewFromFloat(-32768)}},
						{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawTrue},
				{rawFalse, rawFalse},
			},
		),
		nil,
	},
	{
		"None Immediate - int4",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 1), ast.ComposeDataType(ast.DataTypeMajorInt, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Value: decimal.NewFromFloat(-32768)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1), ast.ComposeDataType(ast.DataTypeMajorAddress, 0),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1), ast.ComposeDataType(ast.DataTypeMajorAddress, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte{0x7f, 0xff}}, &Raw{Bytes: []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0x80, 0x00}}},
			},
		),
		nil,
	},
	{
		"None Immediate - uint",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(128)}, &Raw{Value: decimal.NewFromFloat(128)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 2),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 2),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(128)}, &Raw{Value: decimal.NewFromFloat(128)}},
			},
		),
		nil,
	},
	{
		"None Immediate - uint2",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Value: decimal.NewFromFloat(32768)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 1), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 1), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Bytes: []byte{0x80, 0x00}}},
			},
		),
		nil,
	},
	{
		"None Immediate - uint3",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue, rawFalse},
			},
		),
		nil,
	},
	{
		"None Immediate - uint4",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1),
			},
			[]Tuple{
				{&Raw{Bytes: []byte{0x7f, 0xff}}, &Raw{Bytes: []byte{0x00, 0x00}}},
			},
		),
		nil,
	},
	{
		"None Immediate - uint5",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(32767)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorAddress, 1),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorAddress, 1),
			},
			[]Tuple{
				{&Raw{Bytes: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x7f, 0xff}}},
			},
		),
		nil,
	},
	{
		"None Immediate - bytes",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1),
					},
					[]Tuple{
						{&Raw{Bytes: []byte{0xff, 0xff}}, &Raw{Bytes: []byte{0xff, 0xff}}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 2),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 2),
			},
			[]Tuple{
				{&Raw{Bytes: []byte{0xff}}, &Raw{Bytes: []byte{0xff, 0xff, 0x00}}},
			},
		),
		nil,
	},
	{
		"None Immediate - bytes2",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1),
					},
					[]Tuple{
						{&Raw{Bytes: []byte{0x7f, 0xff}}, &Raw{Bytes: []byte{0x80, 0x00}}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 1), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 1), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(32767)}, &Raw{Value: decimal.NewFromFloat(32768)}},
			},
		),
		nil,
	},
	{
		"None Immediate - bytes3",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 1),
					},
					[]Tuple{
						{&Raw{Bytes: []byte{0x7f, 0xff}}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 1),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 1),
			},
			[]Tuple{
				{&Raw{Bytes: []byte{0x7f, 0xff}}},
			},
		),
		nil,
	},
	{
		"Same type",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{rawTrue},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{rawTrue},
			},
		),
		nil,
	},
	{
		"Error Invalid Type",
		Instruction{
			Op: CAST,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 2),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(-32768)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{},
			[]Tuple{},
		),
		errors.ErrorCodeInvalidCastType,
	},
}

func (s *instructionSuite) TestOpCast() {
	s.run(testcasesOpCast, opCast)
}

func BenchmarkOpCast(b *testing.B) {
	runBench(b, testcasesOpCast, opCast)
}

var testcasesOpSort = []opTestcase{
	{
		"Multi-column sorting",
		Instruction{
			Op: SORT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("c")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(2)}, rawTrue},
						{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(3)}, rawTrue},
						{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(1)}, rawFalse},
						{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
						{&Raw{Bytes: []byte("c")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
						{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
						{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{rawFalse, &Raw{Value: decimal.NewFromFloat(1)}},
						{rawTrue, &Raw{Value: decimal.NewFromFloat(2)}},
						{rawFalse, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("c")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
				{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
				{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
				{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(3)}, rawTrue},
				{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
				{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(2)}, rawTrue},
				{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(1)}, rawFalse},
				{&Raw{Bytes: []byte("c")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
			},
		),
		nil,
	},
	{
		"Multi-column sorting - 2",
		Instruction{
			Op: SORT,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte("c")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
						{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(2)}, rawTrue},
						{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(3)}, rawTrue},
						{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(1)}, rawFalse},
						{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
						{&Raw{Bytes: []byte("c")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
						{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
						{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorBool, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{rawTrue, &Raw{Value: decimal.NewFromFloat(0)}},
						{rawTrue, &Raw{Value: decimal.NewFromFloat(1)}},
						{rawFalse, &Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorBool, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(1)}, rawFalse},
				{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(3)}, rawTrue},
				{&Raw{Bytes: []byte("a")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
				{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(2)}, rawTrue},
				{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(2)}, rawFalse},
				{&Raw{Bytes: []byte("b")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
				{&Raw{Bytes: []byte("c")}, &Raw{Value: decimal.NewFromFloat(1)}, rawTrue},
				{&Raw{Bytes: []byte("c")}, &Raw{Value: decimal.NewFromFloat(3)}, rawFalse},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpSort() {
	s.run(testcasesOpSort, opSort)
}

func BenchmarkOpSort(b *testing.B) {
	runBench(b, testcasesOpSort, opSort)
}

var testcasesOpRange = []opTestcase{
	{
		"Range test limit 2 offset 1",
		Instruction{
			Op: RANGE,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}},
						{&Raw{Value: decimal.NewFromFloat(3)}},
						{&Raw{Value: decimal.NewFromFloat(4)}},
						{&Raw{Value: decimal.NewFromFloat(5)}},
						{&Raw{Value: decimal.NewFromFloat(6)}},
						{&Raw{Value: decimal.NewFromFloat(7)}},
						{&Raw{Value: decimal.NewFromFloat(8)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 7), ast.ComposeDataType(ast.DataTypeMajorUint, 7),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(2)}},
				{&Raw{Value: decimal.NewFromFloat(3)}},
			},
		),
		nil,
	},
	{
		"Range test limit 0 offset 1",
		Instruction{
			Op: RANGE,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 7), ast.ComposeDataType(ast.DataTypeMajorUint, 7),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(0)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{},
		),
		nil,
	},
	{
		"Range test offset 20",
		Instruction{
			Op: RANGE,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 7),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(20)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{},
		),
		nil,
	},
	{
		"Range test limit 10 offset 20",
		Instruction{
			Op: RANGE,
			Input: []*Operand{
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorInt, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}},
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 7), ast.ComposeDataType(ast.DataTypeMajorUint, 7),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(20)}, &Raw{Value: decimal.NewFromFloat(10)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorInt, 0),
			},
			[]Tuple{},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpRange() {
	s.run(testcasesOpRange, opRange)
}

func BenchmarkOpRange(b *testing.B) {
	runBench(b, testcasesOpRange, opRange)
}

var testcasesOpFuncBitAnd = []opTestcase{
	{
		"Func BitAnd",
		Instruction{
			Op: FUNC,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(10)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-128)}, &Raw{Bytes: []byte{0x12, 0x34}}, &Raw{Bytes: []byte{0x56, 0x78}}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(6)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Bytes: []byte{0xff, 0xff}}, &Raw{Bytes: []byte{0x00, 0x00}}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-128)}, &Raw{Bytes: []byte{0x12, 0x34}}, &Raw{Bytes: []byte{0x00, 0x00}}},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpFuncBitAnd() {
	s.run(testcasesOpFuncBitAnd, opFunc)
}

func BenchmarkOpFuncBitAnd(b *testing.B) {
	runBench(b, testcasesOpFuncBitAnd, opFunc)
}

var testcasesOpFuncBitOr = []opTestcase{
	{
		"Func BitOr",
		Instruction{
			Op: FUNC,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(11)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-128)}, &Raw{Bytes: []byte{0x12, 0x34}}, &Raw{Bytes: []byte{0x56, 0x78}}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(6)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Bytes: []byte{0xff, 0xff}}, &Raw{Bytes: []byte{0x00, 0x00}}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(6)}, &Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Bytes: []byte{0xff, 0xff}}, &Raw{Bytes: []byte{0x56, 0x78}}},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpFuncBitOr() {
	s.run(testcasesOpFuncBitOr, opFunc)
}

func BenchmarkOpFuncBitOr(b *testing.B) {
	runBench(b, testcasesOpFuncBitOr, opFunc)
}

var testcasesOpFuncBitXor = []opTestcase{
	{
		"Func BitXor",
		Instruction{
			Op: FUNC,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(12)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-128)}, &Raw{Bytes: []byte{0x12, 0x34}}, &Raw{Bytes: []byte{0x56, 0x78}}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(5)}, &Raw{Value: decimal.NewFromFloat(6)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Value: decimal.NewFromFloat(-2)}, &Raw{Bytes: []byte{0xff, 0xff}}, &Raw{Bytes: []byte{0x00, 0x00}}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(4)}, &Raw{Value: decimal.NewFromFloat(4)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(126)}, &Raw{Bytes: []byte{0xed, 0xcb}}, &Raw{Bytes: []byte{0x56, 0x78}}},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpFuncBitXor() {
	s.run(testcasesOpFuncBitXor, opFunc)
}

func BenchmarkOpFuncBitXor(b *testing.B) {
	runBench(b, testcasesOpFuncBitXor, opFunc)
}

var testcasesOpFuncBitNot = []opTestcase{
	{
		"Func BitNot",
		Instruction{
			Op: FUNC,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(13)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(128)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(-1)}, &Raw{Value: decimal.NewFromFloat(-128)}, &Raw{Bytes: []byte{0x12, 0x34}}, &Raw{Bytes: []byte{0xff, 0x00}}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorUint, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorInt, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0), ast.ComposeDataType(ast.DataTypeMajorFixedBytes, 0),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(127)}, &Raw{Value: decimal.NewFromFloat(255)}, &Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(127)}, &Raw{Bytes: []byte{0xed, 0xcb}}, &Raw{Bytes: []byte{0x00, 0xff}}},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpFuncBitNot() {
	s.run(testcasesOpFuncBitNot, opFunc)
}

func BenchmarkOpFuncBitNot(b *testing.B) {
	runBench(b, testcasesOpFuncBitNot, opFunc)
}

var testcasesOpFuncOctetLength = []opTestcase{
	{
		"Func octet length",
		Instruction{
			Op: FUNC,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(14)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0), ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte{}}, &Raw{Bytes: []byte{1}}, &Raw{Bytes: []byte{1, 2}}, &Raw{Bytes: []byte{1, 2, 3}}, &Raw{Bytes: []byte{1, 2, 3, 4}}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorUint, 32), ast.ComposeDataType(ast.DataTypeMajorUint, 32), ast.ComposeDataType(ast.DataTypeMajorUint, 32), ast.ComposeDataType(ast.DataTypeMajorUint, 32), ast.ComposeDataType(ast.DataTypeMajorUint, 32),
			},
			[]Tuple{
				{&Raw{Value: decimal.NewFromFloat(0)}, &Raw{Value: decimal.NewFromFloat(1)}, &Raw{Value: decimal.NewFromFloat(2)}, &Raw{Value: decimal.NewFromFloat(3)}, &Raw{Value: decimal.NewFromFloat(4)}},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpFuncOctetLength() {
	s.run(testcasesOpFuncOctetLength, opFunc)
}

func BenchmarkOpFuncOctetLength(b *testing.B) {
	runBench(b, testcasesOpFuncOctetLength, opFunc)
}

var testcasesOpFuncSubString = []opTestcase{
	{
		"Func sub string",
		Instruction{
			Op: FUNC,
			Input: []*Operand{
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 0),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 1),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(15)}},
					},
				),
				makeOperand(
					false,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
					},
					[]Tuple{
						{&Raw{Bytes: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 7),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(2)}},
					},
				),
				makeOperand(
					true,
					[]ast.DataType{
						ast.ComposeDataType(ast.DataTypeMajorUint, 7),
					},
					[]Tuple{
						{&Raw{Value: decimal.NewFromFloat(5)}},
					},
				),
			},
			Output: 0,
		},
		makeOperand(
			false,
			[]ast.DataType{
				ast.ComposeDataType(ast.DataTypeMajorDynamicBytes, 0),
			},
			[]Tuple{
				{&Raw{Bytes: []byte{3, 4, 5, 6, 7}}},
			},
		),
		nil,
	},
}

func (s *instructionSuite) TestOpFuncSubString() {
	s.run(testcasesOpFuncSubString, opFunc)
}

func BenchmarkOpFuncSubString(b *testing.B) {
	runBench(b, testcasesOpFuncSubString, opFunc)
}
