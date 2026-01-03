This program was created for remote control of a computer via a telegram bot, using plugins written in a scripting language (bash/batch)

> [!WARNING]
>
> I do not guarantee that this application and its code are perfect or standardized. This is  my first golang project and, in general, the first full-fledged project. If you want to improve the code, then upload pull requests
>
## Build
Specify NAME yourself (for Windows, add .exe extension to name)
```
git clone https://github.com/Konare1ka/remote-PC-control
cd remote-PC-control/src
go build -o ../NAME .
```

>If by some miracle the third-party library did not catch up\
>`go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5`

## Settings
In same directory where executable file is located there should be config.json and plugins directory with plugins

config.json must contain a field with:
- telegram bot token, which is issued by [BotFather](https://telegram.me/BotFather)
- your username (without @) to receive admin privileges
- plugins available to everyone

## Usage 
To simple use, run binary without additional arguments

The help command is also available `./bin help`\
There is information for creating plugins, setting up configs, etc.

Messages may arrive with a delay, before you decide that nothing is working for you, wait a minute

## Plugins
Plugins are scripts written in bash/batch (depending on your OS) that are called when you enter a command in telegram bot that matches plugin name
>Ex. file `help.sh` - command in bot `/help`

Plugins can only be accessed by the user specified in config. However, publicly accessible plugins can be specified in third config field,
> like ["help", "getFile", "music"]
> 
The second line of plugin should be a comment describing plugin

### Help plugin
The help plugin scans all scripts in a directory and outputs them to telegram bot in the format 
`/name - description`
You can replace or improve the script, but this plugin allows you to dynamically retrieve available plugins and their descriptions.

### Return value of plugins
Any echo command in plugin will be echoed to telegram bot. To send any media or document (including .jar, .zip, .exe, etc.), you need to send the message type (image-img, video-vid, audio-aud) followed by a space and full path to file. 

#### There should be no spaces in the path

>Ex. img /home/user/Pictures/image.png
>

### Usage arguments
You can pass arguments to the plugin; to do this, when calling, you need to specify the arguments separated by a space
`/plugin arg1 arg2`
And to process arguments in the plugin you need to use %1 - arg1, %2 - arg2, etc.
**I recommend using a check for the presence of these arguments, and also checking whether everything was done as it should. Since the program does not monitor how the script was executed, it only monitors whether it was executed**
>An example of using arguments is in the examples, the name of the script is shareFiles
>

## In future

In the future I will add a system for automatically creating a service/daemon



