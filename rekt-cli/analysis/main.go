package analysis

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"path"
	"rekt/utils"
	"strings"
)

type xmlApplication struct {
	XMlName     xml.Name `xml:"application"`
	AndroidName string   `xml:"name,attr"`
}

type androidManifest struct {
	XMLName     xml.Name       `xml:"manifest"`
	Package     string         `xml:"package,attr"`
	Application xmlApplication `xml:"application"`
}

func isFlutterApp(dir string) bool {
	_, ok := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/flutter.so"}, "/"))
	return ok
}

func isRNApp(dir string) bool {
	_, ok := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/libhermes.so"}, "/"))
	return ok
}

func isNativeScriptApp(dir string) bool {
	_, ok := utils.HasFile(strings.Join([]string{dir, "resources/lib/**/libNativeScript.so"}, "/"))
	return ok
}

func GetBundleId(dir string) (pkgName string, mainApplication string) {
	file, err := os.Open(path.Join(dir, "resources/AndroidManifest.xml"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var m androidManifest
	err = xml.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}

	mainApplication = strings.TrimSuffix(m.Application.AndroidName, ".MainApplication")
	return m.Package, mainApplication
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
