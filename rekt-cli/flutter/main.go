package flutter

import (
	"rekt/utils"
	"strings"
)

func IsFlutterApp(dir string) bool {
	_, okAndroid := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/flutter.so"}, "/"))
	_, okIOS := utils.HasFile(strings.Join([]string{dir, "Payload/**/Frameworks/Flutter.framework"}, "/"))
	return okAndroid || okIOS
}
