@echo off
IF "%1"=="migrate" (
    go run cmd/migrate/main.go
) ELSE IF "%1"=="run" (
    go run cmd/api/main.go
) ELSE (
    echo Invalid command. Use: migrate or run
)