package main

import (
	"runtime"
	"fmt"
)

func printMainHelp() {
	fmt.Println("Created by Konare1ka")
	fmt.Println("\"help\" - list arguments")
	fmt.Println("\"help plugin\" - help related to plugins")
	fmt.Println("\"plugin\" - create a help plugin")
}

func printPluginHelp() {
	fmt.Println("Plugins should be located in the plugin directories, in the same place as the executable file")
	fmt.Println("The additional files must match your wireless system")
	fmt.Println("To connect who can use others using a bot, specify in config.json, in the allowPlugins field,")
	fmt.Println("the name of these plugins without extensions (for example, plugin - example.bat, json - “example”")
	fmt.Println("To send an image/video/document(also exe, jar, zip, etc.)/audio to the bot, you need to indicate with the last echo")
	fmt.Println("that you want to send img/vid/doc/aui and, separated by a space, indicate the full path to the file")
	fmt.Println("(do not forget to indicate the extension Ex. C:\\Users\\PC\\Documents\\doc.txt)")
	fmt.Println("Any echo in the plugin will be output to bot in chat")
	fmt.Println("Line 2 must be a comment describing the plugin")
}

func printHelpPlugin() {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("@echo off")
		fmt.Println("::plugin for display available plugins")
		fmt.Println("setlocal enabledelayedexpansion")
		fmt.Println("cd plugins")
		fmt.Println("for %%I in (*.bat) do (")
		fmt.Println("    < \"%%I\" (")
		fmt.Println("        set /p firstLine=")
		fmt.Println("        set /p secondLine=")
		fmt.Println("    )")
		fmt.Println("    set firstTwoSymb=!secondLine:~0,2!")
		fmt.Println("    if \"!firstTwoSymb!\"==\"::\" (")
		fmt.Println("        echo /%%~nI - !secondLine:~2!")
		fmt.Println("    )")
		fmt.Println(")")
		fmt.Print("\n\nPaste the code into help.bat, which is located in the plugins directory")
	case "linux":
		fmt.Println("#!/bin/bash")
		fmt.Println("#plugin for display available plugins")
		fmt.Println("cd plugins")
		fmt.Println("for file in *.sh; do")
		fmt.Println("    if [[ -f \"$file\" ]]; then")
		fmt.Println("        first_line=$(head -n 1 \"$file\")")
		fmt.Println("        second_line=$(head -n 2 \"$file\" | tail -n 1)")
		fmt.Println("        if [[ \"$second_line\" == \\#* ]]; then")
		fmt.Printf("            filename=\"${file%%.sh}\"\n") //printf because it swears at "%.s" 
		fmt.Println("            echo \"/$filename - ${second_line:1}\"")
		fmt.Println("        fi")
		fmt.Println("    fi")
		fmt.Println("done")
		fmt.Print("\n\nPaste code into help.sh, which is located in plugins directory")
	}
}

func printConfigHelp() {
	fmt.Println("{")
	fmt.Println("	\"token\":\"8466329933:AAE1veE5PwAStOtd-1e5AF4r1OYJKgrbHdw\",")
	fmt.Println("	\"userID\": 5430442454,")
	fmt.Println("   \"allowPlugins\":[\"help\"]")
	fmt.Println("}")
	fmt.Print("\n\nPaste code into config.json, which is located in root directory")
	fmt.Println("In field \"allowPlugins\" you must specify the file name without extension, \nthese plugins will be available to all users")
}