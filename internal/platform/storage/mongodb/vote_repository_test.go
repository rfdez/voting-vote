package mongodb_test

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_OptionRepository_Save_RepositoryError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
}
