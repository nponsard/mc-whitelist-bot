# mc-whitelist-bot project
 A bot that reads from text channels and add the usernames entered to the mincraft whitelist via RCON 

## Dependencies
- make 
- go
- pandoc for manual generation


## Packaging dependencies
- zip
- dpkg

## make commands
- `make all` : builds for Windows, linux generic and Ubuntu/debian (deb), builds the manuals and put everything in the `publish` folder
- `make` : builds for your current platform


## config 
`~/.mc-whitelist-bot/config.json`
```json
{
        "rcons": [
                {
                        "address": "address:port",
                        "password": "password"
                }
        ],
        "discord": {
                "token": "put_your_token_here",
                "channels": [
                        "put_your_discord_channels_to_monitor_here"
                ]
        }
}
```