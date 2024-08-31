@echo off
REM Run TailwindCSS command
echo Running TailwindCSS...
call npx tailwindcss -i styles/style.css -o static/tailwind.css
if %errorlevel% neq 0 (
    echo Failed to run TailwindCSS command.
    exit /b %errorlevel%
)

REM Run templ generate command
echo Generating templates...
call templ generate
if %errorlevel% neq 0 (
    echo Failed to generate templates.
    exit /b %errorlevel%
)

REM Run Go application
echo Starting Go application...
call go run main.go
if %errorlevel% neq 0 (
    echo Failed to start Go application.
    exit /b %errorlevel%
)