import Foundation
import UseCase

protocol rrtBaseUseCase: UseCase {
}

class rrtBaseUseCaseImpl: UseCaseBase, rrtBaseUseCase {
    var rrtRepository: rrtRepository
    // <no value>
    // <no value>
    init(rrtRepository: rrtRepository) {
        self.rrtRepository = rrtRepository
        super.init()
    }
    
    override func cancel() {
        fatalError("not implemented")
    }
 
}
