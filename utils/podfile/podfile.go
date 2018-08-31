package podfile

import (
	"errors"
	"io/ioutil"
	"regexp"

	"github.com/daskioff/jessica/utils/files"
)

// DefaultFilename Default file name for Podfile
const DefaultFilename = "Podfile"

// ErrPodfileExist means that a Podfile already exists
var ErrPodfileExist = errors.New(DefaultFilename + " already exists")

// CreateDefault create Gemfile with default content for iOS Project
func CreateDefault() error {
	content := `# Uncomment the next line to define a global platform for your project
source 'https://github.com/cocoapods/specs.git'
platform :ios, '9.0'
use_frameworks!

target 'YOUR_TARGET_NAME' do
	# pod 'Fabric'
	# pod 'Crashlytics'
	# pod 'AlamofireImage', '~> 3.3'
	# pod 'Sourcery', '~> 0.14' # https://github.com/krzysztofzablocki/Sourcery
	# pod 'Swinject', '2.4.1' # https://github.com/Swinject/Swinject
	
	# https://github.com/mac-cain13/R.swift
	#
	# RunScript: "$PODS_ROOT/R.swift/rswift" generate "$SRCROOT/$PROJECT_NAME/presentation/resources/r"
	pod 'R.swift', '4.0.0'
	
	# pod 'MBProgressHUD', '1.1.0' # https://github.com/jdg/MBProgressHUD

	# pod 'RKKeyboardManager', '~> 0.1' # https://github.com/DaskiOFF/RKKeyboardManager
	# pod 'RKTableAdapter', '~> 0.1' # https://github.com/DaskiOFF/RKTableAdapter
	# pod 'RKFoundationExtensions', '~> 0.1' # https://github.com/DaskiOFF/RKFoundationExtensions
	# pod 'RKUIExtensions', '~> 0.1' # https://github.com/DaskiOFF/RKUIExtensions
	pod 'RKAutoLayout', '~> 0.1' # https://github.com/DaskiOFF/RKAutoLayout
	pod 'RKCoordinator', '~> 0.1' # https://github.com/DaskiOFF/RKCoordinator
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
	
	installer.pods_project.targets.each do |target|
        target.build_configurations.each do |config|
            config.build_settings['SWIFT_VERSION'] = '4.1'

            if config.name.include?("Debug")
                config.build_settings['SWIFT_OPTIMIZATION_LEVEL'] = '-Onone'
                config.build_settings['GCC_OPTIMIZATION_LEVEL'] = '0'
            end
        end
    end
end`

	fileName := DefaultFilename
	if files.IsFileExist(fileName) {
		return ErrPodfileExist
	}
	files.WriteToFile(fileName, content)

	return nil
}

// Dependencies returns a list of targets with dependencies described in the Podfile
func Dependencies() ([]string, error) {
	var re = regexp.MustCompile(`(?ms)target (.*?)end`)

	filename := DefaultFilename
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}
