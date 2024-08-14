package nativescript

import (
	"rekt/utils"
	"strings"
)

func IsNativeScriptApp(dir string) bool {
	_, ok := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/libNativeScript.so"}, "/"))
	return ok
}
