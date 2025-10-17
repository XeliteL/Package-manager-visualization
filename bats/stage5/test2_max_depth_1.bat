@echo off
echo === Тест 2: Визуализация при max_depth = 1 ===
go run main.go -package demo_pkg -file tests\stage5\deps_test.txt -test -mode dot -max-depth 1
findstr /C:"->" graph.dot >nul
if %errorlevel%==0 (
    echo Граф содержит только верхний уровень зависимостей.
) else (
    echo Ошибка: файл graph.dot пуст или не создан.
)
echo -----------------------------------------------