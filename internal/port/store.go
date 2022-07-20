package port

import "github.com/JoLePheno/Fizz-Buzz/internal/model"

type Store interface {
	StoreParameters(params *model.Parameters) error
	RetrieveMostUsedRequestParameters() ([]*model.MostUsedParameters, error)
}
