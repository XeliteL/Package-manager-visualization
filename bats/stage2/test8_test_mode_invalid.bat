@echo off
chcp 65001 >nul
echo === Тест 8: Ошибка — test_mode без файла ===
go run main.go -config tests/stage2/config_test_mode_invalid.json
echo -----------------------------------------------