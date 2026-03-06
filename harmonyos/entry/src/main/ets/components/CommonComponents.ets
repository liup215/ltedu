/**
 * LoadingOverlay displays a full-screen loading indicator with an optional message.
 */
@Component
export struct LoadingOverlay {
  @Prop message: string = 'Loading...';
  @Prop visible: boolean = false;

  build() {
    if (this.visible) {
      Stack() {
        Column() {
          LoadingProgress()
            .width(48)
            .height(48)
            .color('#3B82F6')
          Text(this.message)
            .fontSize(14)
            .fontColor('#64748B')
            .margin({ top: 12 })
        }
        .padding(24)
        .backgroundColor(Color.White)
        .borderRadius(16)
        .shadow({ radius: 16, color: 'rgba(0,0,0,0.1)', offsetX: 0, offsetY: 4 })
      }
      .width('100%')
      .height('100%')
      .backgroundColor('rgba(0,0,0,0.3)')
      .justifyContent(FlexAlign.Center)
      .alignItems(HorizontalAlign.Center)
    }
  }
}

/**
 * ErrorBanner displays an inline error message with a dismiss button.
 */
@Component
export struct ErrorBanner {
  @Prop message: string = '';
  onDismiss: () => void = () => {};

  build() {
    if (this.message) {
      Row() {
        Text('⚠')
          .fontSize(16)
          .margin({ right: 8 })

        Text(this.message)
          .fontSize(14)
          .fontColor('#991B1B')
          .layoutWeight(1)
          .lineHeight(20)

        Text('✕')
          .fontSize(14)
          .fontColor('#991B1B')
          .margin({ left: 8 })
          .onClick(() => {
            this.onDismiss();
          })
      }
      .width('100%')
      .padding({ left: 16, right: 16, top: 12, bottom: 12 })
      .backgroundColor('#FEE2E2')
      .borderRadius(8)
      .margin({ bottom: 12 })
    }
  }
}

/**
 * EmptyState shows a friendly empty state with an icon and message.
 */
@Component
export struct EmptyState {
  @Prop icon: string = '📭';
  @Prop title: string = 'Nothing here yet';
  @Prop subtitle: string = '';

  build() {
    Column() {
      Text(this.icon)
        .fontSize(48)
        .margin({ bottom: 12 })
      Text(this.title)
        .fontSize(16)
        .fontWeight(FontWeight.Medium)
        .fontColor('#374151')
        .margin({ bottom: 6 })
      if (this.subtitle) {
        Text(this.subtitle)
          .fontSize(13)
          .fontColor('#94A3B8')
          .textAlign(TextAlign.Center)
      }
    }
    .width('100%')
    .padding(40)
    .justifyContent(FlexAlign.Center)
    .alignItems(HorizontalAlign.Center)
  }
}
