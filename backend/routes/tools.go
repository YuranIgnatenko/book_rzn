package routes

import (
	"fmt"
	"strings"
)

// argRow := r.URL.Path
// 		argRow = strings.ReplaceAll(argRow, ":/", "://")
// 		argRow = strings.ReplaceAll(argRow, `"`, "")
// 		temp := strings.ReplaceAll(argRow, "/delete_table_orders_history/", "")
// 		target_id_order := strings.Split(temp, "/")[0]
// 		rout.ConfirmOFFTableOrders(rout.GetCookieTokenValue(w, r), target_id_order)
// 		http.Redirect(w, r, "/orders_history", http.StatusPermanentRedirect)


func GetFilterUrlArgs(rUrlPath string) *PathUrlArgs {
	var argRow, argCase = "", ""
	argRow = rUrlPath
	argRow = strings.ReplaceAll(argRow, ":/", "://")
	argRow = strings.ReplaceAll(argRow, `"`, "")
	elems := strings.Split(argRow, "/")
	args := elems[len(elems)-1]
	type_filter := strings.Split(strings.Split(args, "?")[0], "=")
	fmt.Println(type_filter)
	return &PathUrlArgs{
		ArgRow:  argRow,
		ArgCase: argCase,
		Arg1:    type_filter[1],
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
