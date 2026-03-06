import { ChatPage as ChatPageComponent } from './ChatPage';

/**
 * Standalone entry point for the Chat page (used by deep links and widget actions).
 */
@Entry
@Component
struct ChatPageEntry {
  build() {
    ChatPageComponent()
  }
}
