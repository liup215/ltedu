package com.ltedu.app.data.db.dao

import androidx.room.*
import com.ltedu.app.data.db.entity.ChatMessageEntity
import kotlinx.coroutines.flow.Flow

@Dao
interface ChatMessageDao {

    @Query("SELECT * FROM chat_message WHERE session_id = :sessionId ORDER BY timestamp ASC")
    fun getMessages(sessionId: String): Flow<List<ChatMessageEntity>>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insert(message: ChatMessageEntity)

    @Query("DELETE FROM chat_message WHERE session_id = :sessionId")
    suspend fun clearSession(sessionId: String)

    @Query("SELECT DISTINCT session_id FROM chat_message ORDER BY timestamp DESC")
    fun getSessions(): Flow<List<String>>
}
