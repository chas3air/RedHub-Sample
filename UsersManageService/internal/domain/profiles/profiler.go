package profiles

import (
	"usersManageService/internal/domain/models"

	umv1 "github.com/chas3air/protos/gen/go/usersManager"
	"github.com/google/uuid"
)

func UsrToProroUsr(user models.User) *umv1.User {
	return &umv1.User{
		Id:       user.Id.String(),
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Nick:     user.Nick,
	}
}

func ProtoUsrToUsr(proto_usr *umv1.User) models.User {
	parsedUUID, err := uuid.Parse(proto_usr.GetId())
	if err != nil {
		return models.User{}
	}

	return models.User{
		Id:       parsedUUID,
		Email:    proto_usr.GetEmail(),
		Password: proto_usr.GetPassword(),
		Role:     proto_usr.GetRole(),
		Nick:     proto_usr.GetNick(),
	}
}
