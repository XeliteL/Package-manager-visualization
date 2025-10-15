@echo off
chcp 65001 >nul
echo === Тест 7: Тестовый режим с файлом зависимостей ===
go run main.go -config tests/stage2/config_test_mode_valid.json
echo -----------------------------------------------