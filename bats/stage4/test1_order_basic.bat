@echo off
echo === Тест 1: Порядок загрузки зависимостей (основной) ===
go run main.go -package demo_pkg -file tests\stage4\deps_dup.txt -test -mode order
echo -----------------------------------------------