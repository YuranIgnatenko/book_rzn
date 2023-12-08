package routes

import (
	"strings"
)

// argRow := r.URL.Path
// 		argRow = strings.ReplaceAll(argRow, ":/", "://")
// 		argRow = strings.ReplaceAll(argRow, `"`, "")
// 		temp := strings.ReplaceAll(argRow, "/delete_table_orders_history/", "")
// 		target_id_order := strings.Split(temp, "/")[0]
// 		rout.ConfirmOFFTableOrders(rout.GetCookieTokenValue(w, r), target_id_order)
// 		http.Redirect(w, r, "/orders_history", http.StatusPermanentRedirect)

type PathUrlArgs struct {
	ArgRow  string
	ArgCase string
	Arg1    string
	Arg2    string
	Arg3    string
}

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

// func ParsePathUrl(r *http.Request, path string) string {
// 	argRow := r.URL.Path
// 	argRow = strings.ReplaceAll(argRow, ":/", "://")
// 	argRow = strings.ReplaceAll(argRow, `"`, "")
// 	temp := strings.ReplaceAll(argRow, fmt.Sprintf("/%s/", path), "")
// 	target_hash := strings.Split(temp, "/")[0]
// 	target_count := strings.Split(temp, "/")[1]
// 	target_id_order := strings.Split(temp, "/")[2]
// }
