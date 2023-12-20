package routes

import (
	"strings"
)

type PathUrlArgs struct {
	ArgRow  string
	ArgCase string
	Arg1    string
	Arg2    string
	Arg3    string
}

// разобрать url и получить  из него структуру с полями
func NewPathUrlArgs(rUrlPath string) *PathUrlArgs {
	var argRow, argCase, arg1, arg2, arg3 = "", "", "", "", ""

	argRow = rUrlPath
	argRow = strings.ReplaceAll(argRow, ":/", "://")
	argRow = strings.ReplaceAll(argRow, `"`, "")
	elems := strings.Split(argRow, "/")
	elems = elems[1:]
	count_elems := len(elems)

	switch count_elems {
	case 1:
		argCase = elems[0]
	case 2:
		argCase = elems[0]
		arg1 = elems[1]
	case 3:
		argCase = elems[0]
		arg1 = elems[1]
		arg2 = elems[2]
	case 4:
		argCase = elems[0]
		arg1 = elems[1]
		arg2 = elems[2]
		arg3 = elems[3]
		// default:
	}

	return &PathUrlArgs{
		ArgRow:  argRow,
		ArgCase: argCase,
		Arg1:    arg1,
		Arg2:    arg2,
		Arg3:    arg3,
	}
}

// получение фильтра из url (ex.: test.com?key1=value1)
func (pua *PathUrlArgs) NewFilterUrlArgs(rUrlPath string) *PathUrlArgs {
	var argRow, argCase = "", ""
	argRow = rUrlPath
	argRow = strings.ReplaceAll(argRow, ":/", "://")
	argRow = strings.ReplaceAll(argRow, `"`, "")
	elems := strings.Split(argRow, "/")
	args := elems[len(elems)-1]
	type_filter := strings.Split(strings.Split(args, "?")[0], "=")
	return &PathUrlArgs{
		ArgRow:  argRow,
		ArgCase: argCase,
		Arg1:    type_filter[1],
	}
}
