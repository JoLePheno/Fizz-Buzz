package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/JoLePheno/Fizz-Buzz/internal/controller"
	"github.com/JoLePheno/Fizz-Buzz/internal/model"
	"github.com/JoLePheno/Fizz-Buzz/internal/port"
	"github.com/JoLePheno/Fizz-Buzz/internal/utils"
	"github.com/gorilla/mux"
)

type AlgoService struct {
	FizzbuzzController *controller.Fizzbuzz
}

func (s *AlgoService) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.Methods("GET").Name("GetFizzBuzz").Handler(s.FizzBuzzHandler()).Path("/fizzbuzz")

	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	message := utils.Message(true, "Leboncoin looks cool")
	utils.Respond(w, message, 200)
}

func (a *AlgoService) FizzBuzzHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request /fizzbuzz, starting fizzbuzz")

		in := &model.Parameters{}
		err := json.NewDecoder(r.Body).Decode(in) //decode the request body into struct, failed if any error occured
		if err != nil {
			fmt.Println(errors.New("An error occurred while decoding request, err: " + err.Error()))
			utils.Respond(w, utils.Message(false, "Invalid number of parameters in request"), 400)
			return
		}

		fizzbuzzList, err := a.FizzbuzzController.Do(in)
		if err != nil {
			switch {
			case errors.Is(err, port.ErrInvalidIntegers), errors.Is(err, port.ErrInvalidLimit):
				utils.Respond(w, utils.Message(false, "Invalid request: "+err.Error()), 400)
			default:
				utils.Respond(w, utils.Message(false, "Internal Error"), 500)
			}
			return
		}
		utils.Respond(w, map[string]interface{}{
			"fizzBuzzList": fizzbuzzList,
		}, 200)
	})
}
