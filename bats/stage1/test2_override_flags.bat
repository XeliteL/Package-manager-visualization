@echo off
echo === Тест 2: Переопределение флагами ===
go run main.go -config tests/stage1/config_valid.json -package override_pkg -max-depth 5 -ascii=false
echo ----------------------------------------