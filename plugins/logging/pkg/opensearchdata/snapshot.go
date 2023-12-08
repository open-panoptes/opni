package opensearchdata

import (
	"context"
	"fmt"
	"strings"
	"time"

	opensearchtypes "github.com/open-panoptes/opni/pkg/opensearch/opensearch/types"
	loggingerrors "github.com/open-panoptes/opni/plugins/logging/pkg/errors"
	"github.com/opensearch-project/opensearch-go/opensearchutil"
)

func (m *Manager) DoSnapshot(ctx context.Context, repository string, indices []string) error {
	m.WaitForInit()

	snapshotName := fmt.Sprintf("upgrade-%s", time.Now().Format(time.UnixDate))

	settings := opensearchtypes.SnapshotRequest{
		Indices: strings.Join(indices, ","),
	}

	resp, err := m.Client.Snapshot.CreateSnapshot(ctx, snapshotName, repository, opensearchutil.NewJSONReader(settings), false)
	if err != nil {
		return loggingerrors.WrappedOpensearchFailure(err)
	}

	defer resp.Body.Close()

	if resp.IsError() {
		m.logger.Error(fmt.Sprintf("opensearch request failed: %s", resp.String()))
		return loggingerrors.ErrOpensearchResponse
	}

	return nil
}
