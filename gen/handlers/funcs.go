package handlers

import (
	"strings"
	"text/template"
)

var TplFuncs = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"Title": func(in string) string {
		if in == "" {
			return ""
		}

		return strings.ToTitle(string(in[0])) + in[1:]
	},
	"Replace": func(old, new string, n int, in string) string {
		return strings.Replace(in, old, new, n)
	},
	"ReplaceAll": func(old, new, in string) string {
		return strings.ReplaceAll(in, old, new)
	},
	"TrimSpace": strings.TrimSpace,
	"TrimPrefix": func(prefix, in string) string {
		return strings.TrimPrefix(in, prefix)
	},
	"HasPrefix": func(prefix, in string) bool {
		return strings.HasPrefix(in, prefix)
	},
	"TrimSuffix": func(suffix, in string) string {
		return strings.TrimSuffix(in, suffix)
	},
	"Contains": func(substr, in string) bool {
		return strings.Contains(in, substr)
	},
	"Index": func(substr, in string) int {
		return strings.Index(in, substr)
	},
	"Add": func(l, r int, other ...int) int {
		base := l + r
		for _, v := range other {
			base += v
		}
		return base
	},

	"CCaller":     CCaller,
	"CCallee":     CCallee,
	"GoCaller":    GoCaller,
	"GoCallee":    GoCallee,
	"GoParamName": GoParamName,
	"GoParamType": GoParamType,
	"GoType":      GoType,
	"GoTypeName":  GoTypeName,
	"CgoCaller":   CgoCaller,
	"CgoCallee":   CgoCallee,
}
