@echo off
echo === Тест 2: Ограничение глубины анализа (max_depth = 1) ===
go run main.go -package demo_pkg -file tests\stage4\deps_dup.txt -test -mode order -max-depth 1
echo -----------------------------------------------