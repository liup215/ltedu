/**
 * WidgetCard is the Atomic Service widget for LTEdu.
 * It provides quick-access to the AI assistant and study progress
 * directly from the HarmonyOS home screen.
 */
@Entry
@Component
struct WidgetCard {
  @LocalStorageProp('title') title: string = 'LTEdu';
  @LocalStorageProp('subtitle') subtitle: string = 'AI Learning Assistant';
  @LocalStorageProp('status') status: string = 'Ready to learn';

  build() {
    Column() {
      Row() {
        Text('📚')
          .fontSize(28)
          .margin({ right: 10 })

        Column() {
          Text(this.title)
            .fontSize(16)
            .fontWeight(FontWeight.Bold)
            .fontColor('#1E293B')
          Text(this.subtitle)
            .fontSize(11)
            .fontColor('#64748B')
            .margin({ top: 2 })
        }
        .alignItems(HorizontalAlign.Start)
        .layoutWeight(1)
      }
      .width('100%')
      .margin({ bottom: 12 })

      Text(this.status)
        .fontSize(13)
        .fontColor('#374151')
        .width('100%')
        .backgroundColor('#EFF6FF')
        .borderRadius(8)
        .padding({ left: 12, right: 12, top: 8, bottom: 8 })

      Row() {
        Button('Chat')
          .height(32)
          .fontSize(13)
          .fontWeight(FontWeight.Medium)
          .fontColor('#3B82F6')
          .backgroundColor('#DBEAFE')
          .borderRadius(8)
          .padding({ left: 16, right: 16 })
          .onClick(() => {
            postCardAction(this, {
              action: 'router',
              abilityName: 'EntryAbility',
              params: { page: 'chat' }
            });
          })

        Blank()

        Button('Practice')
          .height(32)
          .fontSize(13)
          .fontWeight(FontWeight.Medium)
          .fontColor('#7C3AED')
          .backgroundColor('#EDE9FE')
          .borderRadius(8)
          .padding({ left: 16, right: 16 })
          .onClick(() => {
            postCardAction(this, {
              action: 'router',
              abilityName: 'EntryAbility',
              params: { page: 'practice' }
            });
          })
      }
      .width('100%')
      .margin({ top: 12 })
    }
    .width('100%')
    .height('100%')
    .padding(16)
    .backgroundColor(Color.White)
    .borderRadius(16)
  }
}
