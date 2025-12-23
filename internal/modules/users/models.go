package users

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

/* Combining Repository and service files to simplify process
// The Repository Interface (Port)
type Repository interface {
	FindByID(id int) (*User, error)
}
*/

// // The Service Interface (Port). This should be in service file
// type Service interface {
// 	GetUser(id int) (User, error)
// }
