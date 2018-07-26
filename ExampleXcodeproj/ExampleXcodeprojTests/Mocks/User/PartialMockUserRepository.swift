import Foundation
import XCTest
import Mirage

@testable import ExampleXcodeproj
import UseCase

class PartialMockUsertUseCaseImpl: UsertUseCaseImpl, PartialMock {
    
    lazy var mockManager: MockManager = MockManager(self, callRealFuncClosure: { [weak self] (funcName, args) -> Any? in
        guard let __self = self else { return nil }
        return __self.callRealFunc(funcName, args)
    })
    
    fileprivate func callRealFunc(_ funcName: String, _ args: [Any?]?) -> Any? {
        switch funcName {
        case sel_changeObservingPredicate:
            return super.changeObservingPredicate(args![0] as! NSPredicate)

        case sel_repositoryCompletionHandleResponse:
            return super.repositoryCompletionHandleResponse(args![0] as! RepositoryResponse)

        case sel_shouldExecute:
            return super.shouldExecute()

        default:
            return nil
        }
    }
    
    // MARK: - UsertUseCaseImpl
    
    // MARK: - UseCaseImpl
    let sel_shouldExecute = "sel_shouldExecute"
    override func shouldExecute() -> Bool {
        return mockManager.handle(sel_shouldExecute, withDefaultReturnValue: true, withArgs: nil) as! Bool
    }

    let sel_changeObservingPredicate = "sel_changeObservingPredicate"
    override func changeObservingPredicate(_ predicate: NSPredicate) {
        mockManager.handle(sel_changeObservingPredicate, withDefaultReturnValue: nil, withArgs: predicate)
    }

    let sel_repositoryCompletionHandleResponse = "sel_repositoryCompletionHandleResponse"
    override func repositoryCompletionHandleResponse(_ response: RepositoryResponse) {
        mockManager.handle(sel_repositoryCompletionHandleResponse, withDefaultReturnValue: nil, withArgs: response)
    }
}
