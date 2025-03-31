package config

import "os"

var (
	Mode          string
	ServerPort    string
	RedisPort     string
	RedisServer   string
	UploadPath    string
	ServerBaseUrl string
	FullBaseUrl   string
	RedisBaseUrl  string
	ImageBasePath string
	DbHost        string
	DbPort        string
	DbName        string
	DbUsername    string
	DbPassword    string
)

// กำหนดค่าต่างๆที่ต้องใช้ลงในตัวแปรเพื่อนำไปใช้ต่อ
func Init() {
	Mode = getEnv("GIN_MODE")
	ServerPort = getEnv("SERVER_PORT")
	RedisPort = getEnv("REDIS_PORT")
	RedisServer = getEnv("REDIS_SERVER")
	UploadPath = getEnv("UPLOAD_PATH")
	ServerBaseUrl = getEnv("BASE_URL")
	FullBaseUrl = ServerBaseUrl + ServerPort
	RedisBaseUrl = RedisServer + RedisPort
	ImageBasePath = getEnv("IMAGE_BASE_PATH")
	DbHost = getEnv("DB_HOST")
	DbPort = getEnv("DB_PORT")
	DbName = getEnv("DB_NAME")
	DbUsername = getEnv("DB_USER")
	DbPassword = getEnv("DB_PASSWORD")
}

// รับค่าจาก env ไฟล์
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return ""
	}
	return value
}
