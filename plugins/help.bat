@echo off
::plugin for display available plugins
setlocal enabledelayedexpansion

cd plugins

for %%I in (*.bat) do (
    < "%%I" (
        set /p firstLine=
        set /p secondLine=
    )
    set firstTwoSymb=!secondLine:~0,2!
    if "!firstTwoSymb!"=="::" (
        echo /%%~nI - !secondLine:~2!
    )
)

echo Created by Konare1ka