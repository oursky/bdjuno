package remote

import (
	"fmt"

	inflationtypes "github.com/e-money/em-ledger/x/inflation/types"
	"github.com/forbole/juno/v2/node/remote"

	inflationsource "github.com/forbole/bdjuno/v2/modules/inflation/source"
)

var (
	_ inflationsource.Source = &Source{}
)

// Source implements inflationsource.Source using a remote node
type Source struct {
	*remote.Source
	client inflationtypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, client inflationtypes.QueryClient) *Source {
	return &Source{
		Source: source,
		client: client,
	}
}

// GetInflation implements inflationsource.Source
func (s *Source) GetInflation(height int64) (inflationtypes.InflationState, error) {
	res, err := s.client.Inflation(remote.GetHeightRequestContext(s.Ctx, height), &inflationtypes.QueryInflationRequest{})
	if err != nil {
		return inflationtypes.InflationState{}, fmt.Errorf("error while querying inflation state: %s", err)
	}

	return res.State, nil
}