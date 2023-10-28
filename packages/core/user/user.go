package user

type User struct {
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Avatar           string `json:"avatar" bson:"avatar"`
}
