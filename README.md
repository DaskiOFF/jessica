# Jessica
Last version: 1.3.5

# Usage
```
brew tap daskioff/jessica
brew install daskioff/jessica/jessica

brew upgrade daskioff/jessica/jessica

brew reinstall daskioff/jessica/jessica
brew uninstall --force daskioff/jessica/jessica
```

Переходим в папку проекта и вызываем `jessica <command> [action] [args]`

## Commands
`jessica help <command>` - Помощь по команде

|Command|Description|
|----|---|
|`hi`|Тестовая команда, которая поприветствует вас и напишет свою версию|
|`version`|Выводит номер текущей версии приложения|
|[setup](#setup)|Настройка конфигурации|
|[readme](#readme-command)|Создание необходимых файлов и шаблонов для генерации `README.md` файла|
|[struct](#struct)|Создание и описание структуры проекта|
|[generator](#generator)|Генерация файлов для проекта|

# Example
Пример находится в папке `ExampleXcodeProj` 

# Setup
Конфигурация. 

Запрашиваются только недостающие поля (для обновления всей конфигурации можно использовать параметры)

|Params|Description|
|----|---|
|`--force, --f`|Полное обновление конфигурации|

## Результатом команды являются два файла
### ~/.jessica.yml – глобальный файл конфигурации
  - `user_name` – Имя пользователя

### [project_path]/.jessica.yml – Файл конфигурации проекта
  - `company_name` – Имя компании (Для шаблонов)
  - `project_type` – Тип проекта [iOS|other]
  - `readme_template_filename` – Имя файла для шаблона README файла
  - `custom_project_struct_use` – Использовать или нет кастомную структуру проекта
  - `custom_project_struct_description` – Описание структуры проекта
  - `custom_project_struct_description_template_filename` – Имя файла с шаблоном описания структуры проекта
  - `templates_use` – Использовать шаблоны или нет
  - `templates_folder_name` – Имя папки, содержащей шаблоны

##### Для проекта типа `iOS`
  - `ios_dependencies_gemfile_use` – Использовать Gemfile или нет
  - `ios_dependencies_podfile_use` – Использовать Podfile или нет
  - `ios_xcodeproj_filename` – Имя xcodeproj файла проекта
  - `ios_target_name_code` – Название таргета кода проекта
  - `ios_folder_name_code` – Имя папки с кодом проекта
  - `ios_target_name_unit_tests` – Название таргета unit тестов проекта
  - `ios_folder_name_unit_tests` – Имя папки unit тестов проекта
  - `ios_target_name_ui_tests` – Название таргета ui тестов проекта
  - `ios_folder_name_ui_tests` – Имя папки ui тестов проекта

##### Для проекта типа `other`
  - `other_project_name` – Название проекта
  - `other_project_folder_name` – Имя папки с кодом проекта

# Readme command
Поддержка актуальности `README.md` файла

Пока есть недостающие данные – они будут запрашиваться, все последующие вызовы будут просто обновлять файл `README.md` по шаблону

Для iOS проекта
  - Будут запрошены:
    - Версия xcode, с которой проект последний раз собирался. Ответ будет сохранен в файл `.xcode-version`
    - Версия swift. Ответ будет сохранен в файл `.swift-version`
  - Будут созданы:
    - Файл `Gemfile` с первоначальными зависимостями
    - Файл `Podfile` с первоначальными зависимостями
    - Файл [readme_template_filename](#setup), описывающий шаблон результирующего файла `README.md`

Переменные, используемые при генерации `README.md` из шаблона [readme_template_filename](#setup):
  - Для всех типов проектов:
    - `projectName` – Имя проекта
  - Для iOS проекта:
    - `xcodeVersion` – Версия xcode из файла
    - `swiftVersion` – Версия swift из файла
    - `gemFileDependencies` – Список зависимостей Gemfile
    - `podFileDependencies` – Список зависимостей проекта Podfile

Если файл описания структуры [templates_folder_name](#setup) существует, он подключается в конец файла [readme_template_filename](#setup)

# Struct
|Action|Description|
|----|---|
|`gen`|Генерация структуры проекта, описанной в конфигурационном файле проекта|
|`example`|Пример описания структуры проекта в конфигурационном файле|

## gen
- Генерация структуры проекта
- Создание шаблона [custom_project_struct_description_template_filename](#setup) описания структуры проекта. В шаблоне доступны все те же переменные, что и для [readme_template_filename](#setup) файла

## example
Выводит пример описания структуры в конфигурационном файле

### Пример описания структуры проекта
```yml
custom_project_struct_description:
  - Data:
    - Entities
  - Domain:
    - Entities
  - Presentation:
    - Resources:
      - Fonts
    - Components:
      - Cells
    - Flows
    - ViewControllers
  - Support
```

# Generator
|Action|Description|
|----|---|
|`list`|Список шаблонов|
|`pull git_url [branch]`|Клонирование репозитория в папку шаблонов|
|`gen TEMPLATE_NAME MODULE_NAME [PARAMS] [CUSTOM_KEYS]`|Генерация|

## list
Находит и выводит список всех доступных шаблонов из папки шаблонов, указанной в файле конфигурации по ключу [templates_folder_name](#setup), доступных для генерации с помощью действия `gen`

## pull
Пример использования
```
jessica generator pull github.com/daskioff/jessica_templates
```
или
```
jessica generator pull github.com/daskioff/jessica_templates my
```
или
```
jessica generator pull https://github.com/daskioff/jessica_templates.git
```

## gen
После указания действия `gen` необходимо указать имя шаблона и имя генерируемого модуля. Далее перечисляются параметры, кастомные ключи и значения, которые доступны в шаблоне по ключу `{{.custom.имя_переданного_ключа}}`

### Пример
```
jessica generator gen repository User --nomock userCusomKey1:Value1 userCustom2:value2
```

### Описание файла, описывающего шаблон
Название шаблона – это имя папки с файлом `templates.yml`, которая находится в общей папке шаблонов проекта [templates_folder_name](#setup).

Структура файла `templates.yml`. Доступно 4 секции:
1. `questions` – optional
1. `code_files` – required
1. `test_files` – optional
1. `mock_files` – optional

#### Questions
Секция содержит вопросы, ответы на которые можно использовать в шаблоне по ключу `{{.answers.KeyName}}`. Формат описания вопроса в файле шаблона:

|Name|Type|Description|
|---|---|---|
|`key`|string|Ключ, по которому будет доступен ответ в генерируемом шаблоне|
|`text`|string|Текст вопроса|
|`required`|bool|Обязательно ли требуется непустой ответ на вопрос|

#### (Code|Test|Mock)_files
Каждая секция, описывающая файлы, может содержать список генерируемых файлов. В значениях можно использовать значения, описанные ниже в разделе "Шаблонные значения". Описание каждого файла содержит:

|Name|Type|Description|
|---|---|---|
|`name`|string|Суффикс генерируемого файла, префиксом будет переданное имя модуля (Если позиция названия модуля не указана явно)|
|`template_path`|string|Путь внутри папки шаблона, относительно файла, описывающего шаблон|
|`output_path`|string|Выходной путь сгенерированного файла, возможно использование переменных|
|`rewrite`|bool|Значение true или false, означающее стоит ли перезаписывать генерируемый файл, если файл с таким именем по сохраняемому пути уже существует. Если ключ не указан, то значение будет запрошено во время выполнения|

По умолчанию генерируются все секции, `code_files` является обязательной всегда. Другие можно отключить передав `params`:
- `--notest` – для отключения генерации секции `test_files`
- `--nomock` – для отключения генерации секции `mock_files`

#### Шаблонные значения
- Все значения по ключу [Custom](#custom)
- Все значения по ключу [Answers](#answers)
- Все значения по ключу [ModuleInfo](#moduleinfo)
- `projectName` – Имя папки проекта из файла конфигурации
- `projectTestsName` – Имя папки с тестами проекта из файла конфигурации (Для проектов типа iOS)
- `projectUITestsName` – Имя папки с ui тестами проекта из файла конфигурации (Для проектов типа iOS)

#### Пример файла, описывающего шаблон
```yml
questions:
  - {key: versionApi,
    text: "Enter API version: ",
    required: true}

  - {key: entryPoint,
    text: "Enter Entry point: ",
    required: true}

  - {key: entityName,
    text: "Enter Entity name: ",
    required: true}

  - {key: suffix,
    text: "Enter suffix for module name: ",
    required: false}

code_files: 
  - {name: BaseUseCase.swift, 
    template_path: code/baseUseCase.swift, 
    output_path: "{{.projectName}}/Layers/DataLayer/Entities/{{.moduleInfo.name}}", 
    rewrite: true}

  - {name: "{{.moduleInfo.name}}{{.answers.suffix}}UseCase.swift", 
    template_path: code/usecase.swift, 
    output_path: "{{.projectName}}/Layers/DataLayer/Entities/{{.moduleInfo.name}}/usecases", 
    rewrite: false}

test_files: 
  - {name: "{{.moduleInfo.name}}{{.answers.suffix}}UseCaseImplTests.swift",
    template_path: tests/useCaseImplTests.swift, 
    output_path: "{{.projectTestsName}}/Layers/DataLayer/Entities/{{.moduleInfo.name}}"}

mock_files:
  - {name: "PartialMock{{.moduleInfo.name}}Repository.swift", 
    template_path: mocks/partialMockUseCaseImpl.swift, 
    output_path: "{{.projectTestsName}}/Mocks/{{.moduleInfo.name}}"}
```

### Описание генерируемого файла
Переменные необходимо использовать с помощью конструкции `{{.VariableName}}`. Подробнее про используемый шаблонизатор можно прочитать [здесь](https://golang.org/pkg/text/template/)

Список доступных переменных, их типы и описания:

|VariableName|Type|Description|
|----|---|---|
|`fileName`|string|Имя сгенерированного файла|
|`projectName`|string|Имя проекта, для которого генерируется|
|`date`|string|Текущая дата в формате dd.MM.yyyy|
|`year`|int|Текущий год|

#### Custom
Использовать `{{.custom.VariableName}}`

Содержит ключи и значения, переданные при запуске, тип значения всегда `string`

_Например:_
```
jessica generator gen usecase User key1:Value1
```

В шаблоне можно будет использовать как `{{.custom.key1}}`

#### Answers
Использовать `{{.answers.VariableName}}`

Содержит ключи, указанные при описании вопроса, и ответы, которые дал пользователь, тип значения всегда `string`

_Например, в шаблоне описать вопрос:_
```
questions:
  - {key: versionApi,
    text: "Enter API version: ",
    required: true}
```

В шаблоне можно будет использовать как `{{.answers.versionApi}}`

#### ModuleInfo
Использовать `{{.moduleInfo.VariableName}}`

|VariableName|Type|Description|
|----|---|---|
|`name`|string|Имя модуля, которое было передано при генерации (Например, UserModule)|
|`nameCapitalize`|string|Имя модуля, которое было передано при генерации, но с первой `Прописной` буквы (Например, UserModule)|
|`nameFirstLower`|string|Имя модуля, которое было передано при генерации, но с первой `строчной` буквы (Например, userModule)|
|`nameUppercase`|string|Имя модуля, которое было передано при генерации, но со всеми `ПРОПИСНЫМИ` буквами (Например, USERMODULE)|
|`nameLowercase`|string|Имя модуля, которое было передано при генерации, но со всеми `строчными` буквами (Например, usermodule)|

#### Developer
Использовать `{{.developer.VariableName}}`

|VariableName|Type|Description|
|----|---|---|
|`name`|string|Имя разработчика из глобального файла конфигурации|
|`companyName`|string|Имя компании из локального файла конфигурации|

# Changelog
### 1.3.5
  - Обновление шаблона .readme.tpl.md
  - Обновление примера
  - Обновил информативность ошибок о неполных конфигах 
  
### 1.3.4
  - Обновлен README.md файл.
  - Команда [setup](#setup) запрашивает только недостающие данные.
  - Добавлен параметр [--force, --f](#setup) к команде [setup](#setup) принудительно перезапрашивающий данные конфигурации.
  - Обновлено описание команд.
  - Исправлены параметры шаблонизатора.

### 1.3.3
  - Исправлены ссылки на разделы в README.md файле.
  - Добавлена команда `version` для получения текущей версии.
  - Установка недостающих зависимостей (bundle, xcodeproj).

### 1.3.2
  - Исправлены ошибки генератора при отсутствии некоторых секций.
  
### 1.3.1
  - Обновлен номер версии.
  - Обновлен README.md файл. Добавлены ссылки на разделы.

### 1.3
  - Добавлено действие `pull` для `generator`. [Подробнее](#pull).
