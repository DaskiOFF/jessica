import Foundation
import UseCase

protocol User2BaseUseCase: UseCase {
}

class User2BaseUseCaseImpl: UseCaseBase, User2BaseUseCase {
    var user2Repository: User2Repository
    // <no value>
    // <no value>
    init(user2Repository: User2Repository) {
        self.user2Repository = user2Repository
        super.init()
    }
    
    override func cancel() {
        fatalError("not implemented")
    }
 
}
