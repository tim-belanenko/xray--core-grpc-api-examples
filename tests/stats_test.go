package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	statsService "github.com/xtls/xray-core/app/stats/command"
)

func TestStats(t *testing.T) {
	ctx := context.Background()

	// before running this test, make sure you have xray-core running with the following command:
	// xray-core run -config ./config.json
	conn := prepareGrpcConn(t, ctx, "127.0.0.1:3000")
	stats := prepareStatsClient(t, conn)

	t.Run("GetSystemStats", func(t *testing.T) {
		resp, err := stats.GetSysStats(ctx, &statsService.SysStatsRequest{})
		require.NoError(t, err)
		assert.NotZero(t, resp.NumGoroutine)
		assert.NotZero(t, resp.NumGC)
		assert.NotZero(t, resp.Alloc)
		assert.NotZero(t, resp.TotalAlloc)
		assert.NotZero(t, resp.Sys)
		assert.NotZero(t, resp.Mallocs)
		assert.NotZero(t, resp.Frees)
		assert.NotZero(t, resp.LiveObjects)
		assert.NotZero(t, resp.PauseTotalNs)
		assert.NotZero(t, resp.Uptime)
	})

	// This test requires the following configuration in config.json:
	// config:
	// "policy": {
	//     "system": {
	//         "statsInboundDownlink": true,
	//         "statsInboundUplink": true
	//     }
	// },
	//
	// logs:
	//[Debug] app/stats: create new counter inbound>>>api>>>traffic>>>uplink
	//[Debug] app/stats: create new counter inbound>>>api>>>traffic>>>downlink
	t.Run("GetStats - existing counter", func(t *testing.T) {
		resp, err := stats.GetStats(ctx, &statsService.GetStatsRequest{
			Name: "inbound>>>api>>>traffic>>>uplink",
		})
		require.NoError(t, err)
		assert.Greater(t, resp.Stat.Value, int64(0))
	})

	t.Run("GetStats - non-existing counter", func(t *testing.T) {
		resp, err := stats.GetStats(ctx, &statsService.GetStatsRequest{
			Name: "non-existing-counter",
		})
		require.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("QueryStats - all existing counters", func(t *testing.T) {
		resp, err := stats.QueryStats(ctx, &statsService.QueryStatsRequest{
			Pattern: "",
		})
		require.NoError(t, err)
		assert.NotEmpty(t, resp.Stat)
	})

	t.Run("QueryStats - non-existing counters", func(t *testing.T) {
		resp, err := stats.QueryStats(ctx, &statsService.QueryStatsRequest{
			Pattern: "non-existing-counter",
		})
		require.NoError(t, err)
		assert.Empty(t, resp.Stat)
	})
}
