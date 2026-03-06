import type { Question } from '../models/ApiModels';

/**
 * QuestionCard renders a single practice question with its options
 * (for single/multiple choice) or an answer input area (for short answer).
 */
@Component
export struct QuestionCard {
  @Prop question: Question = {
    id: 0,
    type: 1,
    stem: '',
  };
  @Prop index: number = 0;
  @Prop selectedAnswer: string = '';
  onAnswerSelected: (questionId: number, answer: string) => void = () => {};

  private getTypeLabel(): string {
    switch (this.question.type) {
      case 1: return 'Single Choice';
      case 2: return 'Multiple Choice';
      case 3: return 'True/False';
      case 5: return 'Vocabulary';
      case 6: return 'Short Answer';
      default: return 'Question';
    }
  }

  private getTypeBgColor(): string {
    switch (this.question.type) {
      case 1: return '#DBEAFE';
      case 2: return '#EDE9FE';
      case 3: return '#D1FAE5';
      case 5: return '#FEF3C7';
      case 6: return '#FCE7F3';
      default: return '#F1F5F9';
    }
  }

  private getTypeTextColor(): string {
    switch (this.question.type) {
      case 1: return '#1D4ED8';
      case 2: return '#5B21B6';
      case 3: return '#065F46';
      case 5: return '#92400E';
      case 6: return '#9D174D';
      default: return '#374151';
    }
  }

  build() {
    Column() {
      // Question header
      Row() {
        Text(`Q${this.index + 1}`)
          .fontSize(14)
          .fontWeight(FontWeight.Bold)
          .fontColor('#1E293B')

        Text(this.getTypeLabel())
          .fontSize(12)
          .fontColor(this.getTypeTextColor())
          .backgroundColor(this.getTypeBgColor())
          .padding({ left: 8, right: 8, top: 2, bottom: 2 })
          .borderRadius(4)
          .margin({ left: 8 })
      }
      .width('100%')
      .margin({ bottom: 12 })

      // Question stem
      Text(this.question.stem)
        .fontSize(15)
        .fontColor('#1E293B')
        .lineHeight(22)
        .width('100%')
        .margin({ bottom: 16 })

      // Options (for choice questions)
      if (this.question.options && this.question.options.length > 0) {
        ForEach(this.question.options, (option: { label: string; content: string }) => {
          Row() {
            Row() {
              Text(option.label)
                .fontSize(13)
                .fontColor(this.selectedAnswer === option.label ? Color.White : '#3B82F6')
                .fontWeight(FontWeight.Bold)
            }
            .width(28)
            .height(28)
            .borderRadius(14)
            .backgroundColor(this.selectedAnswer === option.label ? '#3B82F6' : '#EFF6FF')
            .justifyContent(FlexAlign.Center)
            .margin({ right: 12 })

            Text(option.content)
              .fontSize(14)
              .fontColor('#374151')
              .layoutWeight(1)
              .lineHeight(20)
          }
          .width('100%')
          .padding({ left: 12, right: 12, top: 10, bottom: 10 })
          .backgroundColor(this.selectedAnswer === option.label ? '#EFF6FF' : '#FAFAFA')
          .borderRadius(8)
          .borderWidth(1)
          .borderColor(this.selectedAnswer === option.label ? '#3B82F6' : '#E2E8F0')
          .margin({ bottom: 8 })
          .onClick(() => {
            this.onAnswerSelected(this.question.id, option.label);
          })
        })
      }
    }
    .width('100%')
    .padding(16)
    .backgroundColor(Color.White)
    .borderRadius(12)
    .shadow({ radius: 4, color: 'rgba(0,0,0,0.06)', offsetX: 0, offsetY: 2 })
    .margin({ bottom: 12 })
  }
}
