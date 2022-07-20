package controller

import (
	"testing"

	"github.com/JoLePheno/Fizz-Buzz/internal/adapter/dummy"
	"github.com/JoLePheno/Fizz-Buzz/internal/adapter/postgres"
	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
	"github.com/stretchr/testify/require"
)

func TestFizzBuzzController(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	t.Run("Simple Test", func(t *testing.T) {
		in := &model.Parameters{
			FirstInteger:  3,
			SecondInteger: 5,
			Limit:         30,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}

		f := Fizzbuzz{
			Store: &dummy.DummyStore{},
		}
		res, err := f.Do(in)
		require.NoError(t, err)
		require.Len(t, res, in.Limit)
	})
	t.Run("Invalid Limit", func(t *testing.T) {
		in := &model.Parameters{
			FirstInteger:  3,
			SecondInteger: 5,
			Limit:         0,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}

		f := Fizzbuzz{
			Store: &dummy.DummyStore{},
		}
		res, err := f.Do(in)
		require.Error(t, err, port.ErrInvalidLimit)
		require.Nil(t, res)
	})
	t.Run("Error With Same Integers", func(t *testing.T) {
		in := &model.Parameters{
			FirstInteger:  5,
			SecondInteger: 5,
			Limit:         0,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}

		f := Fizzbuzz{
			Store: &dummy.DummyStore{},
		}
		res, err := f.Do(in)
		require.Error(t, err, port.ErrInvalidIntegers)
		require.Nil(t, res)
	})
	t.Run("Error Invalid Integers int1 > int2", func(t *testing.T) {
		in := &model.Parameters{
			FirstInteger:  5,
			SecondInteger: 2,
			Limit:         0,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}

		f := Fizzbuzz{
			Store: &dummy.DummyStore{},
		}
		res, err := f.Do(in)
		require.Error(t, err, port.ErrInvalidIntegers)
		require.Nil(t, res)
	})
}

func TestIntegrationFizzBuzzController(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	store := postgres.NewPostgresStore()

	t.Run("Simple Test", func(t *testing.T) {
		store.ClearStore()
		in := &model.Parameters{
			FirstInteger:  3,
			SecondInteger: 5,
			Limit:         30,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}

		f := Fizzbuzz{
			Store: store,
		}
		res, err := f.Do(in)
		require.NoError(t, err)
		require.Len(t, res, in.Limit)

		expectedStoredReq := []*model.MostUsedParameters{
			{
				TotalOccurence: 1,
				FirstInteger:   in.FirstInteger,
				SecondInteger:  in.SecondInteger,
				Limit:          in.Limit,
				FirstString:    in.FirstString,
				SecondString:   in.SecondString,
			},
		}
		storedReq, err := f.Store.RetrieveMostUsedRequestParameters()
		require.NoError(t, err)
		require.Len(t, storedReq, 1)
		require.EqualValues(t, expectedStoredReq, storedReq)
	})
	t.Run("Multiple values stored Test", func(t *testing.T) {
		store.ClearStore()
		in := &model.Parameters{
			FirstInteger:  3,
			SecondInteger: 5,
			Limit:         30,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}
		in2 := &model.Parameters{
			FirstInteger:  3,
			SecondInteger: 5,
			Limit:         30,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}
		in3 := &model.Parameters{
			FirstInteger:  2,
			SecondInteger: 6,
			Limit:         30,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}

		f := Fizzbuzz{
			Store: store,
		}
		_, err := f.Do(in)
		require.NoError(t, err)

		_, err = f.Do(in2)
		require.NoError(t, err)

		_, err = f.Do(in3)
		require.NoError(t, err)

		expectedStoredReq := []*model.MostUsedParameters{
			{
				TotalOccurence: 2,
				FirstInteger:   in.FirstInteger,
				SecondInteger:  in.SecondInteger,
				Limit:          in.Limit,
				FirstString:    in.FirstString,
				SecondString:   in.SecondString,
			},
			{
				TotalOccurence: 1,
				FirstInteger:   in3.FirstInteger,
				SecondInteger:  in3.SecondInteger,
				Limit:          in3.Limit,
				FirstString:    in3.FirstString,
				SecondString:   in3.SecondString,
			},
		}

		storedReq, err := f.Store.RetrieveMostUsedRequestParameters()
		require.NoError(t, err)
		require.Len(t, storedReq, 2)
		require.EqualValues(t, expectedStoredReq, storedReq)
	})
}
