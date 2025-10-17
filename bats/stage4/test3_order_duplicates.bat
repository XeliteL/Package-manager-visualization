@echo off
echo === Тест 3: Повторяющиеся зависимости ===
go run main.go -package demo_pkg -file tests\stage4\deps_dup.txt -test -mode order
echo -----------------------------------------------