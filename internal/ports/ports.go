package ports

type AuthService interface {
	Authorize()
	Validate()
	UserAuth()
	AdminAuth()
	GenerateToken()
	HashPassword()
	ComparePassword()
}

type UserService interface {
	SignUp()
	LogIn()
	UploadResume()
	GetAllJobs()
	GetJobByID()
}

type AdminService interface {
	CreateJob()
	GetJobInfo()
	GetAllUsers()
	GetUserProfile()
}

type UserRepository interface {
	CreateUser()
	GetUserByID()
	GetAllJobs()
	AddUserToJob()
	CreateProfile()
}

type AdminRepository interface {
	CreateJob()
	GetJobByID()
	GetAllUsers()
	GetProfileByUserID()
}
