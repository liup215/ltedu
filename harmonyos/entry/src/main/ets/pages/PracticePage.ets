import router from '@ohos.router';
import practiceService from '../services/PracticeService';
import type { Question } from '../models/ApiModels';
import { QuestionCard } from '../components/QuestionCard';
import { LoadingOverlay, ErrorBanner, EmptyState } from '../components/CommonComponents';

/**
 * PracticePage lists available practice modes and loads a set of questions
 * for the user to answer. Supports quick practice and answer submission.
 */
@Component
export struct PracticePage {
  @State questions: Question[] = [];
  @State answers: Map<number, string> = new Map();
  @State isLoading: boolean = false;
  @State errorMessage: string = '';
  @State mode: 'select' | 'practice' | 'results' = 'select';
  @State score: number = 0;
  @State totalAnswered: number = 0;

  async startQuickPractice(count: number): Promise<void> {
    this.isLoading = true;
    this.errorMessage = '';
    this.answers = new Map();
    try {
      const response = await practiceService.quickPractice({ count });
      if (response.code === 0 && response.data?.questions) {
        this.questions = response.data.questions;
        this.mode = 'practice';
      } else {
        this.errorMessage = response.msg || 'Failed to load questions.';
      }
    } catch {
      this.errorMessage = 'Network error. Please check your connection.';
    } finally {
      this.isLoading = false;
    }
  }

  async submitAnswers(): Promise<void> {
    if (this.answers.size === 0) {
      this.errorMessage = 'Please answer at least one question.';
      return;
    }

    this.isLoading = true;
    const answersObj: Record<string, string> = {};
    this.answers.forEach((value, key) => {
      answersObj[key.toString()] = value;
    });

    try {
      const response = await practiceService.gradeAnswers({ answers: answersObj });
      if (response.code === 0) {
        const results = response.data as Record<string, boolean>;
        let correct = 0;
        Object.values(results).forEach(v => { if (v) correct++; });
        this.score = correct;
        this.totalAnswered = Object.keys(results).length;
        this.mode = 'results';
      } else {
        this.errorMessage = response.msg || 'Failed to submit answers.';
      }
    } catch {
      this.errorMessage = 'Network error. Please check your connection.';
    } finally {
      this.isLoading = false;
    }
  }

  build() {
    Stack() {
      Column() {
        // Header
        Row() {
          if (this.mode !== 'select') {
            Text('←')
              .fontSize(20)
              .fontColor('#3B82F6')
              .margin({ right: 8 })
              .onClick(() => {
                this.mode = 'select';
                this.questions = [];
                this.errorMessage = '';
              })
          }
          Text(this.mode === 'select' ? 'Practice' :
            this.mode === 'practice' ? `Questions (${this.questions.length})` : 'Results')
            .fontSize(20)
            .fontWeight(FontWeight.Bold)
            .fontColor('#1E293B')
          Blank()
        }
        .width('100%')
        .padding({ left: 20, right: 20, top: 16, bottom: 12 })
        .backgroundColor(Color.White)
        .shadow({ radius: 4, color: 'rgba(0,0,0,0.04)', offsetX: 0, offsetY: 2 })

        ErrorBanner({ message: this.errorMessage, onDismiss: () => { this.errorMessage = ''; } })

        if (this.mode === 'select') {
          this.SelectModeView()
        } else if (this.mode === 'practice') {
          this.PracticeView()
        } else {
          this.ResultsView()
        }
      }
      .width('100%')
      .height('100%')
      .backgroundColor('#F8FAFC')

      LoadingOverlay({ visible: this.isLoading, message: 'Loading...' })
    }
    .width('100%')
    .height('100%')
  }

  @Builder
  SelectModeView() {
    Scroll() {
      Column() {
        Text('Choose a Practice Mode')
          .fontSize(16)
          .fontColor('#64748B')
          .margin({ top: 24, bottom: 20 })

        this.PracticeModeCard('⚡', 'Quick Practice', '10 random questions', '#DBEAFE', '#1D4ED8', () => {
          this.startQuickPractice(10);
        })
        this.PracticeModeCard('🎯', 'Deep Practice', '20 questions', '#EDE9FE', '#5B21B6', () => {
          this.startQuickPractice(20);
        })
        this.PracticeModeCard('📚', 'Mini Session', '5 quick questions', '#D1FAE5', '#065F46', () => {
          this.startQuickPractice(5);
        })
      }
      .width('100%')
      .padding({ left: 20, right: 20, bottom: 20 })
      .alignItems(HorizontalAlign.Center)
    }
    .layoutWeight(1)
    .scrollBar(BarState.Off)
  }

  @Builder
  PracticeModeCard(icon: string, title: string, subtitle: string,
    bgColor: string, textColor: string, onTap: () => void) {
    Row() {
      Text(icon)
        .fontSize(32)
        .margin({ right: 16 })

      Column() {
        Text(title)
          .fontSize(17)
          .fontWeight(FontWeight.SemiBold)
          .fontColor('#1E293B')
          .margin({ bottom: 4 })
        Text(subtitle)
          .fontSize(13)
          .fontColor('#64748B')
      }
      .alignItems(HorizontalAlign.Start)
      .layoutWeight(1)

      Text('→')
        .fontSize(18)
        .fontColor(textColor)
    }
    .width('100%')
    .padding(20)
    .backgroundColor(Color.White)
    .borderRadius(16)
    .shadow({ radius: 6, color: 'rgba(0,0,0,0.06)', offsetX: 0, offsetY: 2 })
    .margin({ bottom: 12 })
    .onClick(onTap)
  }

  @Builder
  PracticeView() {
    Column() {
      Scroll() {
        Column() {
          if (this.questions.length === 0) {
            EmptyState({ icon: '📭', title: 'No questions available', subtitle: 'Try a different practice mode' })
          } else {
            ForEach(this.questions, (q: Question, index: number) => {
              QuestionCard({
                question: q,
                index: index,
                selectedAnswer: this.answers.get(q.id) || '',
                onAnswerSelected: (questionId: number, answer: string) => {
                  const newAnswers = new Map(this.answers);
                  newAnswers.set(questionId, answer);
                  this.answers = newAnswers;
                }
              })
            }, (q: Question) => q.id.toString())
          }
        }
        .padding({ left: 16, right: 16, top: 16, bottom: 16 })
      }
      .layoutWeight(1)
      .scrollBar(BarState.Off)

      // Submit button
      Button(`Submit (${this.answers.size}/${this.questions.length} answered)`)
        .width('calc(100% - 32vp)')
        .height(52)
        .fontSize(16)
        .fontWeight(FontWeight.SemiBold)
        .fontColor(Color.White)
        .backgroundColor(this.answers.size > 0 ? '#3B82F6' : '#CBD5E1')
        .borderRadius(12)
        .margin({ left: 16, right: 16, top: 8, bottom: 16 })
        .enabled(this.answers.size > 0)
        .onClick(() => { this.submitAnswers(); })
    }
    .layoutWeight(1)
  }

  @Builder
  ResultsView() {
    Scroll() {
      Column() {
        // Score card
        Column() {
          Text(this.score >= this.totalAnswered * 0.8 ? '🎉' : this.score >= this.totalAnswered * 0.6 ? '👍' : '💪')
            .fontSize(56)
            .margin({ bottom: 16 })

          Text(`${this.score} / ${this.totalAnswered}`)
            .fontSize(40)
            .fontWeight(FontWeight.Bold)
            .fontColor('#1E293B')
            .margin({ bottom: 8 })

          Text(`${Math.round((this.score / this.totalAnswered) * 100)}% correct`)
            .fontSize(16)
            .fontColor('#64748B')
            .margin({ bottom: 4 })

          Text(this.score >= this.totalAnswered * 0.8 ? 'Excellent work! Keep it up!' :
            this.score >= this.totalAnswered * 0.6 ? 'Good job! Room for improvement.' :
              'Keep practicing, you\'ll get there!')
            .fontSize(15)
            .fontColor('#374151')
            .textAlign(TextAlign.Center)
            .margin({ top: 16 })
        }
        .width('100%')
        .padding(32)
        .backgroundColor(Color.White)
        .borderRadius(20)
        .shadow({ radius: 8, color: 'rgba(0,0,0,0.08)', offsetX: 0, offsetY: 4 })
        .margin({ top: 24, bottom: 24 })
        .alignItems(HorizontalAlign.Center)

        Button('Practice Again')
          .width('100%')
          .height(52)
          .fontSize(16)
          .fontWeight(FontWeight.SemiBold)
          .fontColor(Color.White)
          .backgroundColor('#3B82F6')
          .borderRadius(12)
          .margin({ bottom: 12 })
          .onClick(() => { this.mode = 'select'; })

        Button('Back to Home')
          .width('100%')
          .height(52)
          .fontSize(16)
          .fontWeight(FontWeight.Medium)
          .fontColor('#3B82F6')
          .backgroundColor('#EFF6FF')
          .borderRadius(12)
          .onClick(() => { this.mode = 'select'; })
      }
      .padding({ left: 20, right: 20, bottom: 24 })
    }
    .layoutWeight(1)
    .scrollBar(BarState.Off)
  }
}
