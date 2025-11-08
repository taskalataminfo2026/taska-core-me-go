package constants

// Client DB.
const (
	ErrGettingDatabase       = "Error obteniendo la base de datos."
	ErrOpeningDBConnection   = "Error abriendo la conexión a la base de datos: %v."
	OpenDBConnectionsMessage = "Conexiones abiertas: %v."
)

// Core services - Error messages
const (
	ErrGeneratingAccessToken = "Error al generar tokens de acceso."
)

// Password hash
const (
	ErrProcessingPassword = "Error al procesar la contraseña."
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

// Auth services - Error messages
const (
	ErrPermissionDenied  = "No tienes permisos suficientes para acceder a este recurso."
	ErrRoleHierarchyLoad = "Error al cargar jerarquía de roles."
	ErrTokenRevoked      = "El token ha sido revocado. Por favor inicia sesión nuevamente."
)
