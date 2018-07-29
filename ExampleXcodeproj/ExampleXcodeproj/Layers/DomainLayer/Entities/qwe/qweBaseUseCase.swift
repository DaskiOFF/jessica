import Foundation
import UseCase

protocol qweBaseUseCase: UseCase {
}

class qweBaseUseCaseImpl: UseCaseBase, qweBaseUseCase {
    var qweRepository: qweRepository
    // <no value>
    // <no value>
    init(qweRepository: qweRepository) {
        self.qweRepository = qweRepository
        super.init()
    }
    
    override func cancel() {
        fatalError("not implemented")
    }
 
}
