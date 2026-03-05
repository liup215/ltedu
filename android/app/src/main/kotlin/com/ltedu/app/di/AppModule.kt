package com.ltedu.app.di

import android.content.Context
import androidx.room.Room
import com.ltedu.app.BuildConfig
import com.ltedu.app.data.api.AuthInterceptor
import com.ltedu.app.data.api.LTEduApiService
import com.ltedu.app.data.db.LTEduDatabase
import com.ltedu.app.data.db.dao.ChatMessageDao
import com.ltedu.app.data.local.AuthPreferences
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.android.qualifiers.ApplicationContext
import dagger.hilt.components.SingletonComponent
import okhttp3.OkHttpClient
import okhttp3.logging.HttpLoggingInterceptor
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.util.concurrent.TimeUnit
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
object AppModule {

    @Provides
    @Singleton
    fun provideAuthPreferences(@ApplicationContext context: Context): AuthPreferences =
        AuthPreferences(context)

    @Provides
    @Singleton
    fun provideOkHttpClient(authPreferences: AuthPreferences): OkHttpClient {
        val logging = HttpLoggingInterceptor().apply {
            level = if (BuildConfig.DEBUG) {
                HttpLoggingInterceptor.Level.BODY
            } else {
                HttpLoggingInterceptor.Level.NONE
            }
        }
        return OkHttpClient.Builder()
            .addInterceptor(AuthInterceptor(authPreferences))
            .addInterceptor(logging)
            .connectTimeout(30, TimeUnit.SECONDS)
            .readTimeout(60, TimeUnit.SECONDS)
            .writeTimeout(30, TimeUnit.SECONDS)
            .build()
    }

    @Provides
    @Singleton
    fun provideRetrofit(client: OkHttpClient): Retrofit =
        Retrofit.Builder()
            .baseUrl(BuildConfig.API_BASE_URL)
            .client(client)
            .addConverterFactory(GsonConverterFactory.create())
            .build()

    @Provides
    @Singleton
    fun provideLTEduApiService(retrofit: Retrofit): LTEduApiService =
        retrofit.create(LTEduApiService::class.java)

    @Provides
    @Singleton
    fun provideDatabase(@ApplicationContext context: Context): LTEduDatabase =
        Room.databaseBuilder(context, LTEduDatabase::class.java, "ltedu.db")
            .fallbackToDestructiveMigration()
            .build()

    @Provides
    @Singleton
    fun provideChatMessageDao(db: LTEduDatabase): ChatMessageDao = db.chatMessageDao()
}
