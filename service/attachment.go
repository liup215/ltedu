package service

import (
	"edu/model"
	"edu/repository"
)

var AttachmentSvr *AttachmentService = &AttachmentService{
	baseService: newBaseService(),
}

type AttachmentService struct {
	baseService
}

func (s *AttachmentService) CreateAttachment(a *model.Attachment) error {
	if a.ID != 0 {
		a.ID = 0
	}

	return repository.AttachmentRepo.Create(a)
}

func (s *AttachmentService) SelectAttachmentById(id uint) (*model.Attachment, error) {
	return repository.AttachmentRepo.FindByID(id)
}
