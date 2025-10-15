@echo off
chcp 65001 >nul
echo === Тест 9: Вывод зависимостей в ASCII-виде ===
go run main.go -package demo_pkg -file tests/stage2/deps_test.txt -test -ascii
echo -----------------------------------------------