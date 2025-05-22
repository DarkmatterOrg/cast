### Base
#### Self
    * Self update command
        * Check what version is installed and where (incase user placed it somewhere else)
        * Download the latest version
        * Replace the old binary with the new version

#### Package Manager
    * Search
        * Add support for more package managers
    * Install
        * Add support for adding custom copr stuff `--copr --coprUrlTHing --coprName` or something
        * Flag to allow AUR?

#### Umbra
    * For Arch, add the custom repo command
    * MCreator install command (If Arch then use custom repo)
    * setup-new-system
        * Flag or other argument should ask for path to a file with a list of apps to be installed
        * Flag for base distro (--arch, --debian etc)
        * The toml file for it should have different categories for things one would want
            * Apps
            * Commands wanting to be ran
    * command to deal with BitFocus companion
        * Unmute and mute mic
        * Change output device
        * Change input device
    * Make the chall command in python instead?

### Possible ideas
    * Wrapper for ytdlp
    * Twitch stuff like that one cli https://github.com/krathalan/wtwitch