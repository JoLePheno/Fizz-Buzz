package fizzbuzz

import (
	"fmt"

	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
)

type Fizzbuzz struct {
}

func (u *Fizzbuzz) Do(in *model.Parameters) ([]string, error) {
	var res []string

	if in.Limit == 0 {
		return nil, fmt.Errorf("invalid limit in parameters: %w", port.ErrInvalidLimit)
	}
	if in.FirstInteger == in.SecondInteger || in.FirstInteger > in.SecondInteger {
		return nil, fmt.Errorf("invalid integers in parameters: %w", port.ErrInvalidIntegers)
	}

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

	return res, nil
}
