@echo off
::plugin for checking if a process is running
if not "%~1" == "" (
    tasklist | find "%1" >nul
    if %errorlevel%==0 (
        echo %1 launched!
    ) else (
        echo %1 not found.
    )
) else ( 
    echo Specify name of process
    echo Ex. /system cmd.exe
    )