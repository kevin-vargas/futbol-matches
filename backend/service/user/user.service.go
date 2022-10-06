package user

import (
	"backend/model"
	ur "backend/repository/user"
	"backend/service"
)

func (us *UserService) Create(user model.User) (string, error) {
	encryptedPass, err := us.encrypter.Generate(user.Password)

	if err == nil {
		user.Password = encryptedPass
		_, createdErr := us.userRepo.Create(user)

		if createdErr != nil {
			return "", createdErr
		} else {
			return us.jwt.Generate(user.Username)
		}
	} else {
		return "", err
	}
}

func (us *UserService) GetAll() []model.User {
	return us.userRepo.GetAll()
}

func (us *UserService) GetByUsername(username string) model.User {
	return us.userRepo.GetByUsername(username)
}

func (us *UserService) Update(username string, user model.User) error {
	encryptedPass, err := us.encrypter.Generate(user.Password)

	if err == nil {
		user.Password = encryptedPass
		return us.userRepo.Update(username, user)
	}
	return err
}

func (us *UserService) Delete(username string) error {
	return us.userRepo.Delete(username)
}

func (us *UserService) Login(username string, password string) (string, error) {
	user := us.userRepo.GetByUsername(username)

	if user.Password != "" {
		compare, err := us.encrypter.Compare(user.Password, password)
		if err != nil {
			return "", err
		}
		if compare {
			return us.jwt.Generate(user.Username)
		}
	}
	return "", nil
}

type UserService struct {
	userRepo  ur.UserRepository
	encrypter service.Encrypt
	jwt       service.JWT
}

func NewUserService(userRepo ur.UserRepository, encryptService service.Encrypt, jwt service.JWT) UserService {
	return UserService{
		userRepo:  userRepo,
		encrypter: encryptService,
		jwt:       jwt,
	}
}
