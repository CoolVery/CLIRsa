package models

//Структура User из БД
//json - как называть поле при сериализации
//db - как называется поле в БД для сопоставления в запросах
//binding:"required" - то, что строка не пустая, число не 0 и т.п., есть другие теги
type User struct {
	UserId int 			`json:"user_id" db:"user_id"` 
	RsaNum string 		`json:"rsa_num" binding:"required" db:"rsa_num"` 
	RsaKey string 		`json:"rsa_key" binding:"required" db:"rsa_key"`
	UserNameFull string `json:"user_name_full" binding:"required" db:"user_name_full"`
	UserLogin string 	`json:"user_login" binding:"required" db:"user_login"`
}