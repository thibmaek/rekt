package analysis

import (
	"encoding/xml"
	"fmt"
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

func IsAndroidApp(dir string) bool {
	_, ok := utils.HasFile(fmt.Sprintf("%s/resources/AndroidManifest.xml", dir))
	return ok
}

type Manifest struct {
	MainApplication string
}

func GetAndroidBundleId(dir string) (pkgName string, manifest Manifest) {
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

	mainApplication := strings.TrimSuffix(m.Application.AndroidName, ".MainApplication")
	return m.Package, Manifest{
		MainApplication: mainApplication,
	}
}
