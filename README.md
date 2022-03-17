# kbd-switch
Simple program to cycle between multiple keyboard layouts (For Linux/Xorg)

# Prerequisites
 - Go programming language ([read this](https://golang.org/doc/install) , [ArchLinux](https://archlinux.org/packages/community/x86_64/go/))
 - notify-send (package libnotify-bin on Debian/Ubuntu, libnotify ArchLinux)
- setkxbmap (package x11-xkb-utils on Debian/Ubuntu, xorg-setxkbmap ArchLinux)

# How to install
**Optimization note**: Since Go 1.18, if you are using an x86_64 CPU you can setup the environment variable `GOAMD64`

You can found a detailed description [here](https://github.com/golang/go/wiki/MinimumRequirements#amd64)

***TL;DR:** Use v1 for any x64 cpu, v2 for circa 2009 "Nehalem and Jaguar", v3 for circa 2015 "Haswell and Excavator" and v4 for AVX-512.*

**Example:** `export GOAMD64=v2`

	git clone https://github.com/lucie-cupcakes/kbd-switch.git
	cd kbd-switch
	make
	sudo make install
	cp config.json.default ~/.config/kbd-switch.json

# How to use it
Edit the config file to your likings, you must type the keyboard layouts as setxkbmap accepts it.

On your startup (.xinitrc or your window manager config), you must add this:
	
	kbd-switch --startup

This command sets the keyboard layout to match the last you used.

On your Window manager config simply add:
	
	kbd-switch

As your favorite key-combination to cycle between keymaps.

