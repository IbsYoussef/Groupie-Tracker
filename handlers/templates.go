package handlers

import "html/template"

var Tpl = template.Must(template.ParseGlob("templates/*.html"))
