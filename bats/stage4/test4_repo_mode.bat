@echo off
echo === Тест 4: Эмуляция из репозитория (order-режим) ===
go run main.go -package example_pkg -repo https://github.com/example/example.git -mode order
echo -----------------------------------------------
