import Foundation

// MARK: - API Error

enum APIError: LocalizedError {
    case invalidURL
    case noData
    case decodingFailed(Error)
    case serverError(code: Int, message: String)
    case unauthorized
    case forbidden
    case networkError(Error)
    case timeout

    var errorDescription: String? {
        switch self {
        case .invalidURL:
            return "Invalid URL"
        case .noData:
            return "No data received from server"
        case .decodingFailed(let error):
            return "Failed to parse response: \(error.localizedDescription)"
        case .serverError(_, let message):
            return message
        case .unauthorized:
            return "Session expired. Please log in again."
        case .forbidden:
            return "You don't have permission to perform this action."
        case .networkError(let error):
            return "Network error: \(error.localizedDescription)"
        case .timeout:
            return "Request timed out. Please check your connection."
        }
    }
}

// MARK: - APIClient

final class APIClient {
    static let shared = APIClient()

    private let session: URLSession
    private let decoder: JSONDecoder
    private let encoder: JSONEncoder

    private init() {
        let config = URLSessionConfiguration.default
        config.timeoutIntervalForRequest = APIConfig.requestTimeout
        config.timeoutIntervalForResource = APIConfig.requestTimeout * 2
        session = URLSession(configuration: config)

        decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .useDefaultKeys

        encoder = JSONEncoder()
        encoder.keyEncodingStrategy = .useDefaultKeys
    }

    // MARK: - Generic POST Request

    func post<RequestBody: Encodable, ResponseData: Decodable>(
        endpoint: String,
        body: RequestBody,
        requiresAuth: Bool = true
    ) async throws -> ResponseData {
        let urlString = APIConfig.apiV1 + endpoint
        guard let url = URL(string: urlString) else {
            throw APIError.invalidURL
        }

        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        request.setValue("application/json", forHTTPHeaderField: "Accept")

        if requiresAuth {
            if let token = KeychainManager.shared.retrieveToken() {
                request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
            } else {
                throw APIError.unauthorized
            }
        }

        request.httpBody = try encoder.encode(body)

        return try await performRequest(request)
    }

    // MARK: - GET Request

    func get<ResponseData: Decodable>(
        endpoint: String,
        requiresAuth: Bool = true
    ) async throws -> ResponseData {
        let urlString = APIConfig.apiV1 + endpoint
        guard let url = URL(string: urlString) else {
            throw APIError.invalidURL
        }

        var request = URLRequest(url: url)
        request.httpMethod = "GET"
        request.setValue("application/json", forHTTPHeaderField: "Accept")

        if requiresAuth {
            if let token = KeychainManager.shared.retrieveToken() {
                request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
            } else {
                throw APIError.unauthorized
            }
        }

        return try await performRequest(request)
    }

    // MARK: - Request Execution

    private func performRequest<ResponseData: Decodable>(_ request: URLRequest) async throws -> ResponseData {
        do {
            let (data, response) = try await session.data(for: request)

            guard let httpResponse = response as? HTTPURLResponse else {
                throw APIError.noData
            }

            switch httpResponse.statusCode {
            case 401:
                throw APIError.unauthorized
            case 403:
                throw APIError.forbidden
            default:
                break
            }

            let apiResponse = try decoder.decode(APIResponse<ResponseData>.self, from: data)

            if apiResponse.code == 403 {
                throw APIError.forbidden
            }

            guard apiResponse.isSuccess, let responseData = apiResponse.data else {
                throw APIError.serverError(code: apiResponse.code, message: apiResponse.message)
            }

            return responseData
        } catch let error as APIError {
            throw error
        } catch let error as URLError where error.code == .timedOut {
            throw APIError.timeout
        } catch let error as DecodingError {
            throw APIError.decodingFailed(error)
        } catch {
            throw APIError.networkError(error)
        }
    }

    // MARK: - No-Body POST (returns empty response check)

    func postEmpty<RequestBody: Encodable>(
        endpoint: String,
        body: RequestBody,
        requiresAuth: Bool = true
    ) async throws {
        let urlString = APIConfig.apiV1 + endpoint
        guard let url = URL(string: urlString) else {
            throw APIError.invalidURL
        }

        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        if requiresAuth {
            if let token = KeychainManager.shared.retrieveToken() {
                request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
            } else {
                throw APIError.unauthorized
            }
        }

        request.httpBody = try encoder.encode(body)

        do {
            let (data, _) = try await session.data(for: request)
            let apiResponse = try decoder.decode(APIResponseEmpty.self, from: data)
            if !apiResponse.isSuccess {
                throw APIError.serverError(code: apiResponse.code, message: apiResponse.message)
            }
        } catch let error as APIError {
            throw error
        } catch let error as DecodingError {
            throw APIError.decodingFailed(error)
        } catch {
            throw APIError.networkError(error)
        }
    }
}
