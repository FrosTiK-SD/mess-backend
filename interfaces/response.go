package interfaces

type GenericResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
	Error   interface{} `json:"error"`
}

func GetGenericResponse(success bool, message string, details interface{}, err error) GenericResponse {
	if err != nil {
		return GenericResponse{
			Success: success,
			Message: message,
			Details: details,
			Error:   err.Error(),
		}
	} else {
		return GenericResponse{
			Success: success,
			Message: message,
			Details: details,
		}
	}
}
