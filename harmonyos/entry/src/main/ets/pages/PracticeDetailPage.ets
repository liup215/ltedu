import router from '@ohos.router';
import practiceService from '../services/PracticeService';
import aiService from '../services/AiService';
import type { Question } from '../models/ApiModels';
import { QuestionCard } from '../components/QuestionCard';
import { LoadingOverlay, ErrorBanner } from '../components/CommonComponents';

/**
 * PracticeDetailPage shows a full paper-based practice session with AI feedback.
 * Receives paperId as a router parameter.
 */
@Entry
@Component
struct PracticeDetailPage {
  @State questions: Question[] = [];
  @State answers: Map<number, string> = new Map();
  @State isLoading: boolean = false;
  @State errorMessage: string = '';
  @State aiFeedback: string = '';
  @State currentQuestionIndex: number = 0;
  private paperId: number = 0;

  aboutToAppear(): void {
    const params = router.getParams() as Record<string, number>;
    this.paperId = params?.paperId || 0;
    if (this.paperId > 0) {
      this.loadPaper();
    }
  }

  async loadPaper(): Promise<void> {
    this.isLoading = true;
    this.errorMessage = '';
    try {
      const response = await practiceService.paperPractice({ paperId: this.paperId });
      if (response.code === 0 && response.data?.questions) {
        this.questions = response.data.questions;
      } else {
        this.errorMessage = response.msg || 'Failed to load paper.';
      }
    } catch {
      this.errorMessage = 'Network error. Please check your connection.';
    } finally {
      this.isLoading = false;
    }
  }

  async getAiFeedback(): Promise<void> {
    const currentQ = this.questions[this.currentQuestionIndex];
    if (!currentQ) return;

    const userAnswer = this.answers.get(currentQ.id) || '';
    if (!userAnswer) {
      this.aiFeedback = 'Please answer the question first to get AI feedback.';
      return;
    }

    this.isLoading = true;
    this.aiFeedback = '';
    try {
      const feedback = await aiService.sendChatMessage(
        `I answered "${userAnswer}" for: "${currentQ.stem}". The correct answer is: "${currentQ.answer || 'not specified'}". Please give me feedback.`
      );
      this.aiFeedback = feedback;
    } finally {
      this.isLoading = false;
    }
  }

  build() {
    Stack() {
      Column() {
        // Header
        Row() {
          Text('←')
            .fontSize(20)
            .fontColor('#3B82F6')
            .onClick(() => router.back())
          Text('Paper Practice')
            .fontSize(18)
            .fontWeight(FontWeight.Bold)
            .fontColor('#1E293B')
            .margin({ left: 12 })
          Blank()
          Text(`${this.currentQuestionIndex + 1}/${this.questions.length}`)
            .fontSize(14)
            .fontColor('#64748B')
        }
        .width('100%')
        .padding({ left: 20, right: 20, top: 16, bottom: 12 })
        .backgroundColor(Color.White)

        ErrorBanner({ message: this.errorMessage, onDismiss: () => { this.errorMessage = ''; } })

        Scroll() {
          Column() {
            if (this.questions.length > 0) {
              QuestionCard({
                question: this.questions[this.currentQuestionIndex],
                index: this.currentQuestionIndex,
                selectedAnswer: this.answers.get(this.questions[this.currentQuestionIndex]?.id) || '',
                onAnswerSelected: (questionId: number, answer: string) => {
                  const newAnswers = new Map(this.answers);
                  newAnswers.set(questionId, answer);
                  this.answers = newAnswers;
                }
              })

              // AI feedback section
              if (this.aiFeedback) {
                Column() {
                  Row() {
                    Text('🤖')
                      .fontSize(16)
                      .margin({ right: 8 })
                    Text('AI Feedback')
                      .fontSize(14)
                      .fontWeight(FontWeight.SemiBold)
                      .fontColor('#5B21B6')
                  }
                  .margin({ bottom: 10 })

                  Text(this.aiFeedback)
                    .fontSize(14)
                    .fontColor('#374151')
                    .lineHeight(22)
                    .width('100%')
                }
                .width('100%')
                .padding(16)
                .backgroundColor('#F5F3FF')
                .borderRadius(12)
                .borderWidth(1)
                .borderColor('#DDD6FE')
                .margin({ bottom: 12 })
              }

              Button('Get AI Feedback')
                .width('100%')
                .height(44)
                .fontSize(14)
                .fontWeight(FontWeight.Medium)
                .fontColor('#7C3AED')
                .backgroundColor('#EDE9FE')
                .borderRadius(10)
                .margin({ bottom: 12 })
                .onClick(() => { this.getAiFeedback(); })
            }
          }
          .padding({ left: 16, right: 16, top: 16, bottom: 16 })
        }
        .layoutWeight(1)
        .scrollBar(BarState.Off)

        // Navigation
        Row() {
          Button('← Prev')
            .height(44)
            .fontSize(14)
            .fontColor(this.currentQuestionIndex > 0 ? '#3B82F6' : '#CBD5E1')
            .backgroundColor(this.currentQuestionIndex > 0 ? '#EFF6FF' : '#F8FAFC')
            .borderRadius(10)
            .padding({ left: 20, right: 20 })
            .enabled(this.currentQuestionIndex > 0)
            .onClick(() => {
              this.currentQuestionIndex -= 1;
              this.aiFeedback = '';
            })

          Blank()

          Button('Next →')
            .height(44)
            .fontSize(14)
            .fontColor(this.currentQuestionIndex < this.questions.length - 1 ? '#3B82F6' : '#CBD5E1')
            .backgroundColor(this.currentQuestionIndex < this.questions.length - 1 ? '#EFF6FF' : '#F8FAFC')
            .borderRadius(10)
            .padding({ left: 20, right: 20 })
            .enabled(this.currentQuestionIndex < this.questions.length - 1)
            .onClick(() => {
              this.currentQuestionIndex += 1;
              this.aiFeedback = '';
            })
        }
        .width('100%')
        .padding({ left: 16, right: 16, top: 10, bottom: 16 })
        .backgroundColor(Color.White)
      }
      .width('100%')
      .height('100%')
      .backgroundColor('#F8FAFC')

      LoadingOverlay({ visible: this.isLoading, message: 'Loading...' })
    }
    .width('100%')
    .height('100%')
  }
}
