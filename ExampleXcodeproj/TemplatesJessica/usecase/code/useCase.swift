import Foundation
import UseCase

protocol {{ .moduleInfo.name }}UseCase: {{ .moduleInfo.name }}BaseUseCase {
}

class {{ .moduleInfo.name }}UseCaseImpl: {{ .moduleInfo.name }}BaseUseCaseImpl, {{ .moduleInfo.name }}UseCase {
    
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
