package opensearchdata

import (
	"context"
	"fmt"
	"sync"

	"log/slog"

	"github.com/open-panoptes/opni/pkg/plugins/apis/system"
	"github.com/open-panoptes/opni/pkg/util/future"
	loggingutil "github.com/open-panoptes/opni/plugins/logging/pkg/util"
)

const (
	pendingValue     = "job pending"
	opensearchPrefix = "os_"
)

type DeleteStatus int

const (
	DeletePending DeleteStatus = iota
	DeleteRunning
	DeleteFinished
	DeleteFinishedWithErrors
	DeleteError
)

type ClusterStatus int

// Ready func should return true if there is a critical error
// That would stop the opensearch query from running.
type ReadyFunc func() bool

const (
	ClusterStatusGreen = iota
	ClusterStatusYellow
	ClusterStatusRed
	ClusterStatusError
	ClusterStatusNoClient
)

type Manager struct {
	*loggingutil.AsyncOpensearchClient

	systemKV future.Future[system.KeyValueStoreClient]
	logger   *slog.Logger

	adminInitStateRW sync.RWMutex
}

func NewManager(logger *slog.Logger, kv future.Future[system.KeyValueStoreClient]) *Manager {
	return &Manager{
		AsyncOpensearchClient: loggingutil.NewAsyncOpensearchClient(),
		systemKV:              kv,
		logger:                logger,
	}
}

func (m *Manager) keyExists(keyToCheck string) (bool, error) {
	prefixKey := &system.ListKeysRequest{
		Key: opensearchPrefix,
	}
	keys, err := m.systemKV.Get().ListKeys(context.Background(), prefixKey)
	if err != nil {
		return false, err
	}
	for _, key := range keys.GetKeys() {
		if key == fmt.Sprintf("%s%s", opensearchPrefix, keyToCheck) {
			return true, nil
		}
	}
	return false, nil
}
