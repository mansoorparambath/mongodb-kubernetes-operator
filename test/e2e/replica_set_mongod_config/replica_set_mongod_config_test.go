package replica_set_mongod_config

import (
	"testing"

	e2eutil "github.com/mongodb/mongodb-kubernetes-operator/test/e2e"
	"github.com/mongodb/mongodb-kubernetes-operator/test/e2e/mongodbtests"
	setup "github.com/mongodb/mongodb-kubernetes-operator/test/e2e/setup"
	f "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/stretchr/objx"
)

func TestMain(m *testing.M) {
	f.MainEntry(m)
}

func TestReplicaSet(t *testing.T) {
	ctx, shouldCleanup := setup.InitTest(t)

	if shouldCleanup {
		defer ctx.Cleanup()
	}
	mdb, user := e2eutil.NewTestMongoDB("mdb0")

	_, err := setup.GeneratePasswordForUser(user, ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Override the journal compressor setting
	setting := "storage.wiredTiger.engineConfig.journalCompressor"
	value := "zlib"
	mongodConfig := objx.New(map[string]interface{}{})
	mongodConfig.Set(setting, value)
	mdb.Spec.AdditionalMongodConfig.Object = mongodConfig

	t.Run("Create MongoDB Resource", mongodbtests.CreateMongoDBResource(&mdb, ctx))
	t.Run("Basic tests", mongodbtests.BasicFunctionality(&mdb))
	t.Run("Test Basic Connectivity", mongodbtests.Connectivity(&mdb))
	t.Run("AutomationConfig has the correct version", mongodbtests.AutomationConfigVersionHasTheExpectedVersion(&mdb, 1))
	t.Run("Mongod config has been set", mongodbtests.EnsureMongodConfig(&mdb, setting, value))
}
