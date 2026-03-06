import Foundation
import Combine

// MARK: - AuthViewModel

@MainActor
final class AuthViewModel: ObservableObject {
    @Published var isAuthenticated = false
    @Published var currentUser: User?
    @Published var isLoading = false
    @Published var errorMessage: String?

    // Captcha state
    @Published var captchaId = ""
    @Published var captchaBase64 = ""

    private let authService = AuthService.shared
    private let cacheManager = CacheManager.shared

    init() {
        isAuthenticated = authService.isAuthenticated
        if isAuthenticated {
            Task { await loadCurrentUser() }
        }
    }

    // MARK: - Login

    func login(username: String, password: String, captchaValue: String) async {
        guard !captchaId.isEmpty else {
            errorMessage = "Please load a captcha first"
            return
        }
        isLoading = true
        errorMessage = nil
        do {
            _ = try await authService.login(
                username: username,
                password: password,
                captchaId: captchaId,
                captchaValue: captchaValue
            )
            isAuthenticated = true
            await loadCurrentUser()
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
            await loadCaptcha()
        }
        isLoading = false
    }

    // MARK: - Register

    func register(
        username: String,
        email: String,
        mobile: String?,
        password: String,
        passwordConfirm: String,
        verificationCode: String
    ) async {
        isLoading = true
        errorMessage = nil
        do {
            try await authService.register(
                username: username,
                email: email,
                mobile: mobile,
                password: password,
                passwordConfirm: passwordConfirm,
                verificationCode: verificationCode
            )
            // Auto-login after registration - user must log in manually
            errorMessage = nil
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }

    // MARK: - Logout

    func logout() {
        authService.logout()
        isAuthenticated = false
        currentUser = nil
    }

    // MARK: - Load Current User

    func loadCurrentUser() async {
        do {
            let user = try await authService.getCurrentUser()
            currentUser = user
            cacheManager.cacheUser(user)
        } catch let error as APIError {
            if case .unauthorized = error {
                logout()
            }
        } catch {
            // Silently fail; use cached user if available
        }
    }

    // MARK: - Update Profile

    func updateProfile(realname: String?, nickname: String?, sex: Int?, mobile: String?) async {
        isLoading = true
        errorMessage = nil
        do {
            let updated = try await authService.updateAccount(
                realname: realname,
                nickname: nickname,
                engname: nil,
                sex: sex,
                mobile: mobile
            )
            currentUser = updated
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }

    // MARK: - Captcha

    func loadCaptcha() async {
        do {
            let captcha = try await authService.getCaptcha()
            captchaId = captcha.captchaId
            captchaBase64 = captcha.base64
        } catch {
            errorMessage = "Failed to load captcha: \(error.localizedDescription)"
        }
    }

    // MARK: - Send Verification Code

    func sendVerificationCode(email: String) async {
        do {
            try await authService.sendVerificationCode(email: email)
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
    }
}
