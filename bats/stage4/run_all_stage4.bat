@echo off
echo === Запуск всех тестов Этапа 4 ===
call bats\stage4\test1_order_basic.bat
call bats\stage4\test2_order_depth.bat
call bats\stage4\test3_order_duplicates.bat
call bats\stage4\test4_repo_mode.bat
echo === Все тесты Этапа 4 завершены ===