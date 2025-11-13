@echo off
REM build\build.bat - build helper for Windows cmd.exe
cd /d %~dp0\..
if not exist bin mkdir bin
echo Building go-reference into .\bin\goref.exe
go build -o bin\goref.exe .
echo Build complete: .\bin\goref.exe
