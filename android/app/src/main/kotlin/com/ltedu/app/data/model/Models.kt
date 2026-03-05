package com.ltedu.app.data.model

import com.google.gson.annotations.SerializedName

/** Generic API response wrapper matching the LTEdu backend (code=0 success, code=1 error). */
data class ApiResponse<T>(
    @SerializedName("code") val code: Int,
    @SerializedName("msg") val msg: String?,
    @SerializedName("data") val data: T?
) {
    val isSuccess: Boolean get() = code == 0
}

data class LoginRequest(
    @SerializedName("username") val username: String,
    @SerializedName("password") val password: String
)

data class LoginResponse(
    @SerializedName("token") val token: String,
    @SerializedName("user") val user: User
)

data class User(
    @SerializedName("id") val id: Long,
    @SerializedName("username") val username: String,
    @SerializedName("name") val name: String?,
    @SerializedName("email") val email: String?,
    @SerializedName("isAdmin") val isAdmin: Boolean = false
)

data class ChatMessage(
    @SerializedName("role") val role: String,   // "user" | "assistant"
    @SerializedName("content") val content: String
)

data class ChatRequest(
    @SerializedName("messages") val messages: List<ChatMessage>,
    @SerializedName("stream") val stream: Boolean = false
)

data class ChatResponse(
    @SerializedName("reply") val reply: String
)

data class ListResponse<T>(
    @SerializedName("list") val list: List<T>,
    @SerializedName("total") val total: Int
)

data class Course(
    @SerializedName("id") val id: Long,
    @SerializedName("name") val name: String,
    @SerializedName("description") val description: String?
)

data class LearningPlan(
    @SerializedName("id") val id: Long,
    @SerializedName("name") val name: String,
    @SerializedName("startDate") val startDate: String?,
    @SerializedName("endDate") val endDate: String?,
    @SerializedName("status") val status: String?
)
