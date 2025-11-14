package constants

// Layer
const (
	LayerController = "controller"
	LayerGateway    = "gateway"
	LayerRepository = "repository"
	LayerService    = "service"
)

// Module
const (
	ModuleAuth             = "auth"
	ModuleBlacklistedToken = "blacklisted_token"
	ModuleEmail            = "email"
	ModuleEmailBrevo       = "email_brevo"
	ModuleJwt              = "jwt"
	ModulePassword         = "password"
	ModuleRoles            = "roles"
	ModuleUser             = "user"
)

// Function
const (
	FunctionFindAll              = "find_all"
	FunctionFirstBy              = "first_by"
	FunctionGenerateAccessToken  = "generate_access_token"
	FunctionGenerateRefreshToken = "generate_refresh_token"
	FunctionRefreshToken         = "refresh_token"
	FunctionValidateToken        = "validate_token"
	FunctionDeleteExpired        = "delete_expired"
	FunctionFirstByToken         = "first_by_token"
	FunctionFirstByTokenNil      = "first_by_token_nil"
	FunctionSave                 = "save"
)
