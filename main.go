package main

import (
	"github.com/daskioff/jessica/gemfile"
	"github.com/daskioff/jessica/podfile"
	"github.com/daskioff/jessica/projectStruct"
	"github.com/daskioff/jessica/readme"
	"github.com/daskioff/jessica/versions"
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
