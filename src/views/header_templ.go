// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.906
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

func Header(munroCount int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<header class=\"bg-gradient-to-r from-blue-700 via-teal-600 to-green-500 text-white shadow-lg z-50 relative\"><div class=\"max-w-7xl mx-auto px-4 sm:px-6 lg:px-8\"><div class=\"flex justify-between items-center py-4\"><div class=\"flex items-center\"><a href=\"/\" class=\"flex items-center hover:opacity-80 transition-opacity duration-200\"><h1 class=\"text-2xl font-bold text-white\">MunroMark</h1><span class=\"ml-3 text-sm text-white/80 hidden sm:inline\">Interactive Map of Scottish Munros</span></a></div><div class=\"flex items-center space-x-6\"><nav class=\"hidden sm:flex space-x-4\"><a href=\"/\" class=\"text-white/80 hover:text-white transition-colors duration-200 text-sm font-medium\">Home</a> <a href=\"/map\" class=\"text-white/80 hover:text-white transition-colors duration-200 text-sm font-medium\">Map</a></nav><div class=\"text-sm text-white/80\" id=\"munro-count\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if munroCount > 0 {
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d Munros Available", munroCount))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `src/views/header.templ`, Line: 30, Col: 55}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "Loading...")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div></div></div></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
