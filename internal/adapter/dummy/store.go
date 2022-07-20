package dummy

import (
	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
)

var _ port.Store = (*DummyStore)(nil)

type DummyStore struct{}

// RetrieveParameter implements port.Store
func (*DummyStore) RetrieveParameter(param *model.Parameters) ([]*model.Parameters, error) {
	return nil, nil
}

// RetrieveParameters implements port.Store
func (*DummyStore) RetrieveParameters() ([]*model.Parameters, error) {
	return nil, nil
}

// StoreParameters implements port.Store
func (*DummyStore) StoreParameters(params *model.Parameters) error {
	return nil
}
