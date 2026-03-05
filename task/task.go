package task

import (
	"edu/service"
	"time"
)

func RunTask() {
	go service.DocumentSvr.LoopConvertDocument()
	go runConversationCleanup()
}

// runConversationCleanup periodically removes expired conversation sessions.
func runConversationCleanup() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	for range ticker.C {
		_, _ = service.ConversationSvr.CleanupExpiredSessions()
	}
}

