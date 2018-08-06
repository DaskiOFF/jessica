import Foundation
import XCTest

@testable import {{ .projectName }}

class PartialMock{{ .moduleInfo.name }}{{.answers.suffix}}UseCaseImpl: {{ .moduleInfo.name }}{{.answers.suffix}}UseCaseImpl, PartialMock {
    // MARK: - {{ .moduleInfo.name }}{{.answers.suffix}}UseCaseImpl
}
