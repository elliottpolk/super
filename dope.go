package super

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const repo string = "dope"

// Create will validate fields for the provided records and attempt to create
// new Dope records in the datastore
func Create(ctx context.Context, items []*Dope, db *mongo.Database) error {
	// swap to []interface{} because mongo needs it
	in := make([]interface{}, len(items))

	// validate / enrich required fields
	for i, item := range items {
		// the provider must specify at least the CreatedBy value
		if item.RecordInfo == nil {
			return ErrInvalidRecordInfo
		}

		// verify created_by is populated for an attempt at an audit
		if len(item.RecordInfo.CreatedBy) < 1 {
			return ErrInvalidCreatedBy
		}

		// ensure the Dope has an unique identifier
		if len(item.Id) < 1 {
			item.Id = uuid.New().String()
		}

		// ensure the created value is populated
		if item.RecordInfo.Created == nil || item.RecordInfo.Created.Seconds < 1 {
			item.RecordInfo.Created = &timestamp.Timestamp{Seconds: time.Now().Unix()}
		}

		in[i] = item
	}

	// write Dope to datastore
	if _, err := db.Collection(repo).InsertMany(ctx, in); err != nil {
		return err
	}

	// return the written element
	return nil
}

// RetrieveOne ...
func RetrieveOne(ctx context.Context, id string, db *mongo.Database) (*Dope, error) {
	res, err := Retrieve(ctx, bson.D{{"_id", id}}, db)
	if err != nil {
		return nil, err
	}

	if len(res) < 1 {
		return nil, ErrNotFound
	}

	if len(res) > 1 {
		return nil, ErrMutlipleRecordsReturned
	}

	return res[0], nil
}

// Retrieve ...
func Retrieve(ctx context.Context, filter bson.D, db *mongo.Database) ([]*Dope, error) {
	iter, err := db.Collection(repo).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer iter.Close(ctx)

	items := make([]*Dope, 0)
	for iter.Next(ctx) {
		item := &Dope{}
		if err := iter.Decode(&item); err != nil {
			return nil, errors.Wrapf(err, "unable to decode record")
		}
		items = append(items, item)
	}

	return items, nil
}

// Update ...
func Update(ctx context.Context, user string, filter bson.D, items []*Dope, db *mongo.Database) error {
	// ensure the user provided a username in an attempt to audit
	if len(user) < 1 {
		return ErrInvalidUsername
	}

	log.WithFields(log.Fields{
		"user":        user,
		"action_type": "update",
	}).Infof("attempting to update %d records", len(items))

	count := int64(0)
	for _, item := range items {
		item.RecordInfo.Updated = &timestamp.Timestamp{Seconds: time.Now().Unix()}

		res, err := db.Collection(repo).ReplaceOne(ctx, filter, item)
		if res != nil {
			count += res.ModifiedCount
		}

		if err != nil {
			return errors.Wrapf(err, "update: expected %d - actually %d", len(items), count)
		}
	}

	if want, got := int64(len(items)), count; got < want {
		return errors.Wrapf(ErrIncompleteAction, "update: expected %d - actually %d", want, got)
	}

	log.WithFields(log.Fields{
		"user":        user,
		"action_type": "update",
	}).Infof("updated %d records", len(items))

	return nil
}

// Delete ...
func Delete(ctx context.Context, user string, items []*Dope, db *mongo.Database) error {
	// ensure the user provided a username in an attempt to audit
	if len(user) < 1 {
		return ErrInvalidUsername
	}

	ids := make([]string, len(items))
	for i, item := range items {
		if len(item.Id) < 1 {
			return ErrInvalidId
		}
		ids[i] = item.Id
	}

	log.WithFields(log.Fields{
		"user":        user,
		"action_type": "delete",
	}).Infof("attempting to delete %d records", len(items))

	res, err := db.Collection(repo).DeleteMany(ctx, bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		return errors.Wrapf(err, "deletion: expected %d - actually %d", len(ids), res.DeletedCount)
	}

	if want, got := int64(len(ids)), res.DeletedCount; got < want {
		return errors.Wrapf(ErrIncompleteAction, "deletion: expected %d - actually %d", want, got)
	}

	log.WithFields(log.Fields{
		"user":        user,
		"action_type": "delete",
	}).Infof("deleted %d records", res.DeletedCount)

	return nil
}
