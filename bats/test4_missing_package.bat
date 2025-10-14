@echo off
echo === Тест 4: Отсутствует имя пакета ===
go run main.go -config tests/config_missing_package.json
echo ----------------------------------------