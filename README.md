# go-kindle -- Quick & dirty library to write stuff on a Kindle screen

## Description

This Go library encapsulates calls to various system executables on a Kindle 4 system. This allows developers to manipulate the e-ink screen, read key presses and read settings.

The following applications are wrapped:

- `eips`: clear the screen, draw an image, write text
- `waitforkey`: self-explanatory, waits for a key to be pressed and returns it
- `lipc`: a sort of inter-process communication system, can be used to get information such as battery level or WiFi state.

## Compiling for the Kindle

The Kindle 4 runs on Linux, under an ARMv5 processor. Compiling your application is as easy as making sure the following environment variables are set:

- `GOOS=linux`
- `GOARCH=arm`
- `GOARM=5`

The include `deploy.bat` script will do this for you, and copy your app as a KUAL extension to your Kindle.

## Running the application

While developing, it can be useful to simply run the application from your SSH session. You might see the Kindle interface popping back in sometimes, because your application and the UI are fighting for framebuffer access. The included KUAL configuration (in `app/`) contains a script which will kill the interface, run your application and restart the interface when it exits.

This is slow because you're basically rebooting the Kindle every time, but it fixes everything.