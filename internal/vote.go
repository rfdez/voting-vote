package voting

import (
	"context"

	"github.com/rfdez/voting-vote/kit/uuid"
)

// PollID represents a poll identifier.
type PollID struct {
	value string
}

// NewPollID instantiate the VO for PollID.
func NewPollID(value string) (PollID, error) {
	v, err := uuid.New(value)
	if err != nil {
		return PollID{}, err
	}

	return PollID{
		value: v,
	}, nil
}

// String returns the string representation of the PollID.
func (id PollID) String() string {
	return id.value
}

// OptionID represents a option identifier.
type OptionID struct {
	value string
}

// NewOptionID instantiate the VO for OptionID.
func NewOptionID(value string) (OptionID, error) {
	v, err := uuid.New(value)
	if err != nil {
		return OptionID{}, err
	}

	return OptionID{
		value: v,
	}, nil
}

// String returns the string representation of the OptionID.
func (id OptionID) String() string {
	return id.value
}

// UserID represents an user identifier.
type UserID struct {
	value string
}

// NewUserID instantiate the VO for UserID.
func NewUserID(value string) (UserID, error) {
	v, err := uuid.New(value)
	if err != nil {
		return UserID{}, err
	}

	return UserID{
		value: v,
	}, nil
}

// String returns the string representation of the UserID.
func (id UserID) String() string {
	return id.value
}

// Vote is the data structure that represents a vote.
type Vote struct {
	pollID   PollID
	optionID OptionID
	userID   UserID
}

// VoteRepository is the interface that must be implemented by the vote repository.
type VoteRepository interface {
	Save(context.Context, Vote) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=VoteRepository

// NewVote creates a new Vote.
func NewVote(pollID, optionID, userID string) (Vote, error) {
	pollIDVO, err := NewPollID(pollID)
	if err != nil {
		return Vote{}, err
	}

	optionIDVO, err := NewOptionID(optionID)
	if err != nil {
		return Vote{}, err
	}

	userIDVO, err := NewUserID(userID)
	if err != nil {
		return Vote{}, err
	}

	vote := Vote{
		pollID:   pollIDVO,
		optionID: optionIDVO,
		userID:   userIDVO,
	}

	return vote, nil
}

// PollID returns the poll identifier.
func (v Vote) PollID() PollID {
	return v.pollID
}

// OptionID returns the option identifier.
func (v Vote) OptionID() OptionID {
	return v.optionID
}

// UserID returns the user identifier.
func (v Vote) UserID() UserID {
	return v.userID
}
