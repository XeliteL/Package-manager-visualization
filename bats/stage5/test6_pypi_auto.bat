@echo off
chcp 65001 > nul
echo === Тест 6: Загрузка зависимостей с PyPI без явного указания URL ===

set "EXE=..\..\pkgvis.exe"
set "RESULTS_DIR=.\results"

if exist "%RESULTS_DIR%" (
    rmdir /S /Q "%RESULTS_DIR%"
)
mkdir "%RESULTS_DIR%"

if not exist "%EXE%" (
    set "RUN_CMD=go run .\main.go"
) else (
    set "RUN_CMD=%EXE%"
)

call :Analyze matplotlib
call :Analyze flask
call :Analyze numpy
call :Analyze fastapi
call :Analyze requests

echo -----------------------------------------------
exit /b

:Analyze
set "PKG=%~1"
echo -----------------------------------------------
echo Обработка пакета: %PKG%
%RUN_CMD% -package %PKG% -mode dot

if exist graph.dot (
    move /Y "graph.dot" "%RESULTS_DIR%\%PKG%.dot" > nul
    echo Файл "%RESULTS_DIR%\%PKG%.dot" успешно создан.
    dot -Tpng "%RESULTS_DIR%\%PKG%.dot" -o "%RESULTS_DIR%\%PKG%.png"
    if exist "%RESULTS_DIR%\%PKG%.png" (
        echo Файл "%PKG%.png" успешно создан.
    ) else (
        echo Ошибка при создании изображения для %PKG%.
    )
) else (
    echo graph.dot не найден. Возможно, пакет "%PKG%" не имеет зависимостей.
)