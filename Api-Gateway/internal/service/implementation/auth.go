package implementation

import "github.com/google/uuid"

type Auth struct{}

func (a *Auth) Login(email, password string) (token string, err error)
func (a *Auth) Register(email, password string) (uid uuid.UUID)
func (a *Auth) IsAdmin(uid uuid.UUID) (isAdmin bool, err error)
