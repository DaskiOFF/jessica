import Foundation
import UseCase

protocol User2UseCase: User2BaseUseCase {
}

class User2UseCaseImpl: User2BaseUseCaseImpl, User2UseCase {
    
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
