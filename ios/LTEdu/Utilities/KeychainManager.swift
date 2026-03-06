import Foundation
import Security

// MARK: - Keychain Keys

private enum KeychainKey {
    static let token = "com.ltedu.auth.token"
    static let tokenExpiry = "com.ltedu.auth.token.expiry"
    static let userId = "com.ltedu.auth.userId"
}

// MARK: - KeychainManager

final class KeychainManager {
    static let shared = KeychainManager()

    private init() {}

    // MARK: - Token Management

    func saveToken(_ token: String, expiry: Date? = nil) {
        save(key: KeychainKey.token, value: token)
        if let expiry = expiry {
            let expiryString = ISO8601DateFormatter().string(from: expiry)
            save(key: KeychainKey.tokenExpiry, value: expiryString)
        }
    }

    func retrieveToken() -> String? {
        guard let token = retrieve(key: KeychainKey.token) else { return nil }

        // Check if token is expired
        if let expiryString = retrieve(key: KeychainKey.tokenExpiry),
           let expiryDate = ISO8601DateFormatter().date(from: expiryString),
           expiryDate <= Date() {
            clearToken()
            return nil
        }

        return token
    }

    func isTokenValid() -> Bool {
        return retrieveToken() != nil
    }

    func clearToken() {
        delete(key: KeychainKey.token)
        delete(key: KeychainKey.tokenExpiry)
    }

    // MARK: - User ID Management

    func saveUserId(_ userId: Int) {
        save(key: KeychainKey.userId, value: "\(userId)")
    }

    func retrieveUserId() -> Int? {
        guard let value = retrieve(key: KeychainKey.userId) else { return nil }
        return Int(value)
    }

    // MARK: - Clear All

    func clearAll() {
        clearToken()
        delete(key: KeychainKey.userId)
    }

    // MARK: - Private Keychain Operations

    private func save(key: String, value: String) {
        guard let data = value.data(using: .utf8) else { return }

        let query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrAccount as String: key,
            kSecAttrAccessible as String: kSecAttrAccessibleAfterFirstUnlock
        ]

        let attributes: [String: Any] = [
            kSecValueData as String: data
        ]

        let status = SecItemUpdate(query as CFDictionary, attributes as CFDictionary)
        if status == errSecItemNotFound {
            var newItem = query
            newItem[kSecValueData as String] = data
            SecItemAdd(newItem as CFDictionary, nil)
        }
    }

    private func retrieve(key: String) -> String? {
        let query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrAccount as String: key,
            kSecReturnData as String: kCFBooleanTrue as Any,
            kSecMatchLimit as String: kSecMatchLimitOne
        ]

        var result: AnyObject?
        let status = SecItemCopyMatching(query as CFDictionary, &result)

        guard status == errSecSuccess,
              let data = result as? Data,
              let value = String(data: data, encoding: .utf8) else {
            return nil
        }

        return value
    }

    private func delete(key: String) {
        let query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrAccount as String: key
        ]
        SecItemDelete(query as CFDictionary)
    }
}
