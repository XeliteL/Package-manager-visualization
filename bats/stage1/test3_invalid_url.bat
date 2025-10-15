@echo off
echo === Тест 3: Некорректный URL ===
go run main.go -config tests/stage1/config_invalid_url.json
echo ----------------------------------------