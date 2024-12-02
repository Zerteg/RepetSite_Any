package models

// Error - структура для описания ошибок API
// @Description Error object
// @Param error body models.Error true "Error details"
// @Success 200 {object} models.Error
// @Failure 400 {object} models.Error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
