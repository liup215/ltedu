import Foundation

// MARK: - AuthService

final class AuthService {
    static let shared = AuthService()
    private let apiClient = APIClient.shared

    private init() {}

    // MARK: - Captcha

    func getCaptcha() async throws -> CaptchaResponse {
        return try await apiClient.post(
            endpoint: APIEndpoint.captcha,
            body: EmptyBody(),
            requiresAuth: false
        )
    }

    // MARK: - Login

    func login(username: String, password: String, captchaId: String, captchaValue: String) async throws -> LoginResponse {
        let request = LoginRequest(
            username: username,
            password: password,
            captchaId: captchaId,
            captchaValue: captchaValue
        )
        let response: LoginResponse = try await apiClient.post(
            endpoint: APIEndpoint.login,
            body: request,
            requiresAuth: false
        )
        // Store token in Keychain
        let expiry = ISO8601DateFormatter().date(from: response.expire)
        KeychainManager.shared.saveToken(response.token, expiry: expiry)
        return response
    }

    // MARK: - Register

    func register(
        username: String,
        email: String,
        mobile: String?,
        password: String,
        passwordConfirm: String,
        verificationCode: String
    ) async throws {
        let request = RegisterRequest(
            username: username,
            email: email,
            mobile: mobile,
            password: password,
            passwordConfirm: passwordConfirm,
            verificationCode: verificationCode
        )
        try await apiClient.postEmpty(
            endpoint: APIEndpoint.register,
            body: request,
            requiresAuth: false
        )
    }

    // MARK: - Send Verification Code

    func sendVerificationCode(email: String) async throws {
        let request = SendVerificationRequest(email: email)
        try await apiClient.postEmpty(
            endpoint: APIEndpoint.sendVerificationCode,
            body: request,
            requiresAuth: false
        )
    }

    // MARK: - Current User

    func getCurrentUser() async throws -> User {
        return try await apiClient.get(endpoint: APIEndpoint.currentUser)
    }

    // MARK: - Update Account

    func updateAccount(realname: String?, nickname: String?, engname: String?, sex: Int?, mobile: String?) async throws -> User {
        let request = UpdateAccountRequest(
            realname: realname,
            nickname: nickname,
            engname: engname,
            sex: sex,
            mobile: mobile
        )
        return try await apiClient.post(endpoint: APIEndpoint.updateAccount, body: request)
    }

    // MARK: - Change Password

    func changePassword(oldPassword: String, newPassword: String) async throws {
        let request = ChangePasswordRequest(oldPassword: oldPassword, newPassword: newPassword)
        try await apiClient.postEmpty(endpoint: APIEndpoint.changePassword, body: request)
    }

    // MARK: - Logout

    func logout() {
        KeychainManager.shared.clearAll()
    }

    // MARK: - Check Auth State

    var isAuthenticated: Bool {
        return KeychainManager.shared.isTokenValid()
    }
}

// MARK: - Helper

private struct EmptyBody: Encodable {}
