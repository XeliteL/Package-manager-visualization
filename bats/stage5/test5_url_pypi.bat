@echo off
chcp 65001 > nul
echo === Тест 5: Реальный пакет из PyPI (matplotlib) ===
setlocal

REM Путь к исполняемому файлу (если программа уже собрана)
set EXE=..\..\pkgvis.exe

REM Если запускаем без сборки — используем go run
if not exist %EXE% (
    set RUN_CMD=go run main.go
) else (
    set RUN_CMD=%EXE%
)

echo Режим: dot
echo Получение зависимостей с https://pypi.org/pypi/matplotlib/json ...
%RUN_CMD% -package matplotlib -repo https://pypi.org/pypi/matplotlib/json -mode dot

if exist graph.dot (
    echo Файл graph.dot успешно создан.
    echo Можно визуализировать граф командой:
    echo dot -Tpng graph.dot -o graph.png
) else (
    echo Ошибка: файл graph.dot не создан.
)

echo -----------------------------------------------