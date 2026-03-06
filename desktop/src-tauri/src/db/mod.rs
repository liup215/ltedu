use std::path::Path;
use anyhow::Result;

/// Shared state that holds the path to the SQLite database file.
pub struct DbState {
    pub path: String,
}

/// Initialise the SQLite database and ensure required tables exist.
pub fn init_database(db_path: &Path) -> Result<()> {
    let conn = rusqlite::Connection::open(db_path)?;

    // Key-value store used for auth tokens and settings
    conn.execute_batch(
        "CREATE TABLE IF NOT EXISTS kv_store (
            key        TEXT PRIMARY KEY,
            value      TEXT NOT NULL,
            updated_at TEXT NOT NULL
        );",
    )?;

    // General-purpose response cache
    conn.execute_batch(
        "CREATE TABLE IF NOT EXISTS cache (
            key        TEXT PRIMARY KEY,
            value      TEXT NOT NULL,
            updated_at TEXT NOT NULL
        );",
    )?;

    // AI chat message history
    conn.execute_batch(
        "CREATE TABLE IF NOT EXISTS chat_history (
            id         INTEGER PRIMARY KEY AUTOINCREMENT,
            role       TEXT NOT NULL,
            content    TEXT NOT NULL,
            created_at TEXT NOT NULL DEFAULT (datetime('now'))
        );",
    )?;

    Ok(())
}
