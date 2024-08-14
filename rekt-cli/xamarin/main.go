package xamarin

import (
	"rekt/utils"
	"strings"
)

func IsXamarinApp(dir string) bool {
	_, ok := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/libxamarin-app.so"}, "/"))
	return ok
}
