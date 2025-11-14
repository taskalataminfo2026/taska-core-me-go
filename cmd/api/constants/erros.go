package constants

// Core services - Error message.
const (
	ErrGeneratingAccessToken = "Error al generar tokens de acceso."
)

// Password hash.
const (
	ErrProcessingPassword = "Error al procesar la contraseña."
)

// Client DB.
const (
	ErrGettingDatabase       = "Error obteniendo la base de datos."
	ErrOpeningDBConnection   = "Error abriendo la conexión a la base de datos: %v."
	OpenDBConnectionsMessage = "Conexiones abiertas: %v."
)

// Client JWT.
const (
	ErrMsgAccessTokenGenFailed    = "No se pudo generar el token de acceso."
	ErrMsgRefreshTokenGenFailed   = "No se pudo generar el token de actualización."
	ErrMsgUnexpectedSigningMethod = "Método de firma inesperado"
	InvalidCredentialsMessage     = "Credenciales incorrectas."
	ErrMsgInvalidRefreshToken     = "Token de actualización no válido."
	SessionExpiredMessage         = "Tu sesión ha terminado, por favor inicia sesión nuevamente."
	ErrTokenMissing               = "Token vacío o inválido."
	ErrTokenInvalid               = "Token inválido o corrupto."
	ErrTokenExpired               = "El token ha expirado, solicita uno nuevo."
	ErrDatabaseNotInitialized     = "Conexión a BD no inicializada."
	ErrTokenBadFormat             = "Formato de token inválido."
	ErrTokenEmpty                 = "Token vacío o inválido."
	ErrTokenSignatureInvalid      = "Firma de token inválida."
	ErrTokenBlacklistValidation   = "Error validando token en lista negra"
)

// Auth services - Error messages.
const (
	ErrPermissionDenied  = "No tienes permisos suficientes para acceder a este recurso."
	ErrRoleHierarchyLoad = "Error al cargar jerarquía de roles."
	ErrTokenRevoked      = "El token ha sido revocado. Por favor inicia sesión nuevamente."
)

// Users repositories.
const (
	ErrorMessageErrorFindingUser       = "Error al buscar el usuario: %v."
	ErrorMessageUserNotFoundByUserName = "No se encuentra ningún usuario con el nombre de usuario proporcionado."
)

// BlacklistedToken repositories.
const (
	ErrorMessageTokenNotFound     = "El token no se encuentra en la lista negra."
	ErrorMessageErrorFindingToken = "Error al buscar el token en la lista negra."
	ErrorMessageSavingToken       = "Error al guardar o actualizar  el token en la lista negra: %s"
)

// Users models.
const (
	MsgInvalidEmail         = "Correo electrónico inválido."
	MsgPasswordLowercase    = "La contraseña debe contener al menos una letra minúscula."
	MsgPasswordMinLength    = "La contraseña debe tener al menos 5 caracteres."
	MsgPasswordNumber       = "La contraseña debe contener al menos un número."
	MsgPasswordSpecialChar  = "La contraseña debe contener al menos un carácter especial."
	MsgPasswordUppercase    = "La contraseña debe contener al menos una letra mayúscula."
	MsgUsernameInvalidChars = "El nombre de usuario solo puede contener letras, números, guiones bajos y puntos."
	MsgUsernameMaxLength    = "El nombre de usuario no puede tener más de 75 caracteres."
	MsgUsernameMinLength    = "El nombre de usuario debe tener al menos 3 caracteres."
)
