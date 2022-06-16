package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sharedInterfaces/auth"
	"time"
)

type server struct {
	router *http.ServeMux
	routes auth.Routes
}

func newServer() (*server, error) {
	router := http.NewServeMux()

	s := &server{
		router: router,
		routes: auth.NewRoutes(),
	}
	s.setupRoutes()
	return s, nil
}

func (s *server) printInfo(r *http.Request) {
	fmt.Printf("%v | %v | %v\n", time.Now(), r.URL.Path, r.Method)
}

func (s *server) getJSONFromRequest(r *http.Request, v interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	requestBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(requestBodyBytes, &v)
	return err
}

func (s *server) writeResponse(w http.ResponseWriter, jsonData interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(jsonData)
	if err != nil {
		http.Error(w, "Error: Failed to encode and write response JSON data", http.StatusInternalServerError)
	}
}

func (s *server) writeErrorResponse(w http.ResponseWriter, respErr error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	type errorResponse struct {
		Error      string `json:"error"`
		StatusCode int    `json:"status_code"`
	}
	err := json.NewEncoder(w).Encode(errorResponse{
		Error:      respErr.Error(),
		StatusCode: code,
	})
	if err != nil {
		http.Error(w, "Error: failed to encode and write response data", http.StatusInternalServerError)
	}
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	n, err := w.Write([]byte("hello is it 15?"))
	if err != nil {
		fmt.Printf("Error writing to response\n")
		return
	}
	fmt.Printf("n = %v\n", n)
}

func (s *server) handleNewAuthProfile(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hi"))
	//handleErr := func(err error) {
	//
	//}
	//
	//makeRespData := func(uuid string, isEmailValid bool, isPasswordValid bool) auth.NewAuthProfileResponse {
	//	return auth.NewAuthProfileResponse{
	//		UUID:            uuid,
	//		IsEmailValid:    isEmailValid,
	//		IsPasswordValid: isPasswordValid,
	//	}
	//}

	//return func(w http.ResponseWriter, r *http.Request) {
	s.printInfo(r)

	// Get email & password from request
	var reqData auth.NewAuthProfileRequest
	err := s.getJSONFromRequest(r, &reqData)
	if err != nil {
		s.writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	isEmailValid := isEmailValid(reqData.Email)
	isPasswordValid := isPasswordValid(reqData.Password)

	if !isEmailValid || !isPasswordValid {
		s.writeResponse(w, auth.NewAuthProfileResponse{IsEmailValid: isEmailValid, IsPasswordValid: isPasswordValid},
			http.StatusBadRequest)
		return
	}

	// new auth profile - i.e. new UUID for email & password
	//uuid := newUUID()

	// Hash password
	//passwordHash := hashPassword(reqData.Password)

	// save new auth profile to database

	// return UUID
	//}

	respData := auth.NewAuthProfileResponse{
		UUID:            "some UUID",
		IsEmailValid:    true,
		IsPasswordValid: true,
	}
	s.writeResponse(w, respData, http.StatusOK)
}

func run() error {
	s, err := newServer()
	if err != nil {
		return err
	}

	httpServer := http.Server{
		Addr:    ":8000",
		Handler: s.router,
	}
	err = httpServer.ListenAndServe()
	return err
}

func main() {
	fmt.Printf("TwitterClone-service: Auth\n")

	rand.Seed(time.Now().UTC().UnixNano())

	if err := run(); err != nil {
		fmt.Printf("Error: auth server failed to start %v", err.Error())
	}
}
