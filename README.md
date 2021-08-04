# kbd-switch
Simple program to cycle between multiple keyboard layouts (For Linux/Xorg)

# Prerequisites
 - Go programming language ([read this](https://golang.org/doc/install) , [ArchLinux](https://archlinux.org/packages/community/x86_64/go/))
 - notify-send (package libnotify-bin on Debian/Ubuntu, libnotify ArchLinux)
- setkxbmap (package x11-xkb-utils on Debian/Ubuntu, xorg-setxkbmap ArchLinux)

# How to install

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
