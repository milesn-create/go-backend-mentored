package service

import (
	"errors"
	"rest-api-gin/internal/models"
	"rest-api-gin/internal/repository"
	"testing"
)

func TestCreateUserSuccess(t *testing.T) {
	//arrange - подготовка
	repo := repository.NewInMemoryUserRepository()
	userService := NewUserService(repo)
	newUser := models.User{Name: "Alice", Age: 27}
	//act - действие
	createdUser, err := userService.CreateUser(newUser)
	//Assert - проверка
	if err != nil {
		t.Errorf("ожидалась успешная работа, но получлиась ошибка : %v", err)

	}
	if createdUser.Name != newUser.Name {
		t.Errorf("ожидалось имя созданного пользователя = %v, а создался пользователь с именем : %v", newUser.Name, createdUser.Name)
	}
	if createdUser.ID == 0 {
		t.Errorf("ожидалось что пользователю приссвоится реальный id, но id = %d", createdUser.ID)
	}

}
func TestCreateUserAlreadyExists(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	userService := NewUserService(repo)
	user1 := models.User{Name: "Alice", Age: 27}
	createUser1, _ := userService.CreateUser(user1)
	if createUser1.ID == 0 {
		t.Errorf("ожидалось что пользователю приссвоится реальный id, но id = %d", createUser1.ID)

	}
	user2 := models.User{Name: "Alice", Age: 27}
	createUser2, err := userService.CreateUser(user2)
	if err == nil {
		t.Errorf("ожидалось ошибка о том что пользователь с таким именим уже существует, но err = nil")
	}

	if !errors.Is(err, ErrNameAlreadyExists) {
		t.Errorf("ожидалась ошибка : %v, а получилась: %v", ErrNameAlreadyExists, err)
	}

	if createUser2.ID != createUser1.ID {
		t.Errorf("ожидалось что пользователь не создался ,и метод вернул данные пользоваткля уже существуего с таким именем id = %v, а в реальнсоти метод вернул id =%v", createUser1.ID, createUser2.ID)

	}
	if createUser2.Name != createUser1.Name || createUser2.Age != createUser1.Age {
		t.Errorf("ожидалось что пользователь не создался и метод вернул данные уже суествующего пользователя с таким именем : %v , %v , а в результате Name = %v, Age = %v", createUser1.Name, createUser1.Age, createUser2.Name, createUser2.Age)

	}

}
func TestFindByIdSuccess(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	userService := NewUserService(repo)
	user := models.User{Name: "Alice", Age: 27}
	createdUser1, _ := userService.CreateUser(user)
	FindUser, err := userService.FindByID(createdUser1.ID)
	if err != nil {
		t.Errorf("ожидалось успешное выполнение с возвратом данных о найденом пользователе , а получили ошибку:  %v", err)
	}
	if FindUser.ID != createdUser1.ID {
		t.Errorf("ожидалось что id найденог пользовтаеля равен искомому = %v, а оно равно = %v", createdUser1.ID, FindUser.ID)
	}
	if FindUser.Name != createdUser1.Name {
		t.Errorf("ожидалось что имя найденного пользователя = %v, а получили = %v", createdUser1.Name, FindUser.Name)
	}
	if FindUser.Age != createdUser1.Age {
		t.Errorf("ожидалось что возраст найденного пользователя = %v, а получили = %v", createdUser1.Age, FindUser.Age)
	}

}
func TestFindByIdNotExists(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	userService := NewUserService(repo)

	user_id := 5
	FindByIdUser, err := userService.FindByID(user_id)
	if err == nil {
		t.Errorf("ожидалось ошибка о том что пользователь с таким id не  существует, но err = nil")
	}
	if !errors.Is(err, repository.ErrUserNotFound) {
		t.Errorf("ожидалось ошибка о том что пользователь с таким id не  существует, но err = %v", err)

	}
	if FindByIdUser.ID != 0 || FindByIdUser.Age != 0 || FindByIdUser.Name != "" {
		t.Errorf("ожидалось пользователя с нулевыми значениями , но в результате ID  = %v , Name = %v , Age = %v", FindByIdUser.ID, FindByIdUser.Name, FindByIdUser.Age)

	}

}
