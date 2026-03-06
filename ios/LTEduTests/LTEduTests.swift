import XCTest
@testable import LTEdu

// MARK: - APIConfig Tests

final class APIConfigTests: XCTestCase {

    func testDefaultBaseURL() {
        // When no environment override is set, the default base URL should be used
        let url = APIConfig.baseURL
        XCTAssertFalse(url.isEmpty)
        XCTAssertTrue(url.hasPrefix("http"))
    }

    func testAPIV1URL() {
        let v1 = APIConfig.apiV1
        XCTAssertTrue(v1.hasSuffix("/v1"))
    }

    func testDefaultPageSize() {
        XCTAssertGreaterThan(APIConfig.defaultPageSize, 0)
    }

    func testRequestTimeout() {
        XCTAssertGreaterThan(APIConfig.requestTimeout, 0)
    }
}

// MARK: - NetworkModels Tests

final class NetworkModelsTests: XCTestCase {

    // MARK: - User decoding

    func testUserDecoding() throws {
        let json = """
        {
            "id": 1,
            "username": "testuser",
            "email": "test@example.com",
            "nickname": "Test User",
            "isAdmin": false,
            "isTeacher": false
        }
        """
        let data = try XCTUnwrap(json.data(using: .utf8))
        let user = try JSONDecoder().decode(User.self, from: data)
        XCTAssertEqual(user.id, 1)
        XCTAssertEqual(user.username, "testuser")
        XCTAssertEqual(user.email, "test@example.com")
        XCTAssertFalse(user.isAdmin ?? true)
    }

    func testUserDisplayName() {
        let userWithNickname = User(
            id: 1, username: "user1", email: nil, nickname: "Nick",
            realname: "Real", engname: nil, mobile: nil, avatar: nil,
            sex: nil, status: nil, isAdmin: nil, isTeacher: nil,
            teacherApplyStatus: nil, vipExpireAt: nil, createdAt: nil, updatedAt: nil
        )
        XCTAssertEqual(userWithNickname.displayName, "Nick")

        let userWithRealname = User(
            id: 2, username: "user2", email: nil, nickname: nil,
            realname: "Real Name", engname: nil, mobile: nil, avatar: nil,
            sex: nil, status: nil, isAdmin: nil, isTeacher: nil,
            teacherApplyStatus: nil, vipExpireAt: nil, createdAt: nil, updatedAt: nil
        )
        XCTAssertEqual(userWithRealname.displayName, "Real Name")

        let userWithOnlyUsername = User(
            id: 3, username: "user3", email: nil, nickname: nil,
            realname: nil, engname: nil, mobile: nil, avatar: nil,
            sex: nil, status: nil, isAdmin: nil, isTeacher: nil,
            teacherApplyStatus: nil, vipExpireAt: nil, createdAt: nil, updatedAt: nil
        )
        XCTAssertEqual(userWithOnlyUsername.displayName, "user3")
    }

    func testUserVIPCheck() {
        // Expired VIP
        let expiredUser = User(
            id: 1, username: "u", email: nil, nickname: nil, realname: nil,
            engname: nil, mobile: nil, avatar: nil, sex: nil, status: nil,
            isAdmin: nil, isTeacher: nil, teacherApplyStatus: nil,
            vipExpireAt: "2020-01-01T00:00:00Z",
            createdAt: nil, updatedAt: nil
        )
        XCTAssertFalse(expiredUser.isVip)

        // Active VIP (expires in the future)
        let futureDate = Date().addingTimeInterval(86400)
        let formatter = ISO8601DateFormatter()
        let activeUser = User(
            id: 2, username: "u", email: nil, nickname: nil, realname: nil,
            engname: nil, mobile: nil, avatar: nil, sex: nil, status: nil,
            isAdmin: nil, isTeacher: nil, teacherApplyStatus: nil,
            vipExpireAt: formatter.string(from: futureDate),
            createdAt: nil, updatedAt: nil
        )
        XCTAssertTrue(activeUser.isVip)

        // No VIP date
        let noVipUser = User(
            id: 3, username: "u", email: nil, nickname: nil, realname: nil,
            engname: nil, mobile: nil, avatar: nil, sex: nil, status: nil,
            isAdmin: nil, isTeacher: nil, teacherApplyStatus: nil,
            vipExpireAt: nil, createdAt: nil, updatedAt: nil
        )
        XCTAssertFalse(noVipUser.isVip)
    }

    // MARK: - APIResponse decoding

    func testAPIResponseSuccess() throws {
        let json = """
        {"code": 0, "message": "success", "data": {"id": 1, "username": "testuser"}}
        """
        let data = try XCTUnwrap(json.data(using: .utf8))
        // Decode as raw dictionary first
        let decoded = try JSONDecoder().decode(APIResponse<[String: AnyCodable]>.self, from: data)
        XCTAssertEqual(decoded.code, 0)
        XCTAssertTrue(decoded.isSuccess)
    }

    func testAPIResponseError() throws {
        let json = """
        {"code": 1, "message": "error occurred", "data": null}
        """
        let data = try XCTUnwrap(json.data(using: .utf8))
        let decoded = try JSONDecoder().decode(APIResponseEmpty.self, from: data)
        XCTAssertEqual(decoded.code, 1)
        XCTAssertFalse(decoded.isSuccess)
        XCTAssertEqual(decoded.message, "error occurred")
    }

    // MARK: - Course decoding

    func testCourseDecoding() throws {
        let json = """
        {
            "id": 10,
            "title": "Math Course",
            "slug": "math-course",
            "isFree": 1,
            "charge": 0,
            "userCount": 50
        }
        """
        let data = try XCTUnwrap(json.data(using: .utf8))
        let course = try JSONDecoder().decode(Course.self, from: data)
        XCTAssertEqual(course.id, 10)
        XCTAssertEqual(course.title, "Math Course")
        XCTAssertTrue(course.isFreeAccess)
    }

    // MARK: - ChatMessage

    func testChatMessageInit() {
        let message = ChatMessage(role: .user, content: "Hello")
        XCTAssertEqual(message.role, .user)
        XCTAssertEqual(message.content, "Hello")
        XCTAssertNotNil(message.id)
    }

    func testChatSessionInit() {
        let session = ChatSession()
        XCTAssertEqual(session.title, "New Chat")
        XCTAssertTrue(session.messages.isEmpty)
    }

    func testChatSessionCodable() throws {
        var session = ChatSession(title: "Test Session")
        session.messages.append(ChatMessage(role: .user, content: "Question"))
        session.messages.append(ChatMessage(role: .assistant, content: "Answer"))

        let encoded = try JSONEncoder().encode(session)
        let decoded = try JSONDecoder().decode(ChatSession.self, from: encoded)

        XCTAssertEqual(decoded.id, session.id)
        XCTAssertEqual(decoded.title, "Test Session")
        XCTAssertEqual(decoded.messages.count, 2)
        XCTAssertEqual(decoded.messages[0].role, .user)
        XCTAssertEqual(decoded.messages[1].role, .assistant)
    }

    // MARK: - PaginationRequest

    func testPaginationRequestDefaults() {
        let req = PaginationRequest()
        XCTAssertEqual(req.pageIndex, 1)
        XCTAssertEqual(req.pageSize, APIConfig.defaultPageSize)
    }

    func testPaginationRequestCustom() {
        let req = PaginationRequest(pageIndex: 3, pageSize: 50)
        XCTAssertEqual(req.pageIndex, 3)
        XCTAssertEqual(req.pageSize, 50)
    }
}

// MARK: - KeychainManager Tests

final class KeychainManagerTests: XCTestCase {

    override func setUp() {
        super.setUp()
        KeychainManager.shared.clearAll()
    }

    override func tearDown() {
        KeychainManager.shared.clearAll()
        super.tearDown()
    }

    func testSaveAndRetrieveToken() {
        let token = "test_token_12345"
        let expiry = Date().addingTimeInterval(3600)

        KeychainManager.shared.saveToken(token, expiry: expiry)
        let retrieved = KeychainManager.shared.retrieveToken()

        XCTAssertEqual(retrieved, token)
    }

    func testTokenExpiry() {
        let token = "expired_token"
        let pastExpiry = Date().addingTimeInterval(-3600) // 1 hour ago

        KeychainManager.shared.saveToken(token, expiry: pastExpiry)
        let retrieved = KeychainManager.shared.retrieveToken()

        XCTAssertNil(retrieved, "Expired token should not be returned")
    }

    func testClearToken() {
        let token = "token_to_clear"
        KeychainManager.shared.saveToken(token, expiry: Date().addingTimeInterval(3600))
        XCTAssertNotNil(KeychainManager.shared.retrieveToken())

        KeychainManager.shared.clearToken()
        XCTAssertNil(KeychainManager.shared.retrieveToken())
    }

    func testIsTokenValid() {
        XCTAssertFalse(KeychainManager.shared.isTokenValid())

        let validToken = "valid_token"
        KeychainManager.shared.saveToken(validToken, expiry: Date().addingTimeInterval(3600))
        XCTAssertTrue(KeychainManager.shared.isTokenValid())
    }

    func testSaveAndRetrieveUserId() {
        KeychainManager.shared.saveUserId(42)
        XCTAssertEqual(KeychainManager.shared.retrieveUserId(), 42)
    }

    func testClearAll() {
        KeychainManager.shared.saveToken("tok", expiry: Date().addingTimeInterval(3600))
        KeychainManager.shared.saveUserId(7)

        KeychainManager.shared.clearAll()

        XCTAssertNil(KeychainManager.shared.retrieveToken())
        XCTAssertNil(KeychainManager.shared.retrieveUserId())
    }
}

// MARK: - Extensions Tests

final class ExtensionsTests: XCTestCase {

    func testIsValidEmail() {
        XCTAssertTrue("user@example.com".isValidEmail)
        XCTAssertTrue("user+tag@sub.domain.org".isValidEmail)
        XCTAssertFalse("not-an-email".isValidEmail)
        XCTAssertFalse("missing@tld".isValidEmail)
        XCTAssertFalse("@nodomain.com".isValidEmail)
    }

    func testIsValidPassword() {
        XCTAssertTrue("password".isValidPassword)
        XCTAssertTrue("123456".isValidPassword)
        XCTAssertFalse("12345".isValidPassword)
        XCTAssertFalse("".isValidPassword)
    }

    func testStrippingHTML() {
        let html = "<p>Hello <b>World</b></p>"
        XCTAssertEqual(html.strippingHTML, "Hello World")
    }
}

// MARK: - APIError Tests

final class APIErrorTests: XCTestCase {

    func testErrorDescriptions() {
        XCTAssertNotNil(APIError.invalidURL.errorDescription)
        XCTAssertNotNil(APIError.noData.errorDescription)
        XCTAssertNotNil(APIError.unauthorized.errorDescription)
        XCTAssertNotNil(APIError.forbidden.errorDescription)
        XCTAssertNotNil(APIError.timeout.errorDescription)
    }

    func testServerErrorIncludesMessage() {
        let error = APIError.serverError(code: 1, message: "Custom error message")
        XCTAssertEqual(error.errorDescription, "Custom error message")
    }
}

// MARK: - AnyCodable Tests

final class AnyCodableTests: XCTestCase {

    func testEncodeInt() throws {
        let value = AnyCodable(42)
        let data = try JSONEncoder().encode(value)
        let decoded = try JSONDecoder().decode(AnyCodable.self, from: data)
        XCTAssertEqual(decoded.value as? Int, 42)
    }

    func testEncodeString() throws {
        let value = AnyCodable("hello")
        let data = try JSONEncoder().encode(value)
        let decoded = try JSONDecoder().decode(AnyCodable.self, from: data)
        XCTAssertEqual(decoded.value as? String, "hello")
    }

    func testEncodeBool() throws {
        let value = AnyCodable(true)
        let data = try JSONEncoder().encode(value)
        let decoded = try JSONDecoder().decode(AnyCodable.self, from: data)
        XCTAssertEqual(decoded.value as? Bool, true)
    }

    func testEncodeDouble() throws {
        let value = AnyCodable(3.14)
        let data = try JSONEncoder().encode(value)
        let decoded = try JSONDecoder().decode(AnyCodable.self, from: data)
        XCTAssertEqual(decoded.value as? Double, 3.14, accuracy: 0.001)
    }
}
