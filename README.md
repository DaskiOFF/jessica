# Update README for iOS project

# Usage
`go install`

go to iOS project folder and execute `jessica <command> <action> <args>`

## Commands
`jessica help <command>` - Помощь по команде

|Command|Description|
|----|---|
|`hi`|Тестовая команда, которая поприветствует вас|
|`setup`   |Первичная настройка файла конфигурации|
|`readme`   |Создание необходимых файлов и шаблонов для генерации `README.md` файла|
|`struct`   |Конфигурация, создание и описание структуры проекта|
|`generator`   |Генерация файлов для проекта|

# Example
Пример находится в папке `_Example`

# Setup
Первичная конфигурация

Результатом команды являются два файла:
- `~/.jessica.yml` – глобальный файл конфигурации
    - `info_user_name` – Имя пользователя
- `[project_path]/.jessica.yml` – Файл конфигурации проекта
    - `info_company_name` – Имя компании (Для шаблонов)
    - `project_name` – Имя проекта
    - `project_xcodeproj_name` – Имя xcodeproj файла проекта
    - `templates_struct` – Будут ли использоваться шаблоны
    - `structure_use_custom_project_struct` – Используется ли генерация структуры
    - `structure_custom_project_struct` – Структура проекта

# Readme
Поддержка актуальности `README.md` файла

Пока есть недостающие данные – они будут запрашиваться, все последующие вызовы будут просто обновлять файл `README.md` по шаблону

Будут запрошены:
- Версия xcode, с которой проект последний раз собирался
- Версия swift
- Имя проекта (xcodeproj файла)

Будут созданы:
- Файл `Gemfile` с первоначальными зависимостями
- Файл `Podfile` с первоначальными зависимостями
- Файл `.readme.tpl.md` описывающий шаблон резльтирующего файла `README.md`

Переменные используемые при генерации `README.md` из шаблона `.readme.tpl.md`:
- `xcodeVersion` – Версия xcode из файла
- `swiftVersion` – Версия swift из файла
- `gemFileDependencies` – Список зависимостей Gemfile
- `podFileDependencies` – Список зависимостей проекта Podfile
- `projectName` – Имя проекта

При существовании шаблона описывающего структуру проекта `.project_struct.tpl.md`, он подключается в конец файла `.readme.tpl.md`

# Struct
|Action|Description|
|----|---|
|`setup`|Запрос первоначальных настроек|
|`gen`|Генерация|

## setup
Запрашивает первоначальные настройки:
- Будет ли использоваться кастомная структура роекта
- Описывает пример описания этой структуры в файле конфигурации, если она еще не описана, иначе просто выводит описанную в конфиге структуру проекта
- Будут ли использоваться шаблоны

## gen
- Генерация структуры проекта
- Создание шаблона `.project_struct.tpl.md` описания структуры проекта. В шаблоне доступны все теже переменные, что и для `.readme.tpl.md` файла
- Создание папки для шаблонов, если был утвердительный ответ на соответствующий вопрос во время выполнения действия setup

# Generator
|Action|Description|
|----|---|
|`list`|Запрос первоначальных настроек|
|`gen [template name] [module name] [args]`|Генерация|

## list
Находит и выводит список всех доступных шаблонов из папки шаблонов доступных для генерации с помощью действия `gen`

## gen
После указания действия `gen` необходимо указать имя шаблона и имя генерируемого модуля.

Например
```
jessica generator gen repository User --nomock
```

### Описание файла описывающего шаблон
Файл описывающий шаблон должен иметь имя `templates.yml` и находиться в корне папки с файлами шаблона

Доступно 3 секции:
1. `code_files`
1. `test_files`
1. `mock_files`

Каждая секция может содержать список генерируемых файлов. Описание каждого файла содержит:
- `name` – Суффикс генерируемого файла, префиксом будет переданное имя модуля (Если позиция названия модуля не указана явно)
    - `{{.moduleName}}` – Имя модуля, которое было передано при вызове действия
- `template_path` – путь внутри папки шаблона, относительно файла описывающего шаблон
- `output_path` – выходной путь сгенерированного файла, возможно использование переменных
    - `{{.projectName}}` – Имя проекта из файла конфигурации
    - `{{.projectTestsName}}` – Имя папки с тестами проекта из файла конфигурации
    - `{{.moduleName}}` – Имя модуля, которое было передано при вызове действия
- `replace` – Значение true или false, означающее стоит ли перезаписывать генерируемый файл, если файл с таким именем по сохраняемому пути существует, если ключ не указан, то будет запрошено во время выполнения

По умолчанию генерируются все секции, `code_files` является обязательной всегда. Другие можно отключить передав `args`:
- `--notest` – для отключения генерации секции `test_files`
- `--nomock` – для отключения генерации секции `mock_files`

### Описание генерируемого файла
Переменные использовать с помощью конструкции `{{.VariableName}}`. Подробнее, про используемый шаблонизатор можно прочитать [здесь](https://golang.org/pkg/text/template/)

Список доступных переменных

|VariableName|Type|Description|
|----|---|---|
|`fileName`|string|Имя Сгенерированного файла|
|`projectName`|string|Имя проекта для которого генерируется|
|`date`|string|Текущая дата в формате dd.MM.yyyy|
|`year`|int|Текущий год|

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
|`company_name`|string|Имя компании из локального файла конфигурации|