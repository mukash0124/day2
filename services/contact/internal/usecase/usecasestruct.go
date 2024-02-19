package usecase

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
)

type contactUseCase struct {
	repo repository.ContactRepository
}

// NewContactUseCase создает новый экземпляр ContactUseCase
func NewContactUseCase(repo repository.ContactRepository) ContactUseCase {
	return &contactUseCase{
		repo: repo,
	}
}

func (uc *contactUseCase) CreateContact(contact domain.Contact) error {
	// method call
	return uc.repo.CreateContact(contact)
}

func (uc *contactUseCase) ReadContact(contactID int) (domain.Contact, error) {
	// method call
	return uc.repo.ReadContact(contactID)
}

func (uc *contactUseCase) UpdateContact(contactID int, newContact domain.Contact) error {
	// method call
	return uc.repo.UpdateContact(contactID, newContact)
}

func (uc *contactUseCase) DeleteContact(contactID int) error {
	// method call
	return uc.repo.DeleteContact(contactID)
}

func (uc *contactUseCase) CreateGroup(group domain.Group) error {
	// method call
	return uc.repo.CreateGroup(group)
}

func (uc *contactUseCase) ReadGroup(groupID int) (domain.Group, error) {
	// method call
	return uc.repo.ReadGroup(groupID)
}

func (uc *contactUseCase) AddContactToGroup(contactID, groupID int) error {
	// method call
	return uc.repo.AddContactToGroup(contactID, groupID)
}
