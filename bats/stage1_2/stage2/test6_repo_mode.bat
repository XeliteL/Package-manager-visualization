@echo off
chcp 65001 >nul
echo === Тест 6: Эмуляция получения зависимостей из репозитория ===
go run main.go -config tests/stage2/config_repo_mode.json
echo -----------------------------------------------