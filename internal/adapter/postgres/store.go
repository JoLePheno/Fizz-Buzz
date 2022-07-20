package postgres

import (
	"fmt"
	"time"

	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
	"github.com/go-pg/pg"
)

var _ port.Store = (*ParametersStore)(nil)

type ParametersStore struct {
	db *pg.DB
}

type Parameters struct {
	ID int64 `sql:"id"`

	CreatedAt time.Time `sql:"created_at,notnull"`
	DeletedAt time.Time `pg:"deleted_at,soft_delete"`

	FirstInteger  int `sql:"int1"`
	SecondInteger int `sql:"int2"`
	Limit         int `sql:"limit_number"`

	FirstString  string `sql:"str1"`
	SecondString string `sql:"str2"`
}

// StoreParameters implements port.Store
func (s *ParametersStore) StoreParameters(params *model.Parameters) error {
	if err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		currentTime := time.Now()
		return tx.Insert(&Parameters{
			CreatedAt:     currentTime,
			FirstInteger:  params.FirstInteger,
			SecondInteger: params.SecondInteger,
			Limit:         params.Limit,
			FirstString:   params.FirstString,
			SecondString:  params.SecondString,
		})
	}); err != nil {
		fmt.Println("error when storing params: ", err)
		return err
	}
	fmt.Println("parameters stored")
	return nil
}

// RetrieveParameters implements port.Store
func (s *ParametersStore) RetrieveParameters() ([]*model.Parameters, error) {
	var params []Parameters
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		return tx.Model(&params).Select()
	},
	)
	if err != nil {
		fmt.Println("error when fetching parameters: ", err)
		return nil, err
	}

	paramsModel := make([]*model.Parameters, len(params))
	for i := range params {
		paramsModel[i] = convertParamsModelDBOToModel(params[i])
	}
	return paramsModel, nil
}

// RetrieveParameter implements port.Store
func (s *ParametersStore) RetrieveParameter(param *model.Parameters) ([]*model.Parameters, error) {
	return nil, nil
}

func convertParamsModelDBOToModel(p Parameters) *model.Parameters {
	return &model.Parameters{
		FirstInteger:  p.FirstInteger,
		SecondInteger: p.SecondInteger,
		Limit:         p.Limit,
		FirstString:   p.FirstString,
		SecondString:  p.SecondString,
	}
}
