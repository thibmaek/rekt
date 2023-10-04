package rn

import (
	"fmt"
	"regexp"
	"rekt/utils"
)

func ScanReactNativeBundle(bundlePath string) {
	fmt.Println("Scanning the React Native bundle...")

	keywords := regexp.MustCompile(`(?i)(secret|apikey|api_key|client_secret|clientsecret)`)
	excludes := regexp.MustCompile(`(?i)(__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED|SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED)`)

	if ok, matches := utils.FindInFile(bundlePath, keywords, excludes, nil); ok {
		fmt.Println("🚨 Found keywords in the bundle:", bundlePath)
		fmt.Println(matches)
	} else {
		fmt.Println("✅ React Native bundle looks OK")
	}
}
