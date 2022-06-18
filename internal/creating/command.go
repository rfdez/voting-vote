package creating

import (
	"context"

	"github.com/rfdez/voting-vote/internal/errors"
	"github.com/rfdez/voting-vote/kit/command"
)

const (
	VoteCommandType command.Type = "voting-app.voting-vote.1.command.vote.create"
)

// VoteCommand is the command to create a vote
type VoteCommand struct {
	pollID   string
	optionID string
	userID   string
}

// NewVoteCommand creates a new VoteCommand
func NewVoteCommand(pollID, optionID, userID string) VoteCommand {
	return VoteCommand{
		pollID:   pollID,
		optionID: optionID,
		userID:   userID,
	}
}

func (c VoteCommand) Type() command.Type {
	return VoteCommandType
}

// VoteCommandHandler is the handler for VoteCommand
type VoteCommandHandler struct {
	service Service
}

// NewVoteCommandHandler creates a new VoteCommandHandler
func NewVoteCommandHandler(service Service) VoteCommandHandler {
	return VoteCommandHandler{
		service: service,
	}
}

func (h VoteCommandHandler) Handle(ctx context.Context, c command.Command) error {
	cmd, ok := c.(VoteCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateVote(ctx, cmd.pollID, cmd.optionID, cmd.userID)
}
