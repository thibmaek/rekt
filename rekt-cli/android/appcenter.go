package android

import (
	"fmt"
	"path"
	"regexp"
	utils "rekt/utils"
)

func CheckAppCenterConfig(assetsDir string) {
	cfgPath := path.Join(assetsDir, "appcenter-config.json")
	fmt.Println("Scanning App Center config...")
	_, ok := utils.HasFile(cfgPath)
	if ok {
		keywords := regexp.MustCompile(`app_secret`)
		if ok, matches := utils.FindInFile(cfgPath, keywords, nil, nil); ok {
			fmt.Println("🚨 Found appcenter-config.json containing a secret:", cfgPath)
			fmt.Println(matches)
		} else {
			fmt.Println("✅ App Center config looks OK")
		}
	} else {
		fmt.Println("✅ No App Center config found")
	}
}
