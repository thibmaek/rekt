package analysis

import (
	"rekt/flutter"
	"rekt/nativescript"
	"rekt/rn"
	"rekt/xamarin"
)

func GetAppType(dir string) string {
	if rn.IsRNApp(dir) {
		return "React Native"
	}
	if flutter.IsFlutterApp(dir) {
		return "Flutter"
	}
	if xamarin.IsXamarinApp(dir) {
		return "Xamarin"
	}
	if nativescript.IsNativeScriptApp(dir) {
		return "NativeScript"
	}
	return "Android Native (Java / Kotlin)"
}

func GetArchiveType(dir string) string {
	isIOS, _ := IsIOSApp(dir)
	if isIOS {
		return "ios"
	}
	return "android"
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
