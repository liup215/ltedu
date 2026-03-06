import router from '@ohos.router';
import authService from '../services/AuthService';
import { ErrorBanner, LoadingOverlay } from '../components/CommonComponents';

/**
 * RegisterPage lets new users create an account with email verification.
 */
@Entry
@Component
struct RegisterPage {
  @State username: string = '';
  @State password: string = '';
  @State confirmPassword: string = '';
  @State email: string = '';
  @State verificationCode: string = '';
  @State isLoading: boolean = false;
  @State errorMessage: string = '';
  @State successMessage: string = '';
  @State codeSent: boolean = false;
  @State codeSending: boolean = false;

  async sendCode(): Promise<void> {
    if (!this.email.trim()) {
      this.errorMessage = 'Please enter your email address first.';
      return;
    }
    this.codeSending = true;
    this.errorMessage = '';
    try {
      const response = await authService.sendVerificationCode({
        email: this.email.trim(),
        captchaId: '',
        captchaValue: '',
      });
      if (response.code === 0) {
        this.codeSent = true;
        this.successMessage = 'Verification code sent to your email.';
      } else {
        this.errorMessage = response.msg || 'Failed to send verification code.';
      }
    } catch {
      this.errorMessage = 'Network error. Please try again.';
    } finally {
      this.codeSending = false;
    }
  }

  async register(): Promise<void> {
    if (!this.username.trim() || !this.password || !this.email.trim() || !this.verificationCode.trim()) {
      this.errorMessage = 'Please fill in all required fields.';
      return;
    }
    if (this.password !== this.confirmPassword) {
      this.errorMessage = 'Passwords do not match.';
      return;
    }

    this.isLoading = true;
    this.errorMessage = '';
    try {
      const response = await authService.register({
        username: this.username.trim(),
        password: this.password,
        email: this.email.trim(),
        code: this.verificationCode.trim(),
      });
      if (response.code === 0) {
        router.replaceUrl({ url: 'pages/LoginPage' });
      } else {
        this.errorMessage = response.msg || 'Registration failed. Please try again.';
      }
    } catch {
      this.errorMessage = 'Network error. Please try again.';
    } finally {
      this.isLoading = false;
    }
  }

  build() {
    Stack() {
      Scroll() {
        Column() {
          // Header
          Row() {
            Image($r('sys.media.ohos_ic_back'))
              .width(24)
              .height(24)
              .onClick(() => router.back())
            Text('Create Account')
              .fontSize(20)
              .fontWeight(FontWeight.Bold)
              .fontColor('#1E293B')
              .margin({ left: 12 })
            Blank()
          }
          .width('100%')
          .padding({ left: 20, right: 20, top: 20, bottom: 24 })

          Column() {
            ErrorBanner({ message: this.errorMessage, onDismiss: () => { this.errorMessage = ''; } })

            if (this.successMessage) {
              Row() {
                Text('✓  ' + this.successMessage)
                  .fontSize(14)
                  .fontColor('#065F46')
              }
              .width('100%')
              .padding({ left: 16, right: 16, top: 12, bottom: 12 })
              .backgroundColor('#D1FAE5')
              .borderRadius(8)
              .margin({ bottom: 12 })
            }

            // Username
            this.InputField('Username', 'Choose a username', InputType.Normal,
              (v: string) => { this.username = v; })

            // Email + code button
            Column() {
              Text('Email')
                .fontSize(13)
                .fontColor('#374151')
                .fontWeight(FontWeight.Medium)
                .alignSelf(ItemAlign.Start)
                .margin({ bottom: 6 })
              Row() {
                TextInput({ placeholder: 'Enter your email' })
                  .layoutWeight(1)
                  .height(48)
                  .fontSize(15)
                  .backgroundColor('#F8FAFC')
                  .borderRadius(10)
                  .borderWidth(1)
                  .borderColor('#E2E8F0')
                  .padding({ left: 16, right: 16 })
                  .onChange((value: string) => { this.email = value; })

                Button(this.codeSending ? '...' : 'Send Code')
                  .height(48)
                  .fontSize(13)
                  .fontWeight(FontWeight.Medium)
                  .backgroundColor(this.codeSent ? '#22C55E' : '#3B82F6')
                  .fontColor(Color.White)
                  .borderRadius(10)
                  .margin({ left: 8 })
                  .padding({ left: 14, right: 14 })
                  .onClick(() => { this.sendCode(); })
                  .enabled(!this.codeSending)
              }
            }
            .margin({ bottom: 16 })

            // Verification code
            this.InputField('Verification Code', 'Enter code from email', InputType.Number,
              (v: string) => { this.verificationCode = v; })

            // Password
            this.InputField('Password', 'Choose a password', InputType.Password,
              (v: string) => { this.password = v; })

            // Confirm password
            this.InputField('Confirm Password', 'Repeat your password', InputType.Password,
              (v: string) => { this.confirmPassword = v; })

            Button('Create Account')
              .width('100%')
              .height(52)
              .fontSize(16)
              .fontWeight(FontWeight.SemiBold)
              .fontColor(Color.White)
              .backgroundColor('#3B82F6')
              .borderRadius(12)
              .margin({ top: 8 })
              .onClick(() => { this.register(); })
          }
          .width('100%')
          .padding(24)
          .backgroundColor(Color.White)
          .borderRadius(20)
          .shadow({ radius: 16, color: 'rgba(0,0,0,0.06)', offsetX: 0, offsetY: 4 })
          .margin({ left: 24, right: 24, bottom: 24 })
        }
        .width('100%')
      }
      .width('100%')
      .height('100%')
      .scrollBar(BarState.Off)

      LoadingOverlay({ visible: this.isLoading, message: 'Creating account...' })
    }
    .width('100%')
    .height('100%')
    .backgroundColor('#F8FAFC')
  }

  @Builder
  InputField(label: string, placeholder: string, type: InputType, onChange: (v: string) => void) {
    Column() {
      Text(label)
        .fontSize(13)
        .fontColor('#374151')
        .fontWeight(FontWeight.Medium)
        .alignSelf(ItemAlign.Start)
        .margin({ bottom: 6 })
      TextInput({ placeholder: placeholder })
        .width('100%')
        .height(48)
        .fontSize(15)
        .type(type)
        .backgroundColor('#F8FAFC')
        .borderRadius(10)
        .borderWidth(1)
        .borderColor('#E2E8F0')
        .padding({ left: 16, right: 16 })
        .onChange(onChange)
    }
    .margin({ bottom: 16 })
  }
}
