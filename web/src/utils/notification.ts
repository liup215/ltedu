// src/utils/notification.ts
import { push } from 'notivue'

/**
 * Show a global success notification using Notivue.
 * @param message Success message to display
 */
export function showSuccess(message: string) {
  push.success({
    title: 'Success',
    message: message,
    duration: 2000,
  })
}

/**
 * Show a global error notification using Notivue.
 * @param message Error message to display
 */
export function showError(message: string) {
  push.error({
    title: 'Error',
    message: message,
    duration: 2000, // Show for 5 seconds
  })

}
