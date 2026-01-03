package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var pluginList []string
var pluginPath string

func init() {
	pluginList = *getPluginsList()
}

func pluginHandler(uname* string, com* string, args *[]string) (bool, []byte) {
	pluginList := getPluginsList() //we get every time we call to dynamically add/remove plugins
	isAdmin := checkPermission(*uname)
	err, output := execPlugin(isAdmin, com, pluginList, args)
	return err, output
}

func getPluginsList() *[]string {

	// get path to plugins
	execPath, err := os.Executable()
	if err != nil {logger.Error("Failed get a execPath", "error", err.Error()); os.Exit(1)}
	pluginPath = filepath.Join(filepath.Dir(execPath), "plugins")

	// get files list in directory
	entries, err := os.ReadDir(pluginPath)
	if err != nil { logger.Error("Error when getting a list of files in the plugins directory", "error", err.Error()); os.Exit(1)}
	for _, entry := range entries { // check files in plugins directory
		if !entry.IsDir() { pluginList = append(pluginList, entry.Name()) }
	}
	if len(pluginList) == 0 { logger.Error("No plugins") }
	return &pluginList
}

func checkPermission(uname string) bool {
	if uname == cfg.Username {
		logger.Info("User is privileged")
		return true
	} else {return false}
}

func execPlugin(isAdmin bool, msg* string, plugList *[]string, args *[]string) (bool, []byte) {
	ext := checkOS()
	var output []byte
	var err error
	if stringInSlice(*msg + ext, *plugList) { //checking if the command is in the list
		if isAdmin {
			cmdPath := filepath.Join(pluginPath, *msg + ext) //make command
			cmd := exec.Command(cmdPath, *args...)
			output, err = cmd.CombinedOutput()
		} else { //for no admins we will additionally check whether plugin is in public list
			if stringInSlice(*msg, cfg.AllowPlugList) {
				cmdPath := filepath.Join(pluginPath, *msg + ext)
				cmd := exec.Command(cmdPath, *args...)
				output, err = cmd.CombinedOutput()
			} else {
				logger.Warn("Permission denied", "cmd", *msg)
				return true, nil
			}
		}
		if err != nil { 
			logger.Error("Failed executable plugin", "error", err)
			return true, nil
		} else {
			logger.Info("Plugin called", "plugin", *msg + ext)
			return false, output
		}
	} else {
		logger.Warn("Plugin not found", "plugin", *msg + ext)
		return true, nil
	}
}

func checkOS() string {
	if runtime.GOOS == "windows" { return ".bat" } else { return ".sh" }
}

//func to check an array to see if it contains what you're looking for (like "in" in python)
func stringInSlice(a string, list []string) bool {
    for _, b := range list { if b == a { return true } }
    return false
}