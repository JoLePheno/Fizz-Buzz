package controller

import (
	"fmt"
	"log"

	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
)

type Fizzbuzz struct {
	Store port.Store
}

func (f *Fizzbuzz) Do(in *model.Parameters) ([]string, error) {
	if err := checkIfParametersAreCorrect(in); err != nil {
		return nil, err
	}

	var res []string
	for i := 1; i <= in.Limit; i++ {
		if i%in.FirstInteger == 0 && i%in.SecondInteger == 0 {
			res = append(res, in.FirstString+in.SecondString)
		} else if i%in.FirstInteger == 0 {
			res = append(res, in.FirstString)
		} else if i%in.SecondInteger == 0 {
			res = append(res, in.SecondString)
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}

	if err := f.Store.StoreParameters(in); err != nil {
		log.Default().Println("error when storing parameters: %w", err)
	}
	return res, nil
}

func checkIfParametersAreCorrect(in *model.Parameters) error {
	if in.Limit == 0 {
		return fmt.Errorf("invalid limit in parameters: %w", port.ErrInvalidLimit)
	}
	if in.FirstInteger == in.SecondInteger || in.FirstInteger > in.SecondInteger {
		return fmt.Errorf("invalid integers in parameters: %w", port.ErrInvalidIntegers)
	}
	if in.FirstInteger <= 1 {
		return fmt.Errorf("invalid integers in parameters: %w", port.ErrInvalidInteger)
	}
	return nil
}
