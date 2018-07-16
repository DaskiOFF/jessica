package readme

import (
	"io/ioutil"
	"regexp"

	"github.com/daskioff/jessica/utils"
)

const podFileName = "Podfile"

// Read Читает Podfile и выбирает из него список зависимостей для каждого таргета
func readPodfile() ([]string, error) {
	var re = regexp.MustCompile(`(?ms)target (.*?)end`)

	fileContent, err := ioutil.ReadFile(podFileName)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}

// Check Проверяет существование Podfile, если его нет, то его создает и заполняет значением по умолчанию
func checkPodfile() {
	content := `# Uncomment the next line to define a global platform for your project
platform :ios, '9.0'
use_frameworks!

target 'YOUR_TARGET_NAME' do
	# pod 'Fabric'
	# pod 'Crashlytics'
	# pod 'Sourcery'
	# pod 'Swinject', '2.4.1'
	# pod 'AlamofireImage', '~> 3.3'

	# https://github.com/mac-cain13/R.swift
	#
	# RunScript: "$PODS_ROOT/R.swift/rswift" generate "$SRCROOT/$PROJECT_NAME/presentation/resources/r"
	pod 'R.swift', '4.0.0'
	
	# https://github.com/jdg/MBProgressHUD
	pod 'MBProgressHUD', '1.1.0'

	pod 'RKKeyboardManager'
	pod 'RKTableAdapter'
	pod 'RKFoundationExtensions'
	pod 'RKUIExtensions'
	pod 'RKAutoLayout'
end

post_install do |installer|
	Dir.glob(installer.sandbox.target_support_files_root + "Pods-*/*.sh").each do |script|
		flag_name = File.basename(script, ".sh") + "-Installation-Flag"
		folder = "${TARGET_BUILD_DIR}/${UNLOCALIZED_RESOURCES_FOLDER_PATH}"
		file = File.join(folder, flag_name)
		content = File.read(script)
		content.gsub!(/set -e/, "set -e\nKG_FILE=\"#{file}\"\nif [ -f \"$KG_FILE\" ]; then exit 0; fi\nmkdir -p \"#{folder}\"\ntouch \"$KG_FILE\"")
		File.write(script, content)
	end
end`

	fileName := podFileName
	if !utils.IsFileExist(fileName) {
		utils.WriteToFile(fileName, content)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
		utils.PrintlnAttentionMessage("Update YOUR_TARGET_NAME in Podfile")
	}
}
