import FormExtensionAbility from '@ohos.app.form.FormExtensionAbility';
import type formBindingData from '@ohos.app.form.formBindingData';
import type FormCallbackInfo from '@ohos.app.form.FormCallbackInfo';

/**
 * FormExtensionAbility provides an Atomic Service widget for quick access
 * to LTEdu features from the HarmonyOS home screen without full app install.
 */
export default class LTEduFormExtension extends FormExtensionAbility {
  onAddForm(want: object): formBindingData.FormBindingData {
    const formData: Record<string, string> = {
      title: 'LTEdu',
      subtitle: 'AI Learning Assistant',
      status: 'Ready to learn',
    };
    return formBindingData.createFormBindingData(formData);
  }

  onCastToNormalForm(formId: string): void {
    // Called when widget is cast from temporary to normal
  }

  onUpdateForm(formId: string): void {
    // Called when widget needs to update its data
    const formData: Record<string, string> = {
      title: 'LTEdu',
      subtitle: 'AI Learning Assistant',
      status: 'Tap to continue learning',
    };
    this.formProvider?.updateForm(formId, formBindingData.createFormBindingData(formData));
  }

  onFormEvent(formId: string, message: string): void {
    // Handle widget interaction events (e.g., open chat)
  }

  onRemoveForm(formId: string): void {
    // Widget removed by user
  }
}
