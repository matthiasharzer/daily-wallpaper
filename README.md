# Daily Wallpaper

An app to update the desktop wallpaper with a daily new picture. Currently supports Bing Image Of The Day only.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
<br>

## Installation

### Binary

Download the [latest release](https://github.com/matthiasharzer/daily-wallpaper/releases/latest) for your platform and run it with the appropriate command-line arguments.

## Usage

### `apply` Command

Update the wallpaper once with todays image.

```bash
daily-wallpaper apply [--market <market>]
```

If the `--market` flag is omitted, the systems locale will be used. The market may influence region specific wallpapers. Example markets are `en-US` and `de-DE`.

### `daemon run` Command

Runs an attached daemon process that checks for a new wallpaper regularly. The default interval is every hour, but can be changed with the `--interval` flag.

```bash
daily-wallpaper daemon run [--market <market>] [--interval <interval>]
```

> Note: The app checks the current wallpaper before downloading a new one. If the wallpaper is already up to date, no download will be performed.


### `version` Command

Check the current version of the app.

```bash
daily-wallpaper version
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details
