export interface Config {
  API_BASE_URL: string;
  APP_TITLE: string;
  // Add other config properties as needed
}

/**
 * Fetch config.json at runtime and return a Config object.
 * Usage: const config = await getConfig();
 */
export async function getConfig(): Promise<Config> {
  const response = await fetch('/config.json');
  const config = await response.json();
  return {
    API_BASE_URL: config.API_BASE_URL || 'http://api.example.com',
    APP_TITLE: config.APP_TITLE || 'LTEDU',
  };
}

export const API_BASE_URL = async () => {
  const config = await getConfig();
  return config.API_BASE_URL || 'http://api.example.com'; // Fallback URL
}

export const APP_TITLE = async () => {
  const config = await getConfig();
  return config.APP_TITLE || 'LTEDU'; // Fallback title
};
