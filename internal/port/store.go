package port

import "github.com/JoLePheno/Fizz-Buzz/internal/model"

type Store interface {
	StoreParameters(params *model.Parameters) error
	RetrieveParameters() ([]*model.Parameters, error)
	RetrieveParameter(param *model.Parameters) ([]*model.Parameters, error)
}
