package methods

import (
	"context"

	"github.com/Artem09076/go_api.git/internal/db/sqlc"
	protos "github.com/Artem09076/go_api.git/protos/gen"
	"github.com/google/uuid"
)

type GroupServer struct {
	protos.UnimplementedGroupServiceServer
	queries *sqlc.Queries
}

func NewGroupServer(queries *sqlc.Queries) *GroupServer {
	return &GroupServer{queries: queries}
}

func (s *GroupServer) CreateGroup(ctx context.Context, req *protos.CreateGroupRequest) (*protos.Group, error) {
	group, err := s.queries.CreateGroup(ctx, sqlc.CreateGroupParams{
		Title:        req.GetTitle(),
		Descriptions: req.GetDescriptions(),
	})

	if err != nil {
		return nil, err
	}

	return &protos.Group{
		Title:        group.Title,
		Descriptions: group.Descriptions,
	}, nil
}

func (s *GroupServer) GetGroup(ctx context.Context, req *protos.GetGroupRequest) (*protos.Group, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	group, err := s.queries.GetGroup(ctx, id)
	if err != nil {
		return nil, err
	}
	return &protos.Group{
		Id:           group.ID.String(),
		Title:        group.Title,
		Descriptions: group.Descriptions,
	}, nil
}

func (s *GroupServer) UpdateGroup(ctx context.Context, req *protos.UpdateGroupRequest) (*protos.Group, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	group, err := s.queries.UpdateGroup(ctx, sqlc.UpdateGroupParams{
		ID:           id,
		Title:        req.GetTitle(),
		Descriptions: req.GetDescriptions(),
	})
	if err != nil {
		return nil, err
	}

	return &protos.Group{
		Id:           group.ID.String(),
		Title:        group.Title,
		Descriptions: group.Descriptions,
	}, nil

}

func (s *GroupServer) DeleteGroup(ctx context.Context, req *protos.DeleteGroupRequest) (*protos.EmptyGroup, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.queries.DeleteGroup(ctx, id); err != nil {
		return nil, err
	}
	return &protos.EmptyGroup{}, nil
}
