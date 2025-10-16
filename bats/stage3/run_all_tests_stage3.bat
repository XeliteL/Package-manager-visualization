@echo off
echo === Запуск всех тестов Этапа 3 ===
echo ---------------------------
call bats\stage3\test1_tree_basic.bat
call bats\stage3\test2_depth_limit.bat
call bats\stage3\test3_unlimited_depth.bat
call bats\stage3\test4_non_ascii.bat
call bats\stage3\test5_repo_emulation.bat
echo === Все тесты Этапа 3 завершены ===