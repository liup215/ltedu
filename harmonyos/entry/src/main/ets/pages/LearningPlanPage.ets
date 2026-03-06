import learningPlanService from '../services/LearningPlanService';
import type { LearningPlan } from '../models/ApiModels';
import { LoadingOverlay, ErrorBanner, EmptyState } from '../components/CommonComponents';

/**
 * LearningPlanPage shows the user's personal learning plans, allowing
 * them to view plan details and track their progress across plan types.
 */
@Component
export struct LearningPlanPage {
  @State plans: LearningPlan[] = [];
  @State isLoading: boolean = false;
  @State errorMessage: string = '';
  @State selectedPlan: LearningPlan | null = null;
  @State filterType: string = 'all';

  private readonly PLAN_TYPES: string[] = ['all', 'long', 'mid', 'short'];
  private readonly PLAN_TYPE_LABELS: Record<string, string> = {
    all: 'All',
    long: 'Long-term',
    mid: 'Mid-term',
    short: 'Short-term',
  };

  aboutToAppear(): void {
    this.loadPlans();
  }

  async loadPlans(): Promise<void> {
    this.isLoading = true;
    this.errorMessage = '';
    try {
      const req = this.filterType === 'all' ? {} : { planType: this.filterType };
      const response = await learningPlanService.listMyPlans({ ...req, pageSize: 50 });
      if (response.code === 0) {
        this.plans = response.data?.list || [];
      } else {
        this.errorMessage = response.msg || 'Failed to load learning plans.';
      }
    } catch {
      this.errorMessage = 'Network error. Please check your connection.';
    } finally {
      this.isLoading = false;
    }
  }

  private getPlanTypeColor(planType: string): string {
    switch (planType) {
      case 'long': return '#DBEAFE';
      case 'mid': return '#EDE9FE';
      case 'short': return '#D1FAE5';
      default: return '#F1F5F9';
    }
  }

  private getPlanTypeTextColor(planType: string): string {
    switch (planType) {
      case 'long': return '#1D4ED8';
      case 'mid': return '#5B21B6';
      case 'short': return '#065F46';
      default: return '#374151';
    }
  }

  private formatDate(dateStr: string): string {
    try {
      const d = new Date(dateStr);
      return `${d.getFullYear()}-${(d.getMonth() + 1).toString().padStart(2, '0')}-${d.getDate().toString().padStart(2, '0')}`;
    } catch {
      return dateStr;
    }
  }

  build() {
    Stack() {
      Column() {
        // Header
        Row() {
          Text('Learning Plans')
            .fontSize(20)
            .fontWeight(FontWeight.Bold)
            .fontColor('#1E293B')
          Blank()
          Button('↻')
            .fontSize(18)
            .fontColor('#64748B')
            .backgroundColor('#F1F5F9')
            .borderRadius(8)
            .height(36)
            .width(36)
            .onClick(() => { this.loadPlans(); })
        }
        .width('100%')
        .padding({ left: 20, right: 20, top: 16, bottom: 12 })
        .backgroundColor(Color.White)
        .shadow({ radius: 4, color: 'rgba(0,0,0,0.04)', offsetX: 0, offsetY: 2 })

        // Filter tabs
        Row() {
          ForEach(this.PLAN_TYPES, (type: string) => {
            Text(this.PLAN_TYPE_LABELS[type])
              .fontSize(13)
              .fontColor(this.filterType === type ? '#3B82F6' : '#64748B')
              .fontWeight(this.filterType === type ? FontWeight.SemiBold : FontWeight.Normal)
              .padding({ left: 16, right: 16, top: 8, bottom: 8 })
              .backgroundColor(this.filterType === type ? '#EFF6FF' : Color.Transparent)
              .borderRadius(20)
              .onClick(() => {
                this.filterType = type;
                this.loadPlans();
              })
          })
        }
        .width('100%')
        .padding({ left: 12, right: 12, top: 8, bottom: 8 })
        .backgroundColor(Color.White)

        ErrorBanner({ message: this.errorMessage, onDismiss: () => { this.errorMessage = ''; } })

        // Content
        if (this.selectedPlan !== null) {
          this.PlanDetailView()
        } else {
          this.PlanListView()
        }
      }
      .width('100%')
      .height('100%')
      .backgroundColor('#F8FAFC')

      LoadingOverlay({ visible: this.isLoading, message: 'Loading plans...' })
    }
    .width('100%')
    .height('100%')
  }

  @Builder
  PlanListView() {
    Scroll() {
      Column() {
        if (this.plans.length === 0 && !this.isLoading) {
          EmptyState({
            icon: '📅',
            title: 'No learning plans',
            subtitle: 'Your teacher will create learning plans for you. Check back soon!'
          })
        } else {
          ForEach(this.plans, (plan: LearningPlan) => {
            this.PlanCard(plan)
          }, (plan: LearningPlan) => plan.id.toString())
        }
      }
      .padding({ left: 16, right: 16, top: 16, bottom: 16 })
    }
    .layoutWeight(1)
    .scrollBar(BarState.Off)
  }

  @Builder
  PlanCard(plan: LearningPlan) {
    Column() {
      Row() {
        Text(plan.planType?.toUpperCase() || 'PLAN')
          .fontSize(11)
          .fontColor(this.getPlanTypeTextColor(plan.planType))
          .backgroundColor(this.getPlanTypeColor(plan.planType))
          .padding({ left: 8, right: 8, top: 3, bottom: 3 })
          .borderRadius(4)
          .fontWeight(FontWeight.SemiBold)

        Blank()

        Text(this.formatDate(plan.updatedAt))
          .fontSize(12)
          .fontColor('#94A3B8')
      }
      .width('100%')
      .margin({ bottom: 12 })

      Text(plan.content)
        .fontSize(14)
        .fontColor('#374151')
        .lineHeight(20)
        .maxLines(3)
        .textOverflow({ overflow: TextOverflow.Ellipsis })
        .width('100%')
        .margin({ bottom: 8 })

      if (plan.comment) {
        Text(plan.comment)
          .fontSize(12)
          .fontColor('#94A3B8')
          .fontStyle(FontStyle.Italic)
          .width('100%')
      }

      Row() {
        Blank()
        Text('View Details →')
          .fontSize(13)
          .fontColor('#3B82F6')
          .fontWeight(FontWeight.Medium)
          .margin({ top: 12 })
      }
    }
    .width('100%')
    .padding(16)
    .backgroundColor(Color.White)
    .borderRadius(16)
    .shadow({ radius: 4, color: 'rgba(0,0,0,0.06)', offsetX: 0, offsetY: 2 })
    .margin({ bottom: 12 })
    .onClick(() => {
      this.selectedPlan = plan;
    })
  }

  @Builder
  PlanDetailView() {
    Column() {
      // Back button
      Row() {
        Text('← Back')
          .fontSize(15)
          .fontColor('#3B82F6')
          .fontWeight(FontWeight.Medium)
          .margin({ bottom: 16 })
          .onClick(() => { this.selectedPlan = null; })
      }
      .width('100%')
      .padding({ left: 20, top: 12 })

      Scroll() {
        Column() {
          // Plan type badge
          Row() {
            Text((this.selectedPlan?.planType?.toUpperCase() || 'PLAN') + ' PLAN')
              .fontSize(12)
              .fontColor(this.getPlanTypeTextColor(this.selectedPlan?.planType || ''))
              .backgroundColor(this.getPlanTypeColor(this.selectedPlan?.planType || ''))
              .padding({ left: 10, right: 10, top: 4, bottom: 4 })
              .borderRadius(6)
              .fontWeight(FontWeight.SemiBold)
          }
          .width('100%')
          .margin({ bottom: 16 })

          // Plan content
          Column() {
            Text('Plan Content')
              .fontSize(14)
              .fontWeight(FontWeight.SemiBold)
              .fontColor('#374151')
              .alignSelf(ItemAlign.Start)
              .margin({ bottom: 12 })

            Text(this.selectedPlan?.content || '')
              .fontSize(14)
              .fontColor('#1E293B')
              .lineHeight(22)
              .width('100%')
          }
          .width('100%')
          .padding(16)
          .backgroundColor(Color.White)
          .borderRadius(16)
          .margin({ bottom: 16 })

          if (this.selectedPlan?.comment) {
            Column() {
              Text('Notes')
                .fontSize(14)
                .fontWeight(FontWeight.SemiBold)
                .fontColor('#374151')
                .alignSelf(ItemAlign.Start)
                .margin({ bottom: 12 })
              Text(this.selectedPlan?.comment || '')
                .fontSize(14)
                .fontColor('#64748B')
                .lineHeight(20)
                .fontStyle(FontStyle.Italic)
                .width('100%')
            }
            .width('100%')
            .padding(16)
            .backgroundColor(Color.White)
            .borderRadius(16)
            .margin({ bottom: 16 })
          }

          // Dates
          Row() {
            Column() {
              Text('Created')
                .fontSize(12)
                .fontColor('#94A3B8')
                .margin({ bottom: 4 })
              Text(this.formatDate(this.selectedPlan?.createdAt || ''))
                .fontSize(14)
                .fontColor('#374151')
                .fontWeight(FontWeight.Medium)
            }
            .alignItems(HorizontalAlign.Center)
            .layoutWeight(1)

            Divider().vertical(true).height(32).color('#E2E8F0')

            Column() {
              Text('Updated')
                .fontSize(12)
                .fontColor('#94A3B8')
                .margin({ bottom: 4 })
              Text(this.formatDate(this.selectedPlan?.updatedAt || ''))
                .fontSize(14)
                .fontColor('#374151')
                .fontWeight(FontWeight.Medium)
            }
            .alignItems(HorizontalAlign.Center)
            .layoutWeight(1)
          }
          .width('100%')
          .padding(16)
          .backgroundColor(Color.White)
          .borderRadius(16)
        }
        .padding({ left: 16, right: 16, bottom: 24 })
      }
      .layoutWeight(1)
      .scrollBar(BarState.Off)
    }
    .layoutWeight(1)
  }
}
