# Soteria

A simple bot that sends files to a user.


## Usage

1. Create a bot by talking to https://telegram.me/BotFather. Copy down the TOKEN.
2. compile and run this binary, passing the token as a flag (`--token` or environment variable (`TOKEN=xxx`)).
   This will start up the bot in echo mode. Start a chat with your bot and make note of the chatID printed in the logs.
3. Now you have everything you need. Call the bot with the token and chatID set, and a path to a file: `./soteria --token=xxx --chatID=123 --file=tests/hello_world`

```
Usage of ./soteria:
      --chatid int    [env: `CHATID`] Telegram chatID for your bot + user.
  -f, --file string   [env: `FILE`]   path for file to send
  -t, --token string  [env: `TOKEN`]  Telegram bot token.
```


## What's in a name?

[Soteria](https://en.wikipedia.org/wiki/Soteria_(mythology)) is the Greek Godess of Deliverance. Naming this bot TGFU
 (TeleGram File Uploader) just didn't have the same ring to it. Hence, Soteria lives again in the modern age.

## Contributing

Clone this repo. Run `make init` to get all the tools and deps. `make build` will create a [upx](https://upx.github.io/) minified binary.