package models

type User struct {
	Id	string `gorethink:"id,omitempty" json:"id"`
	Email	string `json:"email" gorethink:"email"`
	Name	string `json:"name" gorethink:"name"`
	Salt string `json:"password" gorethink:"password"`
}

type CreateUser struct {
	Email	string `json:"email" binding:"required"`
	Name	string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email	string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangeUser struct {
	Id	int32 `json:"id" binding:"required"`
	Email	string `json:"email" binding:"required"`
	Name	string `json:"name" binding:"required"`
}
