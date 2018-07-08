package main

import (
	"github.com/daskioff/update_readme_ios/gemfile"
	"github.com/daskioff/update_readme_ios/podfile"
	"github.com/daskioff/update_readme_ios/projectStruct"
	"github.com/daskioff/update_readme_ios/readme"
	"github.com/daskioff/update_readme_ios/versions"
)

func checkFiles() {
	versions.CheckVersionFiles()
	gemfile.Check()
	podfile.Check()
	readme.CheckReadmeTpl()
	projectStruct.Check()
}

func main() {
	checkFiles()

	readme.UpdateREADME()
}
