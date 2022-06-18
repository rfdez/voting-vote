package vote_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/voting-vote/internal/platform/server/http/handler/vote"
	"github.com/rfdez/voting-vote/kit/command/commandmocks"
	"github.com/rfdez/voting-vote/kit/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestHandler_Create(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("creating.VoteCommand"),
	).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/polls/:poll_id/options/:option_id/votes", vote.CreateHandler(commandBus))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		b, err := json.Marshal(map[string]string{
			"user_id": "",
		})
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/polls/1/options/1/votes", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		require.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		b, err := json.Marshal(map[string]interface{}{
			"user_id": uuid.Generate(),
		})
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/polls/%s/options/%s/votes", uuid.Generate(), uuid.Generate()), bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
