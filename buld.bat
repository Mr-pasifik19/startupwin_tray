go get github.com/akavel/rsrc
rsrc -manifest def-manifest.xml -o rsrc.syso

go build -ldflags="-H windowsgui"


@echo off
SETLOCAL

REM Registry key for startup
set KEY="HKCU\Software\Microsoft\Windows\CurrentVersion\Run"

REM Name for the startup entry
set ENTRY_NAME="StartupWin"

REM Path to the executable
set EXECUTABLE_PATH="%CD%\startupwin.exe"

REM StartIn (Working Directory) for the executable
set WORKING_DIRECTORY="%CD%"

REM Add entry to registry with StartIn value
reg add %KEY% /v %ENTRY_NAME% /t REG_SZ /d %EXECUTABLE_PATH% /f

echo "Startup entry added successfully."
