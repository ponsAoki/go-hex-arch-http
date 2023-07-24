package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (httpa *Adapter) DtoHandler(fnc func(a, b int32) (int32, error)) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		queryMap := r.URL.Query()
		a, err := strconv.ParseInt(queryMap["a"][0], 10, 32)
		if err != nil {
			log.Fatalf("failed to convert a string to int: %v\n", err)
		}
		b, err := strconv.ParseInt(queryMap["b"][0], 10, 32)
		if err != nil {
			log.Fatalf("failed to convert b string to int: %v\n", err)
		}

		if a == 0 || b == 0 {
			log.Fatalf("missing required: %v\n", err)
		}

		answer, err := fnc(int32(a), int32(b))
		if err != nil {
			log.Fatalf("controller handling failed: %v", err)
		}

		json.NewEncoder(w).Encode(answer)
	}

	return handler
}

func (httpa *Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := httpa.api.GetAddition(a, b)
	if err != nil {
		log.Fatalf("internal server error: %v\n", err)
	}

	return answer, nil
}

func (httpa *Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := httpa.api.GetSubtraction(a, b)
	if err != nil {
		log.Fatalf("internal server error: %v\n", err)
	}

	return answer, nil
}

func (httpa *Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := httpa.api.GetMultiplication(a, b)
	if err != nil {
		log.Fatalf("internal server error: %v\n", err)
	}

	return answer, nil
}

func (httpa *Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := httpa.api.GetDivision(a, b)
	if err != nil {
		log.Fatalf("internal server error: %v\n", err)
	}

	return answer, nil
}
