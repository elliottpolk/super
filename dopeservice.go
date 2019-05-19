package super

import (
	"context"

	"github.com/elliottpolk/super/config"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DopeServer struct {
	cmp    *config.Composition
	client *mongo.Client
}

func NewDopeServer(cmp *config.Composition, client *mongo.Client) DopeServiceServer {
	return &DopeServer{
		cmp:    cmp,
		client: client,
	}
}

func (s *DopeServer) Create(ctx context.Context, req *DopeRequest) (*Empty, error) {
	empty := &Empty{RequestId: req.RequestId}

	if s.client == nil {
		return empty, ErrNoMongoClient
	}

	client := s.client
	if err := client.UseSession(ctx, func(session mongo.SessionContext) error {
		defer session.EndSession(ctx)

		if err := Create(session, req.Payload, client.Database(repo)); err != nil {
			defer session.AbortTransaction(ctx)
			return err
		}

		return nil
	}); err != nil {
		return empty, err
	}
	return empty, nil
}

func (s *DopeServer) Retrieve(ctx context.Context, req *DopeRequest) (*DopeResponse, error) {
	if s.client == nil {
		return nil, ErrNoMongoClient
	}

	result := &DopeResponse{
		RequestId: req.RequestId,
	}

	client := s.client
	if err := client.UseSession(ctx, func(session mongo.SessionContext) error {
		defer session.EndSession(ctx)

		// retrieve 1 and return by ID if provided in request
		if id := req.Id; len(id) > 0 {
			item, err := RetrieveOne(ctx, id, client.Database(repo))
			if err != nil {
				return errors.Wrapf(err, "unable to retrieve record for id %s", id)
			}
			result.Payload = []*Dope{item}

			return nil
		}

		items, err := Retrieve(ctx, bson.D{}, client.Database(repo))
		if err != nil {
			return errors.Wrap(err, "unable to retrieve records")
		}
		result.Payload = items

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *DopeServer) Update(ctx context.Context, req *DopeRequest) (*DopeResponse, error) {
	if s.client == nil {
		return nil, ErrNoMongoClient
	}

	result := &DopeResponse{
		RequestId: req.RequestId,
	}

	client := s.client
	if err := client.UseSession(ctx, func(session mongo.SessionContext) error {
		defer session.EndSession(ctx)

		if err := Update(session, req.Username, bson.D{}, req.Payload, client.Database(repo)); err != nil {
			defer session.AbortTransaction(ctx)
			return errors.Wrap(err, "unable to update records")
		}
		result.Payload = req.Payload

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *DopeServer) Delete(ctx context.Context, req *DopeRequest) (*Empty, error) {
	empty := &Empty{RequestId: req.RequestId}

	if s.client == nil {
		return empty, ErrNoMongoClient
	}

	client := s.client
	if err := client.UseSession(ctx, func(session mongo.SessionContext) error {
		defer session.EndSession(ctx)

		if err := Delete(session, req.Username, req.Payload, client.Database(repo)); err != nil {
			defer session.AbortTransaction(ctx)
			return errors.Wrap(err, "unable to delete records")
		}

		return nil
	}); err != nil {
		return empty, err
	}

	return empty, nil
}
