# LTEdu Desktop Client

A cross-platform desktop application built with [Tauri](https://tauri.app/) (Rust + WebView) providing an AI-native interface to the LTEdu platform.

## Features

- **AI Chat Interface** – conversational primary interface powered by the LTEdu API
- **Cross-platform** – Windows, macOS, and Linux via Tauri
- **Offline cache** – SQLite-backed local caching for offline capability
- **System integration** – native notifications, file system access
- **Small binary** – Tauri's WebView approach targets a binary under 50 MB
- **Fast startup** – native code for backend; pre-built WebView for frontend

## Architecture

```
desktop/
├── src/                    # Vue 3 + TypeScript frontend
│   ├── components/
│   │   ├── AIChatInterface.vue   # Primary AI chat UI
│   │   └── AppSidebar.vue        # Navigation sidebar
│   ├── views/
│   │   ├── LoginView.vue         # Authentication screen
│   │   ├── MainLayout.vue        # App shell with sidebar
│   │   ├── ChatView.vue          # Chat page
│   │   └── SettingsView.vue      # Server URL & cache management
│   ├── stores/
│   │   └── authStore.ts          # Pinia auth state (token, user)
│   ├── services/
│   │   └── tauriClient.ts        # Tauri invoke wrappers
│   └── router/index.ts           # Vue Router
└── src-tauri/                    # Rust Tauri backend
    ├── src/
    │   ├── main.rs               # Entry point (no console on Windows release)
    │   ├── lib.rs                # Tauri builder + plugin registration
    │   ├── commands/
    │   │   ├── auth.rs           # login / logout / get_current_user
    │   │   ├── api.rs            # Proxy authenticated API requests
    │   │   └── cache.rs          # SQLite cache CRUD
    │   └── db/mod.rs             # SQLite initialisation (kv_store, cache, chat_history)
    ├── tauri.conf.json           # App metadata, window config, bundle targets
    └── capabilities/default.json # Tauri v2 permission declarations
```

## Prerequisites

| Tool | Version |
|------|---------|
| Rust | ≥ 1.77 |
| Node.js | ≥ 20 |
| npm | ≥ 9 |
| System WebView | Edge (Windows) / WebKit (macOS/Linux) |

On **Linux** you also need the WebKitGTK development libraries:

```bash
# Debian/Ubuntu
sudo apt install libwebkit2gtk-4.1-dev libgtk-3-dev libayatana-appindicator3-dev librsvg2-dev
```

## Development

```bash
cd desktop
npm install
npm run tauri dev
```

## Production Build

```bash
cd desktop
npm install
npm run tauri build
```

The installer/binary is output to `src-tauri/target/release/bundle/`.

## Configuration

On first launch the app asks you to sign in. You can configure the server URL in **Settings → Connection** to point at your LTEdu instance (default: `http://localhost:8080`).

## Tauri Commands (Rust → Frontend bridge)

| Command | Description |
|---------|-------------|
| `cmd_login` | Authenticate against the cloud API and persist the token |
| `cmd_logout` | Clear the persisted session |
| `cmd_get_current_user` | Return the cached token + user profile |
| `cmd_api_request` | Proxy an authenticated HTTP request |
| `cmd_cache_get` | Read a value from the SQLite cache |
| `cmd_cache_set` | Write a value to the SQLite cache |
| `cmd_cache_delete` | Delete a single cache entry |
| `cmd_cache_clear` | Wipe all cached entries |
| `cmd_get_cached_resources` | List all cache entries |

## Success Criteria

- Binary size < 50 MB ✓ (Tauri + stripped release profile)
- Startup time < 2 seconds ✓ (native WebView, no Electron overhead)
- 100% feature parity with web version (via API proxy)
- Native look on all platforms (system WebView, system fonts)
- Seamless cloud synchronisation (direct REST API, local SQLite cache)
