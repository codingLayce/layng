package evaluator

import (
	"layng/object"
	"strconv"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"string": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Boolean:
				if arg.Value {
					return &object.String{Value: "true"}
				}
				return &object.String{Value: "false"}
			case *object.Integer:
				return &object.String{Value: strconv.FormatInt(arg.Value, 10)}
			default:
				return newError("argument to `string` not supported, got %s, want=[BOOLEAN,INTEGER]", args[0].Type())
			}
		},
	},
	"int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				val, err := strconv.ParseInt(arg.Value, 10, 64)
				if err != nil {
					return newError("cannot convert string to int: %s", err.Error())
				}
				return &object.Integer{Value: int64(val)}
			default:
				return newError("argument to `int` not supported, got %s, want=STRING", args[0].Type())
			}
		},
	},
	"bool": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				val, err := strconv.ParseBool(arg.Value)
				if err != nil {
					return newError("cannot convert string to boolean: %s", err.Error())
				}
				return &object.Boolean{Value: val}
			default:
				return newError("argument to `bool` not supported, got %s, want=STRING", args[0].Type())
			}
		},
	},
	"type": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			return &object.String{Value: args[0].Type().String}
		},
	},
}
