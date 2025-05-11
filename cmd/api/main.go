package main

import (
	// 基本的な書式処理とI/O機能を提供（文字列フォーマットなど）
	"log" // 標準ログ機能を提供
	// HTTPクライアントとサーバー機能を提供
	"github.com/joho/godotenv"    // 環境変数管理
	"github.com/labstack/echo/v4" // Webフレームワーク
	"go.uber.org/zap"             // ロギング
)

func main() {
	// 環境変数の読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// ロガーの設定
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Echoフレームワークの初期化
	e := echo.New()

	// ヘルスチェックエンドポイントの実装
	e.GET("/health", func(c echo.C))

}
