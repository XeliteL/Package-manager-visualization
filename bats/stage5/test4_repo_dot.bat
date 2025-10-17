@echo off
echo === Тест 4: Эмуляция зависимостей из репозитория (dot-режим) ===
go run main.go -package example_pkg -repo https://github.com/example/example.git -mode dot
if exist graph.dot (
    echo Файл graph.dot создан - OK.
) else (
    echo Ошибка: файл graph.dot не создан.
)
echo -----------------------------------------------