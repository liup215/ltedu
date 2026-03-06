import router from '@ohos.router';
import { ChatPage } from './ChatPage';
import { PracticePage } from './PracticePage';
import { LearningPlanPage } from './LearningPlanPage';
import { ProfilePage } from './ProfilePage';

/**
 * MainPage is the root navigation container after authentication.
 * It uses a TabBar with four tabs:
 *   - AI Chat (primary, index 0)
 *   - Practice (index 1)
 *   - Learning Plans (index 2)
 *   - Profile (index 3)
 */
@Entry
@Component
struct MainPage {
  @State currentTab: number = 0;

  build() {
    Tabs({ barPosition: BarPosition.End, index: this.currentTab }) {
      TabContent() {
        ChatPage()
      }
      .tabBar(this.TabItem('💬', 'AI Chat', 0))

      TabContent() {
        PracticePage()
      }
      .tabBar(this.TabItem('📝', 'Practice', 1))

      TabContent() {
        LearningPlanPage()
      }
      .tabBar(this.TabItem('📅', 'Plans', 2))

      TabContent() {
        ProfilePage()
      }
      .tabBar(this.TabItem('👤', 'Profile', 3))
    }
    .width('100%')
    .height('100%')
    .barMode(BarMode.Fixed)
    .barWidth('100%')
    .barHeight(60)
    .animationDuration(200)
    .onChange((index: number) => {
      this.currentTab = index;
    })
  }

  @Builder
  TabItem(icon: string, label: string, index: number) {
    Column() {
      Text(icon)
        .fontSize(this.currentTab === index ? 24 : 22)
        .margin({ bottom: 2 })
      Text(label)
        .fontSize(11)
        .fontColor(this.currentTab === index ? '#3B82F6' : '#94A3B8')
        .fontWeight(this.currentTab === index ? FontWeight.Medium : FontWeight.Normal)
    }
    .width('100%')
    .height('100%')
    .justifyContent(FlexAlign.Center)
    .alignItems(HorizontalAlign.Center)
    .backgroundColor(Color.White)
  }
}
