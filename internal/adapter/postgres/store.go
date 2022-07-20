package postgres

import (
	"errors"
	"log"
	"time"

	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
	"github.com/go-pg/pg"
)

var _ port.Store = (*ParametersStore)(nil)

type ParametersStore struct {
	db *pg.DB
}

type paramsDBO struct {
	Total int   `sql:"total"`
	ID    int64 `sql:"id"`

	FirstInteger  int `sql:"int1"`
	SecondInteger int `sql:"int2"`
	Limit         int `sql:"limit_number"`

	FirstString  string `sql:"str1"`
	SecondString string `sql:"str2"`
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
		log.Default().Println("error when storing params: ", err)
		return err
	}
	log.Default().Println("parameters stored")
	return nil
}

// RetrieveMostUsedRequestParameters implements port.Store
// Find the top 3 most made request
func (s *ParametersStore) RetrieveMostUsedRequestParameters() ([]*model.MostUsedParameters, error) {
	var params []paramsDBO
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.Model().
			Query(&params, `SELECT int1, int2, str1, str2, limit_number, count(*) as total from parameters group by int1, int2, str1, str2, limit_number ORDER BY total DESC;`)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) { // NO result found we return
			return nil, nil
		} else {
			log.Default().Println("error when fetching parameters: ", err)
			return nil, err
		}
	}

	resp := make([]*model.MostUsedParameters, len(params))
	for i := range params {
		resp[i] = convertParamsModelDBOToModel(params[i])
	}
	return resp, nil
}

func convertParamsModelDBOToModel(p paramsDBO) *model.MostUsedParameters {
	return &model.MostUsedParameters{
		TotalOccurence: p.Total,
		FirstInteger:   p.FirstInteger,
		SecondInteger:  p.SecondInteger,
		Limit:          p.Limit,
		FirstString:    p.FirstString,
		SecondString:   p.SecondString,
	}
}
