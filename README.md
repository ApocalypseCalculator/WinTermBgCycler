# cyclebg

Cycle through various backgrounds in Windows Terminal.

You can use this by doing `go run main.go` or by building the executable.

This is a port from my original Node.JS code [here](https://github.com/ApocalypseCalculator/WinTermBgCycler-JS).

I will be maintaining this version instead since it should be more lightweight.

Configuration options: 
- `wintermsettings` : Location of the Windows Terminal Settings JSON file, if empty, uses default location at `LOCALAPPDATA\Packages\Microsoft.WindowsTerminal_8wekyb3d8bbwe\LocalState\settings.json`
- `picspathfile` : The file that stores the paths to images you wish to cycle through. Changes to the content will be applied without restarting this program. 
- `picsfolder` : If you wish to include images directly, you can place your pictures inside the folder specified here. 
- `interval` : The interval of cycling backgrounds specified in seconds. Default is 120 seconds. 

To install this, download the executable in the releases tab. The executable will automatically create configuration files in the directory it is executed in.
