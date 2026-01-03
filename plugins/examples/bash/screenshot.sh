#!/bin/bash
#screenshot plugin

if command -v flameshot &>/dev/null; then
    flameshot screen -p /tmp/screenshot.png
    echo img /tmp/screenshot.png
else 
    echo Flameshot not found
if