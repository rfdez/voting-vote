package creating_test

import (
	"context"
	"testing"

	"github.com/rfdez/voting-vote/internal/creating"
	"github.com/rfdez/voting-vote/internal/errors"
	"github.com/rfdez/voting-vote/internal/platform/storage/storagemocks"
	"github.com/rfdez/voting-vote/kit/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Service_CreateVote_RepositoryError(t *testing.T) {
	voteRepositoryMock := new(storagemocks.VoteRepository)
	voteRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Vote")).Return(errors.New("error"))

	creatingService := creating.NewService(voteRepositoryMock)

	err := creatingService.CreateVote(context.Background(), uuid.Generate(), uuid.Generate(), uuid.Generate())

	voteRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_CreateVote_Succeed(t *testing.T) {
	voteRepositoryMock := new(storagemocks.VoteRepository)
	voteRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Vote")).Return(nil)

	creatingService := creating.NewService(voteRepositoryMock)

	err := creatingService.CreateVote(context.Background(), uuid.Generate(), uuid.Generate(), uuid.Generate())

	voteRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
