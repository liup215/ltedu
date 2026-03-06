use serde::{Deserialize, Serialize};
use tauri::State;

use crate::db::DbState;

#[derive(Debug, Serialize, Deserialize)]
pub struct ApiRequest {
    pub method: String,
    pub path: String,
    pub body: Option<serde_json::Value>,
    pub query: Option<std::collections::HashMap<String, String>>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct ApiResponse {
    pub status: u16,
    pub body: serde_json::Value,
}

/// Forward an authenticated API request to the LTEdu cloud API.
/// The token and api_base_url are read from the local SQLite store.
#[tauri::command]
pub async fn cmd_api_request(
    request: ApiRequest,
    db: State<'_, DbState>,
) -> Result<ApiResponse, String> {
    let conn = rusqlite::Connection::open(&db.path).map_err(|e| e.to_string())?;

    let token: Option<String> = conn
        .query_row(
            "SELECT value FROM kv_store WHERE key = 'auth_token'",
            [],
            |row| row.get(0),
        )
        .ok();

    let api_base_url: String = conn
        .query_row(
            "SELECT value FROM kv_store WHERE key = 'api_base_url'",
            [],
            |row| row.get(0),
        )
        .map_err(|_| "API base URL not configured. Please log in first.".to_string())?;

    let url = format!(
        "{}/api{}",
        api_base_url.trim_end_matches('/'),
        request.path
    );

    let client = reqwest::Client::new();
    let method = request.method.to_uppercase();

    let mut builder = match method.as_str() {
        "GET" => client.get(&url),
        "POST" => client.post(&url),
        "PUT" => client.put(&url),
        "PATCH" => client.patch(&url),
        "DELETE" => client.delete(&url),
        other => return Err(format!("Unsupported HTTP method: {other}")),
    };

    if let Some(tok) = token {
        builder = builder.bearer_auth(tok);
    }

    if let Some(query) = request.query {
        builder = builder.query(&query.into_iter().collect::<Vec<_>>());
    }

    if let Some(body) = request.body {
        builder = builder.json(&body);
    }

    let response = builder
        .send()
        .await
        .map_err(|e| format!("Network error: {e}"))?;

    let status = response.status().as_u16();
    let body: serde_json::Value = response
        .json()
        .await
        .unwrap_or(serde_json::Value::Null);

    Ok(ApiResponse { status, body })
}
