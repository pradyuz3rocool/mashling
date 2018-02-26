package condition

import (
	"testing"
)

func TestGetConditionOperationAndExpressionType(t *testing.T) {
	tests := []struct {
		test     string
		exprType ExpressionType
		LHS      string
		name     string
		RHS      string
	}{
		{
			test:     "${trigger.content._body[*]?(@._type == \"Element\")._type equals test}",
			exprType: EXPR_TYPE_CONTENT,
			LHS:      "$._body[*]?(@._type == \"Element\")._type+",
			name:     "equals",
			RHS:      "test",
		},
		{
			test:     "${trigger.content._body[*]?(@._type == \"Element\")._type notequals test}",
			exprType: EXPR_TYPE_CONTENT,
			LHS:      "$._body[*]?(@._type == \"Element\")._type+",
			name:     "notequals",
			RHS:      "test",
		},
		{
			test:     "${trigger.content.test == test}",
			exprType: EXPR_TYPE_CONTENT,
			LHS:      "$.test+",
			name:     "==",
			RHS:      "test",
		},
		{
			test:     "${trigger.content.test != test}",
			exprType: EXPR_TYPE_CONTENT,
			LHS:      "$.test+",
			name:     "!=",
			RHS:      "test",
		},
		{
			test:     "${trigger.header.test == test}",
			exprType: EXPR_TYPE_HEADER,
			LHS:      "test",
			name:     "==",
			RHS:      "test",
		},
		{
			test:     "${trigger.header.test != test}",
			exprType: EXPR_TYPE_HEADER,
			LHS:      "test",
			name:     "!=",
			RHS:      "test",
		},
		{
			test:     "${env.test == test}",
			exprType: EXPR_TYPE_ENV,
			LHS:      "test",
			name:     "==",
			RHS:      "test",
		},
		{
			test:     "${env.test != test}",
			exprType: EXPR_TYPE_ENV,
			LHS:      "test",
			name:     "!=",
			RHS:      "test",
		},
	}
	for _, test := range tests {
		condition, exprType, err := GetConditionOperationAndExpressionType(test.test)
		if err != nil {
			t.Fatal(err)
		}
		if exprType != test.exprType {
			t.Fatalf("expression type should be %v", exprType)
		}
		if condition.LHS != test.LHS {
			t.Fatalf("invalid LHS for content expression; should be %v is %v", test.LHS, condition.LHS)
		}
		if operator, exists := OperatorRegistry.Operator(test.name); !exists || condition.Operator != operator {
			t.Fatalf("operator should be %v", test.name)
		}
		if condition.RHS != test.RHS {
			t.Fatalf("invalid RHS for content expression; should be %v is %v", test.RHS, condition.RHS)
		}
	}
}
