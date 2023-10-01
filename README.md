# Discord.go ai (`gpt3.5-turbo`)

To get started:

1. Clone this repository
2. Fill .env with [your values](#configuration)
3. Build (`go build`) the app and start it (`./dgo-ai`)

## Running

### with docker
```bash
docker run -d \
    -e TOKEN=your-token-change-this \
    -e PRESENCE=DM or mention for AI response \
    -e INJECT_USER_PROMPT=... \
    -e INJECT_SYSTEM_PROMPT=... \
    ghcr.io/neu-ma-tic/discord-ai:master
```

### containerless
```bash
# assuming unix-like host
go build
# builds as `dgo-ai` by default
./dgo-ai
```

## Configuration
- `TOKEN`: The bot's token
- `PRESENCE_TEXT`: The bot's rpc text
- `INJECT_USER_PROMPT`: Text to inject as user role as chat context
- `INJECT_SYSTEM_PROMPT`: Text to inject as system role as chat context