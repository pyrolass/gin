package entity

type User struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
