package analysis

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"rekt/utils"

	"howett.net/plist"
)

// Asserts if a decompiled directory is an iOS one by checking if there is any plist file
func IsIOSApp(dir string) (ok bool, appName string) {
	file, ok := utils.HasFile(fmt.Sprintf("%s/Payload/**/*.plist", dir))
	re := regexp.MustCompile(`Payload/(.*?)/`)
	matched := re.FindStringSubmatch(file[0])
	return ok, matched[1]
}

type InfoPlist struct {
	CFBundleName               string
	CFBundleIdentifier         string
	CFBundleShortVersionString string
	CFBundleVersion            string
	MinimumOSVersion           string
}

func GetIosBundleId(dir string, appName string) (pkgName string, infoPlist InfoPlist) {
	file, err := os.Open(path.Join(dir, fmt.Sprintf("Payload/%s/Info.plist", appName)))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var i InfoPlist
	_, err = plist.Unmarshal(data, &i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i.CFBundleIdentifier)

	return i.CFBundleIdentifier, i
}
