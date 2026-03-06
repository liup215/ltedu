use serde::{Deserialize, Serialize};
use tauri::State;

use crate::db::DbState;

#[derive(Debug, Serialize, Deserialize)]
pub struct CachedResource {
    pub key: String,
    pub value: String,
    pub updated_at: String,
}

/// Read a value from the local cache by key.
#[tauri::command]
pub fn cmd_cache_get(key: String, db: State<'_, DbState>) -> Result<Option<String>, String> {
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;

    let result: Option<String> = conn
        .query_row(
            "SELECT value FROM cache WHERE key = ?1",
            rusqlite::params![key],
            |row| row.get(0),
        )
        .ok();

    Ok(result)
}

/// Write a value into the local cache.
#[tauri::command]
pub fn cmd_cache_set(key: String, value: String, db: State<'_, DbState>) -> Result<(), String> {
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;

    conn.execute(
        "INSERT OR REPLACE INTO cache (key, value, updated_at) VALUES (?1, ?2, datetime('now'))",
        rusqlite::params![key, value],
    )
    .map_err(|e| e.to_string())?;

    Ok(())
}

/// Remove a single cache entry.
#[tauri::command]
pub fn cmd_cache_delete(key: String, db: State<'_, DbState>) -> Result<(), String> {
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;

    conn.execute(
        "DELETE FROM cache WHERE key = ?1",
        rusqlite::params![key],
    )
    .map_err(|e| e.to_string())?;

    Ok(())
}

/// Wipe all cached entries.
#[tauri::command]
pub fn cmd_cache_clear(db: State<'_, DbState>) -> Result<(), String> {
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;
    conn.execute("DELETE FROM cache", []).map_err(|e| e.to_string())?;
    Ok(())
}

/// List all cached resources (key + metadata, without values).
#[tauri::command]
pub fn cmd_get_cached_resources(db: State<'_, DbState>) -> Result<Vec<CachedResource>, String> {
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;

    let mut stmt = conn
        .prepare("SELECT key, value, updated_at FROM cache ORDER BY updated_at DESC")
        .map_err(|e| e.to_string())?;

    let rows = stmt
        .query_map([], |row| {
            Ok(CachedResource {
                key: row.get(0)?,
                value: row.get(1)?,
                updated_at: row.get(2)?,
            })
        })
        .map_err(|e| e.to_string())?;

    let mut resources = Vec::new();
    for row in rows {
        resources.push(row.map_err(|e| e.to_string())?);
    }

    Ok(resources)
}
