
| Rama        | Propósito                                          | Quién puede modificarla          | Tipo de merge permitido                           |
| ----------- | -------------------------------------------------- | -------------------------------- | ------------------------------------------------- |
| `master`    | Versión estable y en producción.                   | Solo el Tech Lead o CI/CD.       | **Merge squash** desde `release/*`.               |
| `develop`   | Integración de nuevas funcionalidades validadas.   | Todo el equipo (PR obligatorio). | **Merge squash** desde `feature/*`.               |
| `feature/*` | Desarrollo de nuevas funcionalidades o mejoras.    | Cada desarrollador.              | **Merge → develop** cuando esté lista y revisada. |
| `release/*` | Preparación de versiones estables para despliegue. | Tech Lead / QA.                  | **Merge → master** y luego **→ develop**.         |
| `hotfix/*`  | Corrección urgente en producción.                  | Tech Lead o responsable del bug. | **Merge → master** y luego **→ develop**.         |

