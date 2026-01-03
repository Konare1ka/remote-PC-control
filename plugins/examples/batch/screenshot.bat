@echo off
::screenshot plugin
powershell -Command "Add-Type -AssemblyName System.Windows.Forms; Add-Type -AssemblyName System.Drawing; $scr = [Windows.Forms.SystemInformation]::VirtualScreen; $bmp = New-Object Drawing.Bitmap $scr.Width, $scr.Height; $gfx = [Drawing.Graphics]::FromImage($bmp); $gfx.CopyFromScreen($scr.Location, [Drawing.Point]::Empty, $scr.Size); $bmp.Save('C:\Windows\Temp\screenshot.png')"
echo img C:\Windows\Temp\screenshot.png