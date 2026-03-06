import type { ChatMessage } from '../models/ApiModels';

/**
 * ChatMessageItem renders a single chat message bubble.
 * User messages appear on the right (blue), AI messages on the left (grey).
 */
@Component
export struct ChatMessageItem {
  @Prop message: ChatMessage = {
    id: '',
    role: 'user',
    content: '',
    timestamp: 0
  };

  private formatTime(timestamp: number): string {
    const date = new Date(timestamp);
    const h = date.getHours().toString().padStart(2, '0');
    const m = date.getMinutes().toString().padStart(2, '0');
    return `${h}:${m}`;
  }

  build() {
    Row() {
      if (this.message.role === 'assistant') {
        // AI avatar
        Column() {
          Text('AI')
            .fontSize(12)
            .fontColor(Color.White)
            .fontWeight(FontWeight.Bold)
        }
        .width(36)
        .height(36)
        .borderRadius(18)
        .backgroundColor('#6366F1')
        .justifyContent(FlexAlign.Center)
        .alignItems(HorizontalAlign.Center)
        .margin({ right: 8, top: 4 })

        Column() {
          Text(this.message.content)
            .fontSize(15)
            .fontColor('#1E293B')
            .lineHeight(22)
          Text(this.formatTime(this.message.timestamp))
            .fontSize(11)
            .fontColor('#94A3B8')
            .margin({ top: 4 })
            .alignSelf(ItemAlign.End)
        }
        .padding({ left: 12, right: 12, top: 10, bottom: 10 })
        .backgroundColor('#F1F5F9')
        .borderRadius({ topLeft: 4, topRight: 16, bottomLeft: 16, bottomRight: 16 })
        .constraintSize({ maxWidth: '72%' })
        .alignItems(HorizontalAlign.Start)

        Blank()
      } else {
        Blank()
        Column() {
          Text(this.message.content)
            .fontSize(15)
            .fontColor(Color.White)
            .lineHeight(22)
          Text(this.formatTime(this.message.timestamp))
            .fontSize(11)
            .fontColor('rgba(255,255,255,0.7)')
            .margin({ top: 4 })
            .alignSelf(ItemAlign.End)
        }
        .padding({ left: 12, right: 12, top: 10, bottom: 10 })
        .backgroundColor('#3B82F6')
        .borderRadius({ topLeft: 16, topRight: 4, bottomLeft: 16, bottomRight: 16 })
        .constraintSize({ maxWidth: '72%' })
        .alignItems(HorizontalAlign.End)

        // User avatar
        Column() {
          Text('Me')
            .fontSize(12)
            .fontColor(Color.White)
            .fontWeight(FontWeight.Bold)
        }
        .width(36)
        .height(36)
        .borderRadius(18)
        .backgroundColor('#3B82F6')
        .justifyContent(FlexAlign.Center)
        .alignItems(HorizontalAlign.Center)
        .margin({ left: 8, top: 4 })
      }
    }
    .width('100%')
    .padding({ left: 16, right: 16, top: 6, bottom: 6 })
    .alignItems(VerticalAlign.Top)
  }
}
