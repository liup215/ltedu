import router from '@ohos.router';
import authService from '../services/AuthService';
import storageService from '../services/StorageService';
import type { User } from '../models/ApiModels';
import { LoadingOverlay, ErrorBanner } from '../components/CommonComponents';

/**
 * ProfilePage shows the authenticated user's profile, settings, and
 * provides sign-out functionality.
 */
@Component
export struct ProfilePage {
  @State user: User | null = null;
  @State isLoading: boolean = false;
  @State errorMessage: string = '';
  @State showLogoutDialog: boolean = false;
  @State apiBaseUrl: string = '';

  aboutToAppear(): void {
    this.loadUser();
    this.loadSettings();
  }

  async loadUser(): Promise<void> {
    this.isLoading = true;
    try {
      const cached = await authService.getCachedUser();
      if (cached) {
        this.user = cached;
      }
      // Refresh from server
      const response = await authService.getCurrentUser();
      if (response.code === 0) {
        this.user = response.data;
        await storageService.setUser(JSON.stringify(response.data));
      }
    } catch {
      // Use cached user if network fails
    } finally {
      this.isLoading = false;
    }
  }

  async loadSettings(): Promise<void> {
    this.apiBaseUrl = await storageService.getApiBaseUrl();
  }

  async logout(): Promise<void> {
    this.isLoading = true;
    try {
      await authService.logout();
      router.replaceUrl({ url: 'pages/LoginPage' });
    } catch {
      this.isLoading = false;
    }
  }

  private getInitials(): string {
    if (!this.user) return '?';
    const name = this.user.username || this.user.email || '';
    return name.substring(0, 2).toUpperCase();
  }

  build() {
    Stack() {
      Scroll() {
        Column() {
          // Header
          Text('Profile')
            .fontSize(20)
            .fontWeight(FontWeight.Bold)
            .fontColor('#1E293B')
            .alignSelf(ItemAlign.Start)
            .padding({ left: 20, right: 20, top: 16, bottom: 16 })
            .width('100%')
            .backgroundColor(Color.White)
            .shadow({ radius: 4, color: 'rgba(0,0,0,0.04)', offsetX: 0, offsetY: 2 })

          ErrorBanner({ message: this.errorMessage, onDismiss: () => { this.errorMessage = ''; } })

          // Avatar + name
          Column() {
            // Avatar circle
            Column() {
              Text(this.getInitials())
                .fontSize(28)
                .fontWeight(FontWeight.Bold)
                .fontColor(Color.White)
            }
            .width(80)
            .height(80)
            .borderRadius(40)
            .backgroundColor('#3B82F6')
            .justifyContent(FlexAlign.Center)
            .alignItems(HorizontalAlign.Center)
            .margin({ bottom: 16 })

            Text(this.user?.username || 'Loading...')
              .fontSize(22)
              .fontWeight(FontWeight.Bold)
              .fontColor('#1E293B')
              .margin({ bottom: 4 })

            Text(this.user?.email || '')
              .fontSize(14)
              .fontColor('#64748B')

            if (this.user?.isAdmin) {
              Text('Admin')
                .fontSize(12)
                .fontColor('#7C3AED')
                .backgroundColor('#EDE9FE')
                .padding({ left: 10, right: 10, top: 4, bottom: 4 })
                .borderRadius(12)
                .fontWeight(FontWeight.SemiBold)
                .margin({ top: 8 })
            }
          }
          .width('100%')
          .padding({ top: 32, bottom: 32 })
          .backgroundColor(Color.White)
          .alignItems(HorizontalAlign.Center)
          .margin({ top: 0, bottom: 20 })

          // Settings section
          Column() {
            Text('Settings')
              .fontSize(14)
              .fontWeight(FontWeight.SemiBold)
              .fontColor('#94A3B8')
              .alignSelf(ItemAlign.Start)
              .margin({ bottom: 8 })
              .padding({ left: 4 })

            // API URL setting
            Column() {
              Row() {
                Text('⚙')
                  .fontSize(18)
                  .margin({ right: 12 })

                Column() {
                  Text('API Server URL')
                    .fontSize(15)
                    .fontColor('#1E293B')
                    .fontWeight(FontWeight.Medium)
                  Text(this.apiBaseUrl)
                    .fontSize(12)
                    .fontColor('#94A3B8')
                    .margin({ top: 2 })
                    .maxLines(1)
                    .textOverflow({ overflow: TextOverflow.Ellipsis })
                }
                .alignItems(HorizontalAlign.Start)
                .layoutWeight(1)
              }
              .width('100%')
              .padding({ top: 4, bottom: 4 })

              TextInput({ text: this.apiBaseUrl, placeholder: 'https://api.ltedu.com' })
                .width('100%')
                .height(44)
                .fontSize(14)
                .backgroundColor('#F8FAFC')
                .borderRadius(8)
                .borderWidth(1)
                .borderColor('#E2E8F0')
                .padding({ left: 12, right: 12 })
                .margin({ top: 8 })
                .onChange(async (value: string) => {
                  this.apiBaseUrl = value;
                  await storageService.setApiBaseUrl(value);
                })
            }
            .width('100%')
            .padding(16)
            .backgroundColor(Color.White)
            .borderRadius(12)
            .margin({ bottom: 12 })
          }
          .padding({ left: 16, right: 16 })
          .margin({ bottom: 20 })

          // About section
          Column() {
            Text('About')
              .fontSize(14)
              .fontWeight(FontWeight.SemiBold)
              .fontColor('#94A3B8')
              .alignSelf(ItemAlign.Start)
              .margin({ bottom: 8 })
              .padding({ left: 4 })

            Column() {
              this.InfoRow('📱', 'Platform', 'HarmonyOS')
              Divider().color('#F1F5F9').height(1)
              this.InfoRow('🔖', 'Version', '1.0.0')
              Divider().color('#F1F5F9').height(1)
              this.InfoRow('🌐', 'Backend', 'LTEdu API v1')
            }
            .width('100%')
            .backgroundColor(Color.White)
            .borderRadius(12)
          }
          .padding({ left: 16, right: 16 })
          .margin({ bottom: 24 })

          // Logout button
          Button('Sign Out')
            .width('calc(100% - 32vp)')
            .height(52)
            .fontSize(16)
            .fontWeight(FontWeight.SemiBold)
            .fontColor('#EF4444')
            .backgroundColor('#FEE2E2')
            .borderRadius(12)
            .margin({ left: 16, right: 16, bottom: 32 })
            .onClick(() => {
              this.showLogoutDialog = true;
            })
        }
        .width('100%')
      }
      .width('100%')
      .height('100%')
      .scrollBar(BarState.Off)
      .backgroundColor('#F8FAFC')

      LoadingOverlay({ visible: this.isLoading, message: 'Loading...' })

      // Logout confirmation dialog
      if (this.showLogoutDialog) {
        Column() {
          Column() {
            Text('Sign Out')
              .fontSize(18)
              .fontWeight(FontWeight.Bold)
              .fontColor('#1E293B')
              .margin({ bottom: 12 })

            Text('Are you sure you want to sign out?')
              .fontSize(14)
              .fontColor('#64748B')
              .textAlign(TextAlign.Center)
              .margin({ bottom: 24 })

            Row() {
              Button('Cancel')
                .layoutWeight(1)
                .height(44)
                .fontSize(15)
                .fontColor('#374151')
                .backgroundColor('#F1F5F9')
                .borderRadius(10)
                .margin({ right: 8 })
                .onClick(() => { this.showLogoutDialog = false; })

              Button('Sign Out')
                .layoutWeight(1)
                .height(44)
                .fontSize(15)
                .fontColor(Color.White)
                .backgroundColor('#EF4444')
                .borderRadius(10)
                .margin({ left: 8 })
                .onClick(() => {
                  this.showLogoutDialog = false;
                  this.logout();
                })
            }
          }
          .padding(24)
          .backgroundColor(Color.White)
          .borderRadius(20)
          .width('80%')
          .alignItems(HorizontalAlign.Center)
        }
        .width('100%')
        .height('100%')
        .backgroundColor('rgba(0,0,0,0.4)')
        .justifyContent(FlexAlign.Center)
        .alignItems(HorizontalAlign.Center)
      }
    }
    .width('100%')
    .height('100%')
  }

  @Builder
  InfoRow(icon: string, label: string, value: string) {
    Row() {
      Text(icon)
        .fontSize(16)
        .margin({ right: 12 })
      Text(label)
        .fontSize(14)
        .fontColor('#374151')
        .layoutWeight(1)
      Text(value)
        .fontSize(14)
        .fontColor('#94A3B8')
    }
    .width('100%')
    .padding({ left: 16, right: 16, top: 14, bottom: 14 })
  }
}
