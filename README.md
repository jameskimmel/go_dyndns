# go_dyndns


This program does DynDNS updates for deSEC.io
Please use it with caution, I am not a coder, this program is probably horrible!

What it does currently:
- Wizard will create a config.json file (including your unencrypted token)
- Uses deSEC services to determine your IP(s)
- Always sets the IP(s) in the update URL to prevent MITM attacks
- Always uses HTTPS (thanks to GO)
- Last set IP(s) and last update time will be cached in the config file
- Only ask for an update if something actually has changed
- Only ask for an update if the last update is at least 5min old
- Instead of checking for an IP, you can also hardcode it in the config.json file. This is only useful for special edge cases. 

What is missing:
- Can only check a single domain
- Only uses desec.io services to determine your IP(s), no other services can be selected
- Only works for deSEC.io
- input validation
- some kind of timer to start the service.
- a Windows build
- code review from someone that actually knows how to code :)
