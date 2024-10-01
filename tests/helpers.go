package tests

import (
	"context"
	"testing"

	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

func prepareGrpcConn(t *testing.T, ctx context.Context, address string) *grpc.ClientConn {
	t.Helper()
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func prepareStatsClient(t *testing.T, conn *grpc.ClientConn) statsService.StatsServiceClient {
	t.Helper()
	return statsService.NewStatsServiceClient(conn)
}
