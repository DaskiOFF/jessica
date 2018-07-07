package main

func checkFiles() {
	checkVersionFiles()
	checkGemfile()
	checkPodfile()
	checkReadmeTpl()
}

func main() {
	checkFiles()

	updateREADME()
}
