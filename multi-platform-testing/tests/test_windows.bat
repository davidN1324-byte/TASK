@echo off
echo Running Windows tests...
build\app.exe
IF %ERRORLEVEL% NEQ 0 (
    echo ❌ Tests failed!
    exit /b 1
)
echo ✅ Tests passed!
