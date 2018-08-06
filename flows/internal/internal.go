package internal

import "github.com/daskioff/jessica/configs/keys"

const (
	TemplatesFolderNameDefault    = "TemplatesJessica"
	CustomStructFileNameDefault   = ".project_struct.tpl.md"
	ReadmeTemplateFileNameDefault = ".readme.tpl.md"
	IosProjectFileExtension       = ".xcodeproj"
	GemfileFileName               = "Gemfile"
	PodfileFileName               = "Podfile"
)

func CustomStructDescriptionText() string {
	const exampleStruct = keys.KeyCustomProjectStructDescription + `:
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
  - Support`

	return `Для создания генерируемой структуры вам необходимо описать ее в локальном файле конфигурации .jessica.yml
Описываемая файловая структура будет создаваться внутри папки проекта
	
Например
` + exampleStruct
}
