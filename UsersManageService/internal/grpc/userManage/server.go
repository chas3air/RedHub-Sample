package usermanage

import (
	"context"
	"usersManageService/internal/domain/models"
	"usersManageService/internal/domain/profiles"

	umv1 "github.com/chas3air/protos/gen/go/usersManager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

type UsersManager interface {
	ListUsers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, uid uuid.UUID) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	Insert(ctx context.Context, user models.User) error
	Update(ctx context.Context, uid uuid.UUID, user models.User) error
	Delete(ctx context.Context, uid uuid.UUID) (models.User, error)
}

type serverAPI struct {
	umv1.UnimplementedUsersManagerServer
	userManager UsersManager
}

func Register(grpc *grpc.Server, userManager UsersManager) {
	umv1.RegisterUsersManagerServer(grpc, &serverAPI{userManager: userManager})
}

func (s *serverAPI) ListUsers(ctx context.Context, req *umv1.ListUsersRequest) (*umv1.ListUsersResponse, error) {
	app_users, err := s.userManager.ListUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve app users")
	}

	resp_users := make([]*umv1.User, 0, len(app_users))
	for _, user := range app_users {
		resp_users = append(resp_users, profiles.UsrToProroUsr(user))
	}

	return &umv1.ListUsersResponse{
		Users: resp_users,
	}, nil
}

func (s *serverAPI) GetUserById(ctx context.Context, req *umv1.GetUserByIdRequest) (*umv1.GetUserByIdResponse, error) {
	parsedUUID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}
	requested_user, err := s.userManager.GetUserById(ctx, parsedUUID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user no found")
	}

	return &umv1.GetUserByIdResponse{
		User: profiles.UsrToProroUsr(requested_user),
	}, nil
}

func (s *serverAPI) GetUserByLogin(ctx context.Context, req *umv1.GetUserByEmailRequest) (*umv1.GetUserByEmailResponse, error) {
	requested_user, err := s.userManager.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	return &umv1.GetUserByEmailResponse{
		User: profiles.UsrToProroUsr(requested_user),
	}, nil
}

func (s *serverAPI) InsertUser(ctx context.Context, req *umv1.InsertUserRequest) (*umv1.InsertUserResponse, error) {
	parsedUser := profiles.ProtoUsrToUsr((req.GetUser()))

	err := s.userManager.Insert(ctx, parsedUser)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	return nil, nil
}

func (s *serverAPI) UpdateUser(ctx context.Context, req *umv1.UpdateUserRequest) (*umv1.UpdateUserResponse, error) {
	parsedUser := profiles.ProtoUsrToUsr(req.GetUser())
	parsedUUID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	err = s.userManager.Update(ctx, parsedUUID, parsedUser)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	return nil, nil
}

func (s *serverAPI) Delete(ctx context.Context, req *umv1.DeleteRequest) (*umv1.DeleteResponse, error) {
	parsedUUID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	user, err := s.userManager.Delete(ctx, parsedUUID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid arguments")
	}

	return &umv1.DeleteResponse{
		User: profiles.UsrToProroUsr(user),
	}, nil
}
