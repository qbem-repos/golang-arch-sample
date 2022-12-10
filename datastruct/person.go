package datastruct

type Person struct {
	ID    string `json:",omitempty" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
