# Update README for iOS project

# Usage
`go install`

go to iOS project folder and execute `jessica <command>`

## Commands
`jessica help <command>` - Помощь по команде

|Command|Description|
|----|---|
|`hi`|Тестовая команда, которая поприветствует вас|
|`init`   |Первичная настройка файла конфигурации|
|`readme`   |Создание необходимых файлов и шаблонов для генерации README.md файла|



# Result
Create files if not exists
- `.xcode-version`
- `.swift-version`
- `.readme.tpl.md`
- `Gemfile`
- `Podfile`

Create or update
- `README.md`

## Optional
- .project_struct.tpl.md
- project struct folders in Project folder