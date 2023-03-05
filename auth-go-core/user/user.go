package user

type User struct {
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Avatar           string `json:"avatar" bson:"avatar"`
	OriginalResponse string `json:"original_response" bson:"original_response"`
}
