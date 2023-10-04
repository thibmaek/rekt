package android

import (
	"fmt"
	"regexp"
	utils "rekt/utils"
	"strings"
)

func CheckAirshipConfig(assetsDir string) {
	cfgPathGlob := strings.Join([]string{assetsDir, "airshipconfig.*.properties"}, "/")
	fmt.Println("Scanning for Airship configurations...")
	matches, ok := utils.HasFile(cfgPathGlob)
	if ok {
		for _, cfgPath := range matches {
			keywords := regexp.MustCompile(`productionAppSecret`)
			if ok, matches := utils.FindInFile(cfgPath, keywords, nil, nil); ok {
				fmt.Println("🚨 Found Airship config containing a secret:", cfgPath)
				fmt.Println(matches)
			} else {
				fmt.Println("✅ Airship config looks OK", cfgPath)
			}
		}
	} else {
		fmt.Println("✅ No Airship config found")
	}
}
