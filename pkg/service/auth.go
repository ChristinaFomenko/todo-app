package service //реализация интерфейса
import (
	"crypto/sha1"
	"fmt"
	todo_app "github.com/KrisInferno/todo-app"
	"github.com/KrisInferno/todo-app/pkg/repository"
)

const salt = "fkbno43tn0df90hvm9u9gfdjv"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo_app.User) (int, error) {
	user.Password = generatePasswordHash(string(user.Password))

	return s.repo.CreateUser(user)
}

//хэширование пароля
func generatePasswordHash(password string) int {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
