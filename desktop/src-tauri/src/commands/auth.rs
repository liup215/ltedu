use serde::{Deserialize, Serialize};
use tauri::State;

use crate::db::DbState;

#[derive(Debug, Serialize, Deserialize)]
pub struct LoginRequest {
    pub username: String,
    pub password: String,
    pub api_base_url: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct AuthResult {
    pub token: String,
    pub user: serde_json::Value,
}

/// Perform login against the LTEdu cloud API and persist the token in the
/// local SQLite store so the session survives restarts.
#[tauri::command]
pub async fn cmd_login(
    request: LoginRequest,
    db: State<'_, DbState>,
) -> Result<AuthResult, String> {
    let client = reqwest::Client::new();
    let url = format!("{}/api/v1/login", request.api_base_url.trim_end_matches('/'));

    let payload = serde_json::json!({
        "username": request.username,
        "password": request.password,
    });

    let response = client
        .post(&url)
        .json(&payload)
        .send()
        .await
        .map_err(|e| format!("Network error: {e}"))?;

    let body: serde_json::Value = response
        .json()
        .await
        .map_err(|e| format!("Failed to parse response: {e}"))?;

    if body["code"].as_i64().unwrap_or(1) != 0 {
        let msg = body["message"].as_str().unwrap_or("Login failed").to_string();
        return Err(msg);
    }

    let token = body["data"]["token"]
        .as_str()
        .ok_or("Token missing in response")?
        .to_string();

    // Persist token + api_base_url
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;
    conn.execute(
        "INSERT OR REPLACE INTO kv_store (key, value, updated_at) VALUES (?1, ?2, datetime('now'))",
        rusqlite::params!["auth_token", &token],
    )
    .map_err(|e| e.to_string())?;
    conn.execute(
        "INSERT OR REPLACE INTO kv_store (key, value, updated_at) VALUES (?1, ?2, datetime('now'))",
        rusqlite::params!["api_base_url", &request.api_base_url],
    )
    .map_err(|e| e.to_string())?;

    // Fetch user profile
    let profile_url = format!("{}/api/v1/user", request.api_base_url.trim_end_matches('/'));
    let profile_resp = client
        .get(&profile_url)
        .bearer_auth(&token)
        .send()
        .await
        .map_err(|e| format!("Network error fetching profile: {e}"))?;

    let profile_body: serde_json::Value = profile_resp
        .json()
        .await
        .map_err(|e| format!("Failed to parse profile: {e}"))?;

    let user = if profile_body["code"].as_i64().unwrap_or(1) == 0 {
        profile_body["data"].clone()
    } else {
        serde_json::Value::Null
    };

    // Persist user profile
    conn.execute(
        "INSERT OR REPLACE INTO kv_store (key, value, updated_at) VALUES (?1, ?2, datetime('now'))",
        rusqlite::params!["user_profile", serde_json::to_string(&user).unwrap_or_default()],
    )
    .map_err(|e| e.to_string())?;

    Ok(AuthResult { token, user })
}

/// Clear the persisted session.
#[tauri::command]
pub fn cmd_logout(db: State<'_, DbState>) -> Result<(), String> {
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;
    conn.execute("DELETE FROM kv_store WHERE key IN ('auth_token', 'user_profile')", [])
        .map_err(|e| e.to_string())?;
    Ok(())
}

/// Return persisted token + user profile (if any) so the frontend can restore
/// the session without a round-trip to the cloud API.
#[tauri::command]
pub fn cmd_get_current_user(db: State<'_, DbState>) -> Result<serde_json::Value, String> {
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;

    let token: Option<String> = conn
        .query_row(
            "SELECT value FROM kv_store WHERE key = 'auth_token'",
            [],
            |row| row.get(0),
        )
        .ok();

    let user_json: Option<String> = conn
        .query_row(
            "SELECT value FROM kv_store WHERE key = 'user_profile'",
            [],
            |row| row.get(0),
        )
        .ok();

    let api_base_url: Option<String> = conn
        .query_row(
            "SELECT value FROM kv_store WHERE key = 'api_base_url'",
            [],
            |row| row.get(0),
        )
        .ok();

    let user: serde_json::Value = user_json
        .and_then(|s| serde_json::from_str(&s).ok())
        .unwrap_or(serde_json::Value::Null);

    Ok(serde_json::json!({
        "token": token,
        "user": user,
        "apiBaseUrl": api_base_url,
    }))
}
