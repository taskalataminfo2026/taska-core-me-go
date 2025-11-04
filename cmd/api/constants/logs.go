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
	FunctionCleanupExpiredTokens = "cleanup_expired_tokens"
	FunctionDeleteExpired        = "delete_expired"
	FunctionEmailCodeResend      = "email_code_resend"
	FunctionEmailCodeSend        = "email_code_send"
	FunctionEmailUpdate          = "email_update"
	FunctionEmailVerify          = "email_verify"
	FunctionFindAll              = "find_all"
	FunctionFirstBy              = "first_by"
	FunctionFirstByEmail         = "first_by_email"
	FunctionFirstByEmailNil      = "first_by_email_nil"
	FunctionFirstByToken         = "first_by_token"
	FunctionFirstByTokenNil      = "first_by_token_nil"
	FunctionFirstByUserName      = "first_by_user_name"
	FunctionFirstByUserNameNil   = "first_by_user_name_nil"
	FunctionGenerateAccessToken  = "generate_access_token"
	FunctionGenerateRefreshToken = "generate_refresh_token"
	FunctionLogin                = "login"
	FunctionLogout               = "logout"
	FunctionMakeEmailBody        = "make_email_body"
	FunctionPasswordChange       = "password_change"
	FunctionPasswordRecovery     = "password_recovery"
	FunctionPasswordReset        = "password_reset"
	FunctionRefreshToken         = "refresh_token"
	FunctionRegisterUser         = "register_user"
	FunctionSave                 = "save"
	FunctionSendEmail            = "send_email"
	FunctionUserProfile          = "user_profile"
	FunctionUserUpdate           = "user_update"
	FunctionValidateToken        = "validate_token"
)
