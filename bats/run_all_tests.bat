@echo off
echo === Запуск всех тестов ===
echo ---------------------------

call bats\test1_valid.bat
call bats\test2_override_flags.bat
call bats\test3_invalid_url.bat
call bats\test4_missing_package.bat
call bats\test5_test_mode_valid.bat
call bats\test6_test_mode_invalid.bat
call bats\test7_negative_depth.bat

echo === Все тесты завершены ===
pause