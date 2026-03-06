import UIAbility from '@ohos.app.ability.UIAbility';
import type AbilityConstant from '@ohos.app.ability.AbilityConstant';
import type Want from '@ohos.app.ability.Want';
import type window from '@ohos.window';
import storageService from '../services/StorageService';

/**
 * EntryAbility is the main entry point for the LTEdu HarmonyOS app.
 * It initialises StorageService on startup and routes the user to
 * either the login page or the main page depending on auth state.
 */
export default class EntryAbility extends UIAbility {
  async onCreate(want: Want, launchParam: AbilityConstant.LaunchParam): Promise<void> {
    await storageService.init(this.context);
  }

  onDestroy(): void {
    // Clean-up on ability destruction
  }

  async onWindowStageCreate(windowStage: window.WindowStage): Promise<void> {
    const isLoggedIn = await storageService.isLoggedIn();
    const initialPage = isLoggedIn ? 'pages/MainPage' : 'pages/LoginPage';

    windowStage.loadContent(initialPage, (err) => {
      if (err.code) {
        // Fall back to login page if loading fails
        windowStage.loadContent('pages/LoginPage');
      }
    });
  }

  onWindowStageDestroy(): void {
    // Window stage is about to be destroyed
  }

  onForeground(): void {
    // Ability becomes foreground
  }

  onBackground(): void {
    // Ability enters background
  }
}
