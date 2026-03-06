package com.ltedu.app.data.api

import com.ltedu.app.data.local.AuthPreferences
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.runBlocking
import okhttp3.Interceptor
import okhttp3.Response

/**
 * OkHttp interceptor that attaches the stored JWT bearer token to every request.
 */
class AuthInterceptor(private val authPreferences: AuthPreferences) : Interceptor {

    override fun intercept(chain: Interceptor.Chain): Response {
        val token = runBlocking { authPreferences.token.first() }
        val request = if (token != null) {
            chain.request().newBuilder()
                .header("Authorization", "Bearer $token")
                .build()
        } else {
            chain.request()
        }
        return chain.proceed(request)
    }
}
