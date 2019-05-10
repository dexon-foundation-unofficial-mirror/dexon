package checker

import (
	"fmt"

	"github.com/dexon-foundation/dexon/core/vm/sqlvm/ast"
	"github.com/dexon-foundation/dexon/core/vm/sqlvm/errors"
	"github.com/dexon-foundation/dexon/core/vm/sqlvm/schema"
)

func checkCastOperator(n *ast.CastOperatorNode,
	s schema.Schema, o CheckOptions, c *schemaCache, el *errors.ErrorList,
	tr schema.TableRef, ta typeAction) ast.ExprNode {

	fn := "CheckCastOperator"
	op := "operator CAST"

	hasError := false

	source := n.SourceExpr
	source = checkExpr(source, s, o, c, el, tr, nil)
	if source != nil {
		n.SourceExpr = source
	} else {
		hasError = true
	}
	target := n.TargetType
	dt, code, message := target.GetType()
	if code == errors.ErrorCodeNil {
		if !dt.ValidColumn() {
			el.Append(errors.Error{
				Position: target.GetPosition(),
				Length:   target.GetLength(),
				Category: errors.ErrorCategorySemantic,
				Code:     errors.ErrorCodeInvalidCastDataType,
				Severity: errors.ErrorSeverityError,
				Prefix:   fn,
				Message: fmt.Sprintf(
					"cannot use %s with type %s", op, dt.String()),
			}, &hasError)
		}
	} else {
		el.Append(errors.Error{
			Position: target.GetPosition(),
			Length:   target.GetLength(),
			Category: errors.ErrorCategorySemantic,
			Code:     code,
			Severity: errors.ErrorSeverityError,
			Prefix:   fn,
			Message:  message,
		}, &hasError)
	}

	if hasError {
		return nil
	}
	r := ast.ExprNode(n)

	return verifyTypeAction(r, fn, dt, el, ta)
}
