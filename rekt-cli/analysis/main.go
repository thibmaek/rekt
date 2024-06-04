package analysis

import (
	"rekt/utils"
	"strings"
)

func isFlutterApp(dir string) bool {
	_, okAndroid := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/flutter.so"}, "/"))
	_, okIOS := utils.HasFile(strings.Join([]string{dir, "Payload/**/Frameworks/Flutter.framework"}, "/"))
	return okAndroid || okIOS
}

func isRNApp(dir string) bool {
	_, okAndroid := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/libhermes.so"}, "/"))
	_, okIOS := utils.HasFile(strings.Join([]string{dir, "Payload/**/Frameworks/hermes.framework"}, "/"))
	return okAndroid || okIOS
}

func isNativeScriptApp(dir string) bool {
	_, ok := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/libNativeScript.so"}, "/"))
	return ok
}

func GetAppType(dir string) string {
	if isRNApp(dir) {
		return "React Native"
	}
	if isNativeScriptApp(dir) {
		return "NativeScript"
	}
	if isFlutterApp(dir) {
		return "Flutter"
	}
	return "Android Native (Java / Kotlin)"
}

type BundleExtras struct {
	MainApplication string
	AppVersion      string
}

func GetBundleId(dir string) (id string, extras BundleExtras) {
	isIos, appName := IsIOSApp(dir)
	if isIos {
		bundleId, plist := GetIosBundleId(dir, appName)
		return bundleId, BundleExtras{
			AppVersion: plist.CFBundleVersion,
		}
	}

	bundleId, manifest := GetAndroidBundleId(dir)
	return bundleId, BundleExtras{
		MainApplication: manifest.MainApplication,
	}
}
