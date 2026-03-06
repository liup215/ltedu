import preferences from '@ohos.data.preferences';
import common from '@ohos.app.ability.common';

const STORE_NAME = 'ltedu_prefs';
const KEY_TOKEN = 'auth_token';
const KEY_USER = 'current_user';
const KEY_API_BASE = 'api_base_url';
const DEFAULT_API_BASE = 'https://api.ltedu.com';

/**
 * StorageService wraps HarmonyOS Preferences for persistent local storage
 * of auth tokens, user data, and app settings.
 */
export class StorageService {
  private static instance: StorageService;
  private store: preferences.Preferences | null = null;
  private context: common.UIAbilityContext | null = null;

  static getInstance(): StorageService {
    if (!StorageService.instance) {
      StorageService.instance = new StorageService();
    }
    return StorageService.instance;
  }

  async init(context: common.UIAbilityContext): Promise<void> {
    this.context = context;
    this.store = await preferences.getPreferences(context, STORE_NAME);
  }

  async getToken(): Promise<string> {
    if (!this.store) return '';
    return (await this.store.get(KEY_TOKEN, '')) as string;
  }

  async setToken(token: string): Promise<void> {
    if (!this.store) return;
    await this.store.put(KEY_TOKEN, token);
    await this.store.flush();
  }

  async clearToken(): Promise<void> {
    if (!this.store) return;
    await this.store.delete(KEY_TOKEN);
    await this.store.delete(KEY_USER);
    await this.store.flush();
  }

  async getUser(): Promise<string> {
    if (!this.store) return '';
    return (await this.store.get(KEY_USER, '')) as string;
  }

  async setUser(userJson: string): Promise<void> {
    if (!this.store) return;
    await this.store.put(KEY_USER, userJson);
    await this.store.flush();
  }

  async getApiBaseUrl(): Promise<string> {
    if (!this.store) return DEFAULT_API_BASE;
    return (await this.store.get(KEY_API_BASE, DEFAULT_API_BASE)) as string;
  }

  async setApiBaseUrl(url: string): Promise<void> {
    if (!this.store) return;
    await this.store.put(KEY_API_BASE, url);
    await this.store.flush();
  }

  async isLoggedIn(): Promise<boolean> {
    const token = await this.getToken();
    return token.length > 0;
  }
}

export default StorageService.getInstance();
