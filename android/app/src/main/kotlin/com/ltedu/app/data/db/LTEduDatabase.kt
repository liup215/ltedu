package com.ltedu.app.data.db

import androidx.room.Database
import androidx.room.RoomDatabase
import com.ltedu.app.data.db.dao.ChatMessageDao
import com.ltedu.app.data.db.entity.ChatMessageEntity

@Database(
    entities = [ChatMessageEntity::class],
    version = 1,
    exportSchema = false
)
abstract class LTEduDatabase : RoomDatabase() {
    abstract fun chatMessageDao(): ChatMessageDao
}
