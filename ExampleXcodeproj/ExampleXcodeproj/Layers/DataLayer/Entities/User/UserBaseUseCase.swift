import Foundation
import UseCase

protocol UserBaseUseCase: UseCase {
}

class UserBaseUseCaseImpl: UseCaseBase, UserBaseUseCase {
    var userRepository: UserRepository
    // Roman
    // <no value>
    init(userRepository: UserRepository) {
        self.userRepository = userRepository
        super.init()
    }
    
    override func cancel() {
        fatalError("not implemented")
    }
 
}

 
}
