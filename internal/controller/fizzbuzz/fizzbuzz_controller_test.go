package fizzbuzz

import (
	"testing"

	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
	"github.com/stretchr/testify/require"
)

func TestFizzBuzzController(t *testing.T) {
	t.Run("Simple Test", func(t *testing.T) {
		in := &model.Parameters{
			FirstInteger:  3,
			SecondInteger: 5,
			Limit:         30,
			FirstString:   "fizz",
			SecondString:  "buzz",
		}

		f := Fizzbuzz{}
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

		f := Fizzbuzz{}
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

		f := Fizzbuzz{}
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

		f := Fizzbuzz{}
		res, err := f.Do(in)
		require.Error(t, err, port.ErrInvalidIntegers)
		require.Nil(t, res)
	})
	// t.Run("Big Test", func(t *testing.T) {
	// 	in := &model.Parameters{
	// 		FirstInteger:  3,
	// 		SecondInteger: 5,
	// 		Limit:         math.MaxInt - 1,
	// 		FirstString:   "fizz",
	// 		SecondString:  "buzz",
	// 	}

	// 	f := Fizzbuzz{}
	// 	res, err := f.Do(in)
	// 	require.NoError(t, err)
	// 	require.Len(t, res, in.Limit)
	// })
}
