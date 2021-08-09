# Battery notifier

Just a simple program, which checks every x minutes if a battery is below the specified threshold. If the battery is low (and not being charged) it sends a notification.

Why? I use a i3{,-gaps} for a while, but I haven't set up any "low battery reminder", so I decided to try to create one.

On the screenshot below you may see an example notification popped out by [dunst](https://github.com/dunst-project/dunst) :

![](./dunst_notification.png)

Yeah, I know 94% is not low, but... I've set the threshold to 100% so I could capture a notification as my battery was almost full :D

---

Using these packages:
- [distatus/battery](https://github.com/distatus/battery) for getting battery info 
- [keybase/go-notifier](https://github.com/keybase/go-notifier) for sending a notification


## Install 

```bash
go get gitlab.com/koralowiec/battery-notifier
```

or:

```bash
git clone https://gitlab.com/koralowiec/battery-notifier.git
cd battery-notifier
go build main.go
mv main battery-notifier
```

## Running via cron

```bash
crontab -e
```

To run every minute add something like the line below (remember to adjust path to the binary, optionally you may want to change the threshold via `-t`):

```
*/1 * * * * /home/arek/horsing-around/battery-notifier/battery-notifier -t 40
```
