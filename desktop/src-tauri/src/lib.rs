mod commands;
mod db;

use tauri::Manager;

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_shell::init())
        .plugin(tauri_plugin_notification::init())
        .plugin(tauri_plugin_fs::init())
        .plugin(tauri_plugin_http::init())
        .plugin(tauri_plugin_store::Builder::default().build())
        .setup(|app| {
            let app_data_dir = app
                .path()
                .app_data_dir()
                .expect("Failed to get app data directory");
            std::fs::create_dir_all(&app_data_dir)
                .expect("Failed to create app data directory");

            let db_path = app_data_dir.join("ltedu_cache.db");
            db::init_database(&db_path).expect("Failed to initialize database");

            // Store the DB path for use in commands
            app.manage(db::DbState {
                path: db_path.to_string_lossy().to_string(),
            });

            Ok(())
        })
        .invoke_handler(tauri::generate_handler![
            commands::auth::cmd_login,
            commands::auth::cmd_logout,
            commands::auth::cmd_get_current_user,
            commands::api::cmd_api_request,
            commands::cache::cmd_cache_get,
            commands::cache::cmd_cache_set,
            commands::cache::cmd_cache_delete,
            commands::cache::cmd_cache_clear,
            commands::cache::cmd_get_cached_resources,
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
