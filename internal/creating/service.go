package creating

import (
	"context"

	voting "github.com/rfdez/voting-vote/internal"
)

// Service is the interface that must be implemented by the creating service.
type Service interface {
	CreateVote(ctx context.Context, pollID, optionID, userID string) error
}

type service struct {
	voteRepository voting.VoteRepository
}

// NewService creates a new creating service.
func NewService(voteRepository voting.VoteRepository) Service {
	return &service{
		voteRepository: voteRepository,
	}
}

func (s *service) CreateVote(ctx context.Context, pollID, optionID, userID string) error {
	vote, err := voting.NewVote(pollID, optionID, userID)
	if err != nil {
		return err
	}

	return s.voteRepository.Save(ctx, vote)
}
