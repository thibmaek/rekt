package android

import (
	"fmt"
	"path"
	"regexp"
	utils "rekt/utils"
	"strings"
	"sync"
)

func CheckBuildConfig(sourcesDir string, bundleId string) {
	pkgPath := strings.Split(bundleId, ".")
	cfgPath := path.Join(sourcesDir, path.Join(pkgPath...), "BuildConfig.java")
	fmt.Println("Scanning BuildConfig.java...")
	_, ok := utils.HasFile(cfgPath)
	if ok {
		keywords := regexp.MustCompile(`(?i)(token|secret|apikey|api_key|client_secret|clientsecret|credential|password|license_key|licensekey|access_key|accesskey|private_key|privatekey)`)
		if ok, matches := utils.FindInFile(cfgPath, keywords, nil, nil); ok {
			fmt.Println("🚨 Found BuildConfig.java containing a secret:", cfgPath)
			fmt.Println(matches)
		} else {
			fmt.Println("✅ BuildConfig.java looks OK")
		}
	} else {
		fmt.Println("✅ No BuildConfig.java found")
	}
}

func CheckPrivateKeys(assetsDir string) {
	fmt.Println("Scanning for private key files...")

	var wg sync.WaitGroup
	wg.Add(3)

	key := func(wg *sync.WaitGroup) {
		keyGlob := strings.Join([]string{assetsDir, "*.key"}, "/")
		defer wg.Done()
		matches, ok := utils.HasFile(keyGlob)
		if ok {
			fmt.Println("Found a private key file (Generic):", matches)
		} else {
			fmt.Println("No .key files found")
		}
	}

	pkcs8 := func(wg *sync.WaitGroup) {
		defer wg.Done()
		pkcs8Glob := strings.Join([]string{assetsDir, "*.pkcs8"}, "/")
		matches, ok := utils.HasFile(pkcs8Glob)
		if ok {
			fmt.Println("Found a private key file (PKCS#8):", matches)
		} else {
			fmt.Println("No .pkcs8 files found")
		}
	}

	pkcs12 := func(wg *sync.WaitGroup) {
		defer wg.Done()
		pkcs12Glob := strings.Join([]string{assetsDir, "*.p12"}, "/")
		matches, ok := utils.HasFile(pkcs12Glob)
		if ok {
			fmt.Println("Found a private key file (PKCS#12):", matches)
		} else {
			fmt.Println("No .p12 files found")
		}
	}

	go key(&wg)
	go pkcs8(&wg)
	go pkcs12(&wg)
	wg.Wait()
}
