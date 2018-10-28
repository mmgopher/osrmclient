package config

import (
	"html/template"
	"io/ioutil"
	"os"
	"shipping/logging"
)

var Template *template.Template
func init() {
	Template = template.Must(template.ParseGlob("templates/*.gohtml"))
	logging.InitLogging(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}
