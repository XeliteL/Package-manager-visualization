@echo off
echo === Тест 4: Отсутствует имя пакета ===
go run main.go -config tests/stage1/config_missing_package.json
echo ----------------------------------------