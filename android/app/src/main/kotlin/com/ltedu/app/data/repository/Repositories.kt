package com.ltedu.app.data.repository

import com.ltedu.app.data.api.LTEduApiService
import com.ltedu.app.data.db.dao.ChatMessageDao
import com.ltedu.app.data.db.entity.ChatMessageEntity
import com.ltedu.app.data.local.AuthPreferences
import com.ltedu.app.data.model.*
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.first
import javax.inject.Inject
import javax.inject.Singleton

sealed class Result<out T> {
    data class Success<T>(val data: T) : Result<T>()
    data class Error(val message: String) : Result<Nothing>()
}

@Singleton
class AuthRepository @Inject constructor(
    private val api: LTEduApiService,
    private val prefs: AuthPreferences
) {
    suspend fun login(username: String, password: String): Result<User> {
        return try {
            val response = api.login(LoginRequest(username, password))
            val body = response.body()
            if (response.isSuccessful && body != null && body.isSuccess && body.data != null) {
                prefs.saveToken(body.data.token)
                prefs.saveUsername(username)
                Result.Success(body.data.user)
            } else {
                Result.Error(body?.msg ?: "Login failed")
            }
        } catch (e: Exception) {
            Result.Error(e.message ?: "Network error")
        }
    }

    suspend fun logout() {
        try { api.logout() } catch (_: Exception) {}
        prefs.clear()
    }

    fun tokenFlow(): Flow<String?> = prefs.token
}

@Singleton
class ChatRepository @Inject constructor(
    private val api: LTEduApiService,
    private val dao: ChatMessageDao
) {
    fun getMessages(sessionId: String): Flow<List<ChatMessageEntity>> =
        dao.getMessages(sessionId)

    fun getSessions(): Flow<List<String>> = dao.getSessions()

    suspend fun sendMessage(sessionId: String, userText: String): Result<String> {
        // Persist user message locally first for offline UX
        dao.insert(ChatMessageEntity(sessionId = sessionId, role = "user", content = userText))

        // Collect a single snapshot; keep last 20 messages for context window management
        val history = dao.getMessages(sessionId).first()
            .takeLast(20)
            .map { ChatMessage(role = it.role, content = it.content) }

        return try {
            val response = api.chat(ChatRequest(messages = history))
            val body = response.body()
            if (response.isSuccessful && body != null && body.isSuccess && body.data != null) {
                val reply = body.data.reply
                dao.insert(ChatMessageEntity(sessionId = sessionId, role = "assistant", content = reply))
                Result.Success(reply)
            } else {
                Result.Error(body?.msg ?: "AI error")
            }
        } catch (e: Exception) {
            Result.Error(e.message ?: "Network error")
        }
    }

    suspend fun clearSession(sessionId: String) = dao.clearSession(sessionId)
}
