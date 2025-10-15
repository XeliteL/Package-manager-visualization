@echo off
echo === Запуск всех тестов ===
echo ---------------------------

call bats\stage1\test1_valid.bat
call bats\stage1\test2_override_flags.bat
call bats\stage1\test3_invalid_url.bat
call bats\stage1\test4_missing_package.bat
call bats\stage1\test5_negative_depth.bat
call bats\stage2\test6_repo_mode.bat
call bats\stage2\test7_test_mode_valid.bat
call bats\stage2\test8_test_mode_invalid.bat
call bats\stage2\test9_ascii_mode.bat

echo === Все тесты завершены ===