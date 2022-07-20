package dummy

import (
	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
)

var _ port.Store = (*DummyStore)(nil)

type DummyStore struct{}

// StoreParameters implements port.Store
func (*DummyStore) StoreParameters(params *model.Parameters) error {
	return nil
}

// RetrieveMostUsedRequestParameters implements port.Store
func (*DummyStore) RetrieveMostUsedRequestParameters() ([]*model.MostUsedParameters, error) {
	return nil, nil
}
