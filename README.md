# Battery notifier

Just a simple program, which checks every x minutes if a battery is below the specified threshold. If the battery is low (and not being charged) it sends a notification.

Why? I use a i3{,-gaps} for a while, but I haven't set up any "low battery reminder", so I decided to try to create one.

On the screenshot below you may see an example notification popped out by [dunst](https://github.com/dunst-project/dunst) :

![](./dunst_notification.png)

Yeah, I know 94% is not low, but... I've set the threshold to 100% so I could capture a notification as my battery was almost full :D

---

Using those packages:
- [distatus/battery](https://github.com/distatus/battery) for getting battery info 
- [keybase/go-notifier](https://github.com/keybase/go-notifier) for sending a notification


## Install 

```bash
go get gitlab.com/koralowiec/battery-notifier
```

## Running as a systemd unit

Clone the repo:

```
git clone https://gitlab.com/koralowiec/battery-notifier.git
cd battery-notifier
```

Open the service file with a text editor:

```
vim ./systemd-unit/battery-notifier.service
```

Change the unit file (the line below):
 - replace /home/arek/go with your $GOPATH
 - (optional) change passed arguments (check `battery-notifier -h`)

```ini
ExecStart=/home/arek/go/bin/battery-notifier -t 35 -m 2
```

Copy the file to user's units, reload the configuration, start and enable the service:

```
cp ./systemd-unit/battery-notifier.service ~/.config/systemd/user/battery-notifier.service
systemctl daemon-reload --user
systemctl --user start battery-notifier.service
systemctl --user enable battery-notifier.service
```
