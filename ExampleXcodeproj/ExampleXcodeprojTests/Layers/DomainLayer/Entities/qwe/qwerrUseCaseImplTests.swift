import XCTest
import Mirage

import UseCase
@testable import ExampleXcodeproj

class qweUseCaseImplTests: XCTestCase {

    // <no value>
    // <no value>
    // <no value>
    var sut: PartialMockqwerrUseCaseImpl!
    
    var mockRepository: MockqweRepository!
    
    override func setUp() {
        super.setUp()
        // Put setup code here. This method is called before the invocation of each test method in the class.
        
        mockRepository = MockqweRepository()
        
        sut = PartialMockqwerrUseCaseImpl(mockRepository: MockqweRepository)
    }
    
    override func tearDown() {
        mockRepository = nil
        
        sut = nil
        
        // Put teardown code here. This method is called after the invocation of each test method in the class.
        super.tearDown()
    }

    // MARK: - observe
    func test_GivenAllParams_WhenObserve_ThenItShouldCallRepositoryPassingAllParams() {
        //  given
        let param = anyParam 
        sut.<#param#> = param
        
        //  when
        sut.observe()
        
        //  then
        XCTAssertNoThrow(try mockRepository.verify(mockRepository.<#function name#>, Once()))
        
        guard let args = mockRepository.argsOf(mockRepository.<#function name#>) else {
            XCTFail(<#nilArgumentExceptionMessage#>)
            return
        }

        guard let argParam = args.first as? NSNumber else {
            XCTFail(<#transformationExceptionMessage#>)
            return
        }

        XCTAssert(argParam == param)
    }
    
    func test_GivenNilIdentifier_WhenObserve_ThenItUsesDefaultIdentifier() {
        //  given
        
        //  when
        sut.observe()
        
        //  then
        XCTAssertNoThrow(try mockRepository.verify(mockRepository.<#functionName#>, Once()))
        
        guard let args = mockRepository.argsOf(mockRepository.<#function name#>) else {
            XCTFail(<#nilArgumentExceptionMessage#>)
            return
        }

        guard let argParam = args.first as? NSNumber else {
            XCTFail(<#transformationExceptionMessage#>)
            return
        }

        XCTAssert(argParam == sut.kDefaultParam)
    }

    // MARK: - execute
    func test_GivenShouldExecuteReturnsTrue_WhenExecute_ThenItShouldCallRepositoryPassingAllParams() {
        //  given
        let param = anyParam 
        sut.<#param#> = param

        sut.when(sut.sel_shouldExecute).thenReturn(true)
        
        //  when
        sut.execute()
        
        //  then
        XCTAssertNoThrow(try mockRepository.verify(mockRepository.<#functionName#>, Once()))
        
        guard let args = mockRepository.argsOf(mockRepository.<#function name#>) else {
            XCTFail(<#nilArgumentExceptionMessage#>)
            return
        }

        guard let argParam = args.first as? NSNumber else {
            XCTFail(<#transformationExceptionMessage#>)
            return
        }

        XCTAssert(argParam == sut.kDefaultParam)
    }
    
    func test_GivenShouldExecuteReturnsFalse_WhenExecute_ThenItShouldNotCallRepository() {
        //  given
        sut.when(sut.sel_shouldExecute).thenReturn(false)
        
        //  when
        sut.execute()
        
        //  then
        XCTAssertNoThrow(try mockRepository.verify(mockRepository.<#functionName#>, Never()))
    }

    // MARK: - params validation
    func test_GivenValidParams_WhenValidateParams_ThenItShouldNotThrow() {
        //  given
        sut.<#param#> = anyParam
        
        //  when
        XCTAssertNoThrow(try sut.validateParams())
        
        //  then
    }
    
    func test_GivenNilParam_WhenValidateParams_ThenItShouldThrowNilIdentifierError() {
        //  given
        sut.<#param#> = nil
        
        //  when
        XCTAssertThrowsError(try sut.validateParams(), "") { (error) in
            XCTAssert(error is qweUseCaseNilParamError)
        }
        
        //  then
    }

}