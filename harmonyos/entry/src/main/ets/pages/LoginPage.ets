import router from '@ohos.router';
import authService from '../services/AuthService';
import { ErrorBanner, LoadingOverlay } from '../components/CommonComponents';

/**
 * LoginPage is the entry point for unauthenticated users.
 * It supports username/password login and navigation to registration.
 */
@Entry
@Component
struct LoginPage {
  @State username: string = '';
  @State password: string = '';
  @State isLoading: boolean = false;
  @State errorMessage: string = '';

  async login(): Promise<void> {
    if (!this.username.trim() || !this.password.trim()) {
      this.errorMessage = 'Please enter your username and password.';
      return;
    }

    this.isLoading = true;
    this.errorMessage = '';

    try {
      const response = await authService.login({
        username: this.username.trim(),
        password: this.password,
      });

      if (response.code === 0) {
        router.replaceUrl({ url: 'pages/MainPage' });
      } else {
        this.errorMessage = response.msg || 'Invalid username or password.';
      }
    } catch (err) {
      this.errorMessage = 'Network error. Please check your connection.';
    } finally {
      this.isLoading = false;
    }
  }

  build() {
    Stack() {
      // Background gradient
      Column()
        .width('100%')
        .height('100%')
        .linearGradient({
          angle: 135,
          colors: [['#EFF6FF', 0.0], ['#F5F3FF', 0.5], ['#F0FDF4', 1.0]]
        })

      // Content
      Scroll() {
        Column() {
          // Logo area
          Column() {
            Text('📚')
              .fontSize(64)
              .margin({ bottom: 16 })

            Text('LTEdu')
              .fontSize(32)
              .fontWeight(FontWeight.Bold)
              .fontColor('#1E293B')
              .margin({ bottom: 8 })

            Text('AI-powered learning platform')
              .fontSize(15)
              .fontColor('#64748B')
          }
          .margin({ top: 60, bottom: 48 })
          .alignItems(HorizontalAlign.Center)

          // Login card
          Column() {
            Text('Sign In')
              .fontSize(22)
              .fontWeight(FontWeight.Bold)
              .fontColor('#1E293B')
              .alignSelf(ItemAlign.Start)
              .margin({ bottom: 24 })

            ErrorBanner({ message: this.errorMessage, onDismiss: () => { this.errorMessage = ''; } })

            // Username input
            Column() {
              Text('Username or Email')
                .fontSize(13)
                .fontColor('#374151')
                .fontWeight(FontWeight.Medium)
                .alignSelf(ItemAlign.Start)
                .margin({ bottom: 6 })

              TextInput({ placeholder: 'Enter your username or email' })
                .width('100%')
                .height(48)
                .fontSize(15)
                .backgroundColor('#F8FAFC')
                .borderRadius(10)
                .borderWidth(1)
                .borderColor('#E2E8F0')
                .padding({ left: 16, right: 16 })
                .onChange((value: string) => {
                  this.username = value;
                })
            }
            .margin({ bottom: 16 })

            // Password input
            Column() {
              Text('Password')
                .fontSize(13)
                .fontColor('#374151')
                .fontWeight(FontWeight.Medium)
                .alignSelf(ItemAlign.Start)
                .margin({ bottom: 6 })

              TextInput({ placeholder: 'Enter your password' })
                .width('100%')
                .height(48)
                .fontSize(15)
                .type(InputType.Password)
                .backgroundColor('#F8FAFC')
                .borderRadius(10)
                .borderWidth(1)
                .borderColor('#E2E8F0')
                .padding({ left: 16, right: 16 })
                .onChange((value: string) => {
                  this.password = value;
                })
            }
            .margin({ bottom: 24 })

            // Login button
            Button('Sign In')
              .width('100%')
              .height(52)
              .fontSize(16)
              .fontWeight(FontWeight.SemiBold)
              .fontColor(Color.White)
              .backgroundColor('#3B82F6')
              .borderRadius(12)
              .onClick(() => {
                this.login();
              })

            // Register link
            Row() {
              Text("Don't have an account? ")
                .fontSize(14)
                .fontColor('#64748B')
              Text('Create one')
                .fontSize(14)
                .fontColor('#3B82F6')
                .fontWeight(FontWeight.Medium)
                .onClick(() => {
                  router.pushUrl({ url: 'pages/RegisterPage' });
                })
            }
            .margin({ top: 20 })
            .justifyContent(FlexAlign.Center)
          }
          .width('100%')
          .padding(24)
          .backgroundColor(Color.White)
          .borderRadius(20)
          .shadow({ radius: 16, color: 'rgba(0,0,0,0.06)', offsetX: 0, offsetY: 4 })
          .margin({ left: 24, right: 24 })
        }
        .width('100%')
        .alignItems(HorizontalAlign.Center)
      }
      .width('100%')
      .height('100%')
      .scrollBar(BarState.Off)

      // Loading overlay
      LoadingOverlay({ visible: this.isLoading, message: 'Signing in...' })
    }
    .width('100%')
    .height('100%')
  }
}
