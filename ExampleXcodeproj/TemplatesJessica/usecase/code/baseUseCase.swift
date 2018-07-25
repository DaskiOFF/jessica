import Foundation
import UseCase

protocol {{ .moduleInfo.name }}BaseUseCase: UseCase {
}

class {{ .moduleInfo.name }}BaseUseCaseImpl: UseCaseBase, {{ .moduleInfo.name }}BaseUseCase {
    var {{ .moduleInfo.nameFirstLower }}Repository: {{ .moduleInfo.name }}Repository
    
    init({{ .moduleInfo.nameFirstLower }}Repository: {{ .moduleInfo.name }}Repository) {
        self.{{ .moduleInfo.nameFirstLower }}Repository = {{ .moduleInfo.nameFirstLower }}Repository
        super.init()
    }
    
    override func cancel() {
        fatalError("not implemented")
    }
 
}
