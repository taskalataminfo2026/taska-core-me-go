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
	ModuleSkills           = "skills"
	ModuleRoles            = "roles"
	ModuleBlacklistedToken = "black_listed_token"
	ModuleJwt              = "jwt"
)

// Function Skills
const (
	FunctionSkillsList = "skills_list"
)

// Function Roles
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
	FunctionUpsert               = "upsert"
)
