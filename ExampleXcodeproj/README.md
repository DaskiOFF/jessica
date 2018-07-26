[![Swift Version 4.1](https://img.shields.io/badge/Swift-4.1-blue.svg?style=flat)](https://developer.apple.com/swift)
[![Recommend xcode version 9.3.1](https://img.shields.io/badge/Xcode-9.3.1-blue.svg?style=flat)](https://developer.apple.com/ios)

**Это сгенерированный файл, для изменения контента редактируйте файл .readme.tpl.md**

# Описание проекта ExampleXcodeproj

# Краткие данные по проекту

## [Dependencies](https://ios-factor.com/dependencies)
Последний раз проект собирался с версией **Xcode 9.3.1** указанной в файле `.xcode-version` ([Подробнее](https://github.com/fastlane/ci/blob/master/docs/xcode-version.md))

Последний раз проект собирался с версией **Swift 4.1** указанной в файле `.swift-version`

### Gemfile
В `Gemfile` описаны зависимости инструментов. Для установки использовать команду `bundle install` ([Подробнее](https://bundler.io/))
```
gem "xcodeproj"
gem "fastlane", ">= 2.96.1", "<= 3.0.0"
gem "cocoapods", "~> 1.5"
```

### Podfile
Зависимости проекта подключены через `cocoapods` и описаны в `Podfile`. Для установки использовать: `[bundle exec] pod install` или `[bundle exec] pod update`
```
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

	pod 'RKKeyboardManager', '~> 0.1'
	pod 'RKTableAdapter', '~> 0.1'
	pod 'RKFoundationExtensions', '~> 0.1'
	pod 'RKUIExtensions', '~> 0.1'
	pod 'RKAutoLayout', '~> 0.1'
end
```

# Структура проекта
- `ExampleXcodeproj` – папка проекта
  - Layers
    - ApplicationLayer
      - Config
    - ServiceLayer
    - DataLayer
      - Entities
    - DomainLayer
      - Entities
    - PresentationLayer
      - Components
      - Flows
        - _AppCoordinator
      - Resources
