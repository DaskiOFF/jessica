import Foundation

protocol {{ .moduleInfo.name }}{{.answers.suffix}}UseCase: {{ .moduleInfo.name }}BaseUseCase {
}

class {{ .moduleInfo.name }}{{.answers.suffix}}UseCaseImpl: {{ .moduleInfo.name }}BaseUseCaseImpl, {{ .moduleInfo.name }}UseCase {
}
