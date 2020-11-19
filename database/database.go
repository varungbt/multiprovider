package database

import (
	"context"
	"multiprovider/database/driver"
	muxutil "multiprovider/internal/muxutil"
)

// DbClient ...
type DbClient struct {
	Client driver.DbClient
}

// DbClientConnOpener ...
type DbClientConnOpener interface {
	OpenConnection(ctx context.Context) (*DbClient, error)
}

// NewDbClient ...
func NewDbClient(drv driver.DbClient) *DbClient {
	return &DbClient{
		Client: drv,
	}
}

//---------------------------------------------------- Provider Mux ---------------------------------------------------------

// ProviderMux ...
type ProviderMux struct {
	providers muxutil.ProviderMap
}

var defaultProviderMux = new(ProviderMux)

// DefaultProviderMux ...
func DefaultProviderMux() *ProviderMux {
	return defaultProviderMux
}

// RegisterDbClient ...
func (mux *ProviderMux) RegisterDbClient(provider string, dbClientOpener DbClientConnOpener) {
	mux.providers.Register(provider, dbClientOpener)
}

// OpenConnection ...
func (mux *ProviderMux) OpenConnection(ctx context.Context, provider string) (*DbClient, error) {
	opener := mux.providers.FromString(provider)
	dbClientOpener := opener.(DbClientConnOpener)
	return dbClientOpener.OpenConnection(ctx)
}

//----------------------------------------------------------------------------------------------------------------------------

// OpenConnection ...
func OpenConnection(ctx context.Context, provider string) (*DbClient, error) {
	conn, _ := defaultProviderMux.OpenConnection(ctx, provider)
	return conn, nil
}
