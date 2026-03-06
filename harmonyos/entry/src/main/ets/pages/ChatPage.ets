import aiService from '../services/AiService';
import type { ChatMessage } from '../models/ApiModels';
import { ChatMessageItem } from '../components/ChatMessageItem';
import { LoadingOverlay } from '../components/CommonComponents';

/**
 * ChatPage is the primary AI-native interaction screen.
 * Users can ask questions about their studies, get AI-powered explanations,
 * and receive personalised learning guidance.
 */
@Component
export struct ChatPage {
  @State messages: ChatMessage[] = [];
  @State inputText: string = '';
  @State isAiTyping: boolean = false;
  private scroller: Scroller = new Scroller();

  private readonly WELCOME_MESSAGE: string =
    "Hello! I'm your LTEdu AI assistant 👋\n\n" +
    "I can help you with:\n" +
    "• Explaining difficult concepts\n" +
    "• Reviewing your answers\n" +
    "• Building study strategies\n" +
    "• Practice question guidance\n\n" +
    "What would you like to learn today?";

  aboutToAppear(): void {
    const welcome: ChatMessage = {
      id: 'welcome',
      role: 'assistant',
      content: this.WELCOME_MESSAGE,
      timestamp: Date.now(),
    };
    this.messages = [welcome];
  }

  private generateId(): string {
    return `msg_${Date.now()}_${Math.random().toString(36).substring(2, 9)}`;
  }

  async sendMessage(): Promise<void> {
    const text = this.inputText.trim();
    if (!text || this.isAiTyping) return;

    const userMsg: ChatMessage = {
      id: this.generateId(),
      role: 'user',
      content: text,
      timestamp: Date.now(),
    };

    this.messages = [...this.messages, userMsg];
    this.inputText = '';
    this.isAiTyping = true;

    // Scroll to bottom
    this.scroller.scrollEdge(Edge.Bottom);

    try {
      const reply = await aiService.sendChatMessage(text);
      const aiMsg: ChatMessage = {
        id: this.generateId(),
        role: 'assistant',
        content: reply,
        timestamp: Date.now(),
      };
      this.messages = [...this.messages, aiMsg];
    } finally {
      this.isAiTyping = false;
      this.scroller.scrollEdge(Edge.Bottom);
    }
  }

  build() {
    Column() {
      // Header
      Row() {
        Column() {
          Text('AI Assistant')
            .fontSize(20)
            .fontWeight(FontWeight.Bold)
            .fontColor('#1E293B')
          Text('Powered by LTEdu AI')
            .fontSize(12)
            .fontColor('#94A3B8')
            .margin({ top: 2 })
        }
        .alignItems(HorizontalAlign.Start)

        Blank()

        Button('Clear')
          .fontSize(13)
          .fontColor('#64748B')
          .backgroundColor('#F1F5F9')
          .borderRadius(8)
          .height(32)
          .padding({ left: 12, right: 12 })
          .onClick(() => {
            aiService.clearHistory();
            this.messages = [];
            this.aboutToAppear();
          })
      }
      .width('100%')
      .padding({ left: 20, right: 20, top: 16, bottom: 12 })
      .backgroundColor(Color.White)
      .shadow({ radius: 4, color: 'rgba(0,0,0,0.04)', offsetX: 0, offsetY: 2 })

      // Message list
      Scroll(this.scroller) {
        Column() {
          ForEach(this.messages, (msg: ChatMessage) => {
            ChatMessageItem({ message: msg })
          }, (msg: ChatMessage) => msg.id)

          // Typing indicator
          if (this.isAiTyping) {
            Row() {
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
              .margin({ right: 8 })

              Row() {
                Text('● ● ●')
                  .fontSize(16)
                  .fontColor('#94A3B8')
                  .letterSpacing(4)
              }
              .padding({ left: 14, right: 14, top: 10, bottom: 10 })
              .backgroundColor('#F1F5F9')
              .borderRadius(16)
            }
            .width('100%')
            .padding({ left: 16, right: 16, top: 6, bottom: 6 })
            .alignItems(VerticalAlign.Top)
          }
        }
        .width('100%')
        .padding({ bottom: 16 })
      }
      .layoutWeight(1)
      .scrollBar(BarState.Off)
      .backgroundColor('#F8FAFC')

      // Input area
      Row() {
        TextInput({ placeholder: 'Ask me anything...', text: this.inputText })
          .layoutWeight(1)
          .height(44)
          .fontSize(15)
          .backgroundColor('#F1F5F9')
          .borderRadius(22)
          .padding({ left: 16, right: 16 })
          .borderWidth(0)
          .onChange((value: string) => {
            this.inputText = value;
          })
          .onSubmit(() => {
            this.sendMessage();
          })

        Button() {
          Text('➤')
            .fontSize(18)
            .fontColor(Color.White)
        }
        .width(44)
        .height(44)
        .borderRadius(22)
        .backgroundColor(this.inputText.trim() ? '#3B82F6' : '#CBD5E1')
        .margin({ left: 10 })
        .enabled(this.inputText.trim().length > 0 && !this.isAiTyping)
        .onClick(() => {
          this.sendMessage();
        })
      }
      .width('100%')
      .padding({ left: 16, right: 16, top: 10, bottom: 10 })
      .backgroundColor(Color.White)
      .shadow({ radius: 4, color: 'rgba(0,0,0,0.06)', offsetX: 0, offsetY: -2 })
    }
    .width('100%')
    .height('100%')
    .backgroundColor('#F8FAFC')
  }
}
