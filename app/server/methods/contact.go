package methods

import (
	"context"

	"github.com/Artem09076/go_api.git/internal/db/sqlc"
	protos "github.com/Artem09076/go_api.git/protos/gen"
	"github.com/google/uuid"
)

type ContactServer struct {
	protos.UnimplementedContactServiceServer
	queries *sqlc.Queries
}

func NewContactServer(queries *sqlc.Queries) *ContactServer {
	return &ContactServer{queries: queries}
}

func (s *ContactServer) CreateContact(ctx context.Context, req *protos.CreateContactRequest) (*protos.Contact, error) {

	contact, err := s.queries.CreateContact(ctx, sqlc.CreateContactParams{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}

	return &protos.Contact{
		Username: contact.Username,
		Email:    contact.Email,
	}, nil
}

func (s *ContactServer) GetContact(ctx context.Context, req *protos.GetContactRequest) (*protos.Contact, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	contact, err := s.queries.GetContact(ctx, id)
	if err != nil {
		return nil, err
	}

	return &protos.Contact{
		Id:       contact.ID.String(),
		Username: contact.Username,
		Email:    contact.Email,
	}, nil
}

func (s *ContactServer) UpdateContact(ctx context.Context, req *protos.UpdateContactRequest) (*protos.Contact, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	contact, err := s.queries.UpdateContact(ctx, sqlc.UpdateContactParams{
		ID:       id,
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}

	return &protos.Contact{
		Id:       contact.ID.String(),
		Username: contact.Username,
		Email:    contact.Email,
	}, nil

}

func (s *ContactServer) DeleteContact(ctx context.Context, req *protos.DeleteContactRequest) (*protos.EmptyContact, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	if err := s.queries.DeleteContact(ctx, id); err != nil {
		return nil, err
	}
	return &protos.EmptyContact{}, nil
}
