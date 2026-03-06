package com.ltedu.app.data.api

import com.ltedu.app.data.model.*
import retrofit2.Response
import retrofit2.http.*

interface LTEduApiService {

    @POST("/api/v1/user/login")
    suspend fun login(@Body request: LoginRequest): Response<ApiResponse<LoginResponse>>

    @POST("/api/v1/user/logout")
    suspend fun logout(): Response<ApiResponse<Unit>>

    @GET("/api/v1/user/me")
    suspend fun getMe(): Response<ApiResponse<User>>

    @POST("/api/v1/ai/chat")
    suspend fun chat(@Body request: ChatRequest): Response<ApiResponse<ChatResponse>>

    @GET("/api/v1/learning-plan/list")
    suspend fun getLearningPlans(
        @Query("classId") classId: Long
    ): Response<ApiResponse<ListResponse<LearningPlan>>>

    @GET("/api/v1/course/list")
    suspend fun getCourses(): Response<ApiResponse<ListResponse<Course>>>

    @PUT("/api/v1/user/fcm-token")
    suspend fun updateFcmToken(@Body body: Map<String, String>): Response<ApiResponse<Unit>>
}
