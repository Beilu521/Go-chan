package database

import (
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) *gorm.DB{
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		slog.Error("failed to connect to database","err",err)
		os.Exit(1)
	}

	slog.Info("connected to database")
	return db
}

// slog 日志级别（从低到高，数值越大越严重）
// LevelDebug(-4)：调试信息，开发时用，默认不输出
// LevelInfo(0)：正常流程信息，默认输出（如启动成功、配置加载完成）
// LevelWarn(4)：警告，不影响主流程但要注意（如默认值兜底）
// LevelError(8)：错误，影响功能但不退出（如数据库连接失败、参数非法）