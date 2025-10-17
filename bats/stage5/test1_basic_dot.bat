@echo off
echo === Тест 1: Базовая визуализация графа (dot-режим) ===
go run main.go -package demo_pkg -file tests\stage5\deps_test.txt -test -mode dot
if exist graph.dot (
    echo Файл graph.dot успешно создан.
) else (
    echo Ошибка: файл graph.dot не найден!
)
echo -----------------------------------------------