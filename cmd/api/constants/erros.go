package constants

const (
	ErrInvalidRequestFormat = "Formato de solicitud inválido: %v."
	ErrInvalidInputData     = "Datos de entrada inválidos: %v."
)

// Client DB.
const (
	ErrGettingDatabase       = "Error obteniendo la base de datos."
	ErrOpeningDBConnection   = "Error abriendo la conexión a la base de datos: %v."
	OpenDBConnectionsMessage = "Conexiones abiertas: %v."
)

// Skills repositories.
const (
	LogSkillsFound                 = "Skills encontrado: %+v"
	LogSkillsNotFoundByFilters     = "No se encontraron skills con los filtros aplicados"
	ErrSkillsQueryExecution        = "Error al ejecutar la consulta de skills"
	LogSkillUpsert                 = "Guardando o actualizando skill: %+v"
	ErrorMessageErrorFindingSkills = "Error al buscar por el id %v del skill."
	ErrSkillSave                   = "Error guardando el skill"
	ErrorMessageSkillsNotFound     = "No se encuentra ningún skill proporcionado."
)

// Users repositories.
const (
	ErrorMessageErrorFindingUser       = "Error al buscar por el id %v del usuario."
	ErrorMessageUserNotFoundByUserName = "No se encuentra ningún usuario con el nombre de usuario proporcionado."
)

// BlacklistedToken repositories.
const (
	ErrorMessageSavingToken = "Error al guardar o actualizar  el token en la lista negra: %s"
)
