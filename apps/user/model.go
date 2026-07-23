package user //serializer:json
import (
	"blog/common"

	"golang.org/x/crypto/bcrypt"
)

func NewCreateUserReuqest() *CreateUserReuqest {
	return &CreateUserReuqest{
		Role:  ROLE_FANGKE,
		Label: map[string]string{},
	}
}

type CreateUserReuqest struct {
	Username string            `json:"username" gorm:"column:username"`
	Password string            `json:"password" gorm:"column:password"`
	Role     Role              `json:"role" gorm:"column:role"`
	Label    map[string]string `json:"label" gorm:"column:label;serializer:json"`
}

func (c *CreateUserReuqest) HashPassword() error {
	ps, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Password = string(ps)
	return nil
}

func (c *CreateUserReuqest) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))
}

func NewUser(req *CreateUserReuqest) *User {
	return &User{
		Mate:              common.NewMate(),
		CreateUserReuqest: req,
	}
}

type User struct {
	*CreateUserReuqest
	*common.Mate
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

type UserSet struct {
	Total int64   `json:"total"`
	Items []*User `json:"items"`
}
