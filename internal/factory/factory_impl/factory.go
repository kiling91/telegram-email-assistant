package factory_impl

import (
	"github.com/kiling91/telegram-email-assistant/internal/email"
	"github.com/kiling91/telegram-email-assistant/internal/email/email_impl"
	"github.com/kiling91/telegram-email-assistant/internal/factory"
	"github.com/kiling91/telegram-email-assistant/internal/storage"
	"github.com/kiling91/telegram-email-assistant/internal/storage/storage_impl"
)

type service struct {
	store      storage.Storage
	imapServer email.ImapServer
}

func NewFactory() factory.Factory {
	return &service{}
}

func (s *service) GetStorage() storage.Storage {
	if s.store == nil {
		s.store = storage_impl.NewStorage(s)
	}
	return s.store
}

func (s *service) ImapServer() email.ImapServer {
	if s.imapServer == nil {
		s.imapServer = email_impl.NewImapServer(s)
	}
	return s.imapServer
}