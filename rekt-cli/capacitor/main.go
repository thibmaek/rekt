package capacitor

import (
	"rekt/utils"
	"strings"
)

func IsCapacitorApp(dir string) bool {
	_, okAndroid := utils.HasFile(strings.Join([]string{dir, "resources/assets/capacitor.config.json"}, "/"))
	return okAndroid
}
