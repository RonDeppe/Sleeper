<h1 align="center">Sleeper💤</h1>

<h5 align="center">Command-line tool for sleeping and logging sleep events</h5>

<p align="center">
  <a href="https://github.com/RonDeppe/sleeper/releases">
        <img src="https://img.shields.io/github/v/release/RonDeppe/sleeper?logo=github" alt="Latest Release">
    </a>
    <a href="https://github.com/RonDeppe/sleeper/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/RonDeppe/sleeper" alt="License">
    </a>
  <a href="https://github.com/RonDeppe?tab=repositories&q=&type=&language=go&sort=">
        <img src="https://img.shields.io/github/languages/top/RonDeppe/sleeper?logo=go&label=" alt="Top Language">
    </a>
</p>

Sleeper is a lightweight Command-line tool written in <code>Go</code> that puts your Windows machine into sleep (suspend) mode. It logs the date, time, and duration of each sleep session, providing a simple and efficient way to monitor your system's sleep activity.

## Instalation
1. Dowload the program directly from [the GitHub releases page](https://github.com/RonDeppe/sleeper/releases/latest).
2. Add a folder for scripts to your PATH variable. Usually this is `C:\Users\{Username}\bin`.
3. Place `sleeper.exe` in this folder.

## Usage
- Run the program by typing `sleeper.exe` in the terminal.

## FAQ
### Q: What platforms are supported
Currently only windows 11 has been tested, but older versions of windows should work.

### Q: Where is the data stored?
The data is located at `C:\Users\{Username}\AppData\Local\sleeper\data.csv`.
