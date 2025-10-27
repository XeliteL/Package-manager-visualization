@echo off
echo === Запуск всех тестов Этапа 5 (визуализация графа) ===
echo -----------------------------------------------
call bats\stage5\test1_basic_dot.bat
call bats\stage5\test2_max_depth_1.bat
call bats\stage5\test3_empty_file.bat
call bats\stage5\test4_repo_dot.bat
call bats\stage5\test5_url_pypi.bat
call bats\stage5\test6_pypi_auto.bat
echo === Все тесты Этапа 5 завершены ===