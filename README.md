# digital-pkg/pubsubcli
A really simple CLI for managing local PubSub Emulator.

## Docs

```
Manage topics and subscriptions running on a local PubSub Emulator

Usage:
  pubsubcli [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  publish      Publish message
  subscription Manage subscriptions on PubSub Emulator
  topic        Manage topics on PubSub Emulator

Flags:
  -a, --addr string   PubSub Emulator address (following the pattern <server>:<port>) (default "localhost:8432")
  -h, --help          help for pubsubcli
  -p, --proj string   the project's ID (default "my-project-id")

Use "pubsubcli [command] --help" for more information about a command.
```

## Contributing
This project is based on [Cobra](https://github.com/spf13/cobra). Read the docs and install their CLI called `cobra-cli` ([details here](https://github.com/spf13/cobra#usage)) to make it easy to edit this code.