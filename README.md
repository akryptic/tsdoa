# TSdoA

TSdoA is a task management application built with Wails, combining a Go backend with a Svelte TypeScript frontend. It provides an intuitive interface for managing tasks and their associated steps.

## Download

Get the latest version of TSdoA from our GitHub releases:

### 📥 [Download Latest Release](https://github.com/akryptic/tsdoa/releases/latest)

**Available platforms:**
- 🪟 **Windows** - `.exe` installer (NSIS)
- 🐧 **Linux** - `.AppImage`, `.deb` and binary
- 🍎 **macOS**: 
    - [x] arm64
    - [x] arch64
    - [x] universal (Apple Silicon / Intel)
- 🤖 **Android** - *Coming soon*

### Installation Instructions

**Windows:**
1. Download the `.exe` file from releases
2. Run the installer and follow the setup wizard
3. Launch TSdoA from your Start Menu

**Linux:**
1. Download the .AppImage or .deb or binary
2. Make it executable: `chmod +x tsdoa`
3. Run the application: `./tsdoa`

**macOS:**
1. Download the macOS `.app.zip` or `.pkg`
2. Extract and move to Applications folder
3. Right-click and select "Open" (may need to allow in Security preferences)


## Tech Stack

- **Backend**: Go with Wails v2
- **Frontend**: Svelte with TypeScript
- **Build Tool**: Vite
- **Styling**: Custom CSS with Space Grotesk font
- **Database**: Local storage via Go backend

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and ensure the build passes
5. Submit a pull request

## License

[MIT](LICENSE)

## Acknowledgments

- Built with [Wails](https://wails.io/) - Go + [Svelte](https://svelte.dev) - Web frontend framework
- Icons powered by [Lucide](https://lucide.dev/) icons
- Font: [Space Grotesk](https://fonts.google.com/specimen/Space+Grotesk)
