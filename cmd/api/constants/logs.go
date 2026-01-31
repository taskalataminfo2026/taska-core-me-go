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
	ModuleCategories       = "categories "
	ModuleSkills           = "skills"
	ModuleRoles            = "roles"
	ModuleBlacklistedToken = "black_listed_token"
	ModuleJwt              = "jwt"
)

// Function Skills
const (
	FunctionSkillsSearch = "skills_search"
	FunctionSkillsList   = "skills_list"
	FunctionSkillsSave   = "skills_save"
	FunctionSkillsUpdate = "skills_update"
)

// Function Categories
const (
	FunctionCategorySearch = "Category_search"
	FunctionCategoryList   = "Category_list"
	FunctionCategorySave   = "Category_save"
	FunctionCategoryUpdate = "skills_update"
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
