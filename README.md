# al

AL is a cli tool for open/close your apps and web URLs.

## Installation

```bash
go install github.com/meinbaumm/al@latest
```

```bash
[~]$ al

NAME:
   al - Open/close web urls/apps

USAGE:
   al [global options] command [command options] [arguments...]

VERSION:
   0.1.3

AUTHOR:
   Maxim Petrenko <meinbaumm@gmail.com>

COMMANDS:
   web      Open web urls
   open     Open apps
   close    Close apps
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## How to setup?

Now AL support 3 command.

```
web      Open web urls
open     Open apps
close    Close apps
```

So to use AL you need to specify your own config file like this.

```yml
urls:
  cambridge: "https://dictionary.cambridge.org/"
  github: "https://github.com/"
  facebook: "https://www.facebook.com/"
apps-to-open:
  safari: "/../../Applications/Safari.app/"
  settings: "/../../System/Applications/System Settings.app/"
  calendar: "/../../System/Applications/Calendar.app/"
  photos: "/../../System/Applications/Photos.app/"
  books: "/../../System/Applications/Books.app/"
apps-to-close: 
  safari: quit app "Safari"
  settings: quit app "System Settings"
  calendar: quit app "Calendar"
  photos: quit app "Photos"
  books: quit app "Books"
```

## How AL decides what config file to use?

All will look for the path in AL_CONFIG env variable.
Just add this line to your .profile file:

```bash
export AL_CONFIG="/home/username/al-config.yml"
```

## Examples of usage

`al help` – shows help command

```bash
NAME:
   al - Open/close web urls/apps

USAGE:
   al [global options] command [command options] [arguments...]

VERSION:
   0.1.3

AUTHOR:
   Maxim Petrenko <meinbaumm@gmail.com>

COMMANDS:
   web      Open web urls
   open     Open apps
   close    Close apps
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

- `al open safari calendar books` – will open safari, calendar and books applications.
- `al close safari calendar books` – will close opened applications
- `al web github facebook` – will open github and facebook websites in your default browser

For any command, there is a `list` subcommand to see list of available apps/urls in your config file. For example  

- `al web list` – will print all url names from your config file in ascending order. If you want to see not only url names, but urls too, use `--verbose` flag.
