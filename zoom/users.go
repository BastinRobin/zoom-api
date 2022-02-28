package zoom

import (
	"encoding/json"
	"log"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PMI       int64  `json:"pmi"`
	Verified  int    `json:"verified"`
}
type Users []User

type UserResponse struct {
	TotalRecords int `json:"total_records"`
	Users        Users
}

// Return the list of all users in zoom account
func (z *Zoom) GetUsers() (UserResponse, error) {
	users, err := z.Request("/users", "GET")
	if err != nil {
		log.Println(err)
		return UserResponse{}, err
	}

	var response UserResponse

	// Unmarshal the response into Response struct
	err = json.Unmarshal(users, &response)
	if err != nil {
		log.Println(err)
		return UserResponse{}, err
	}

	return response, nil

}
