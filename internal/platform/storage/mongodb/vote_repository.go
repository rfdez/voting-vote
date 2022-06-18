package mongodb

import (
	"context"
	"time"

	voting "github.com/rfdez/voting-vote/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "votes"
)

type voteRepository struct {
	collection *mongo.Collection
	dbTimeout  time.Duration
}

func NewVoteRepository(db *mongo.Database, dbTimeout time.Duration) voting.VoteRepository {
	return &voteRepository{
		collection: db.Collection(collection),
		dbTimeout:  dbTimeout,
	}
}

func (r *voteRepository) Save(ctx context.Context, vote voting.Vote) error {
	filter := bson.D{
		primitive.E{
			Key: "_id", Value: bson.D{
				primitive.E{
					Key:   "poll_id",
					Value: vote.PollID().String(),
				},
				primitive.E{
					Key:   "user_id",
					Value: vote.UserID().String(),
				},
			},
		},
	}
	update := bson.D{
		primitive.E{
			Key: "$set", Value: bson.D{
				primitive.E{
					Key: "options",
					Value: bson.A{
						vote.OptionID().String(),
					},
				},
			},
		},
	}
	opts := options.Update().SetUpsert(true)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.collection.UpdateOne(ctxTimeout, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
