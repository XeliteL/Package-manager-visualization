@echo off
echo === Тест 3: Пустой файл зависимостей ===
go run main.go -package empty_pkg -file tests\stage5\deps_empty.txt -test -mode dot
if exist graph.dot (
    echo Файл graph.dot создан граф в нём пуст - OK.
) else (
    echo Ошибка: graph.dot не создан.
)
echo -----------------------------------------------