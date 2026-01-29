package task

import "edu/service"

func RunTask() {
	go service.DocumentSvr.LoopConvertDocument()
}
