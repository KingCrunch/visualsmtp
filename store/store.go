package store

import (
	"github.com/satori/go.uuid"
	"github.com/KingCrunch/visualsmtp/mail"
)

type Store interface {
	Get(id uuid.UUID) (mail.Mail, error)
	List() (map[uuid.UUID]mail.Mail, error)
	Push(mail mail.Mail) (uuid.UUID, error)
}