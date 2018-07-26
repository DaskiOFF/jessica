import Foundation
import UseCase

protocol {{ .moduleInfo.name }}{{.answers.suffix}}UseCase: {{ .moduleInfo.name }}BaseUseCase {
}

class {{ .moduleInfo.name }}{{.answers.suffix}}UseCaseImpl: {{ .moduleInfo.name }}BaseUseCaseImpl, {{ .moduleInfo.name }}UseCase {
    
     // MARK: - observe
//    override func observe() {
//    }
    
//    override func handleObjects(_ objects: [Any]) -> [Any] {
//        return objects
//    }
    
    // MARK: - execute
//    override func execute() {
//        guard shouldExecute() == true else { return }
//        shouldChangePredicateToObserveLoadedObjects = true
//        shouldAccumulateLoadedObjects = true
//    }
    
//    override func validateParams() throws {
//    }
    
}
