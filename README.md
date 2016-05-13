# stow 
[![Build Status](https://travis-ci.org/fabienfoerster/stow.svg?branch=master)](https://travis-ci.org/fabienfoerster/stow)

Stow is a simple app that allow someone to sort his series. To use it you simply have to pass a source folder and a destination folder. The files in the source folder will be match against some regex and move to the destination following the principle :
````
seriename/seasonXX/your-episode.mkv
````

## Utilisation
````bash
stow -src=sourcefolder -dst=destinationfolder
````

## Installation

### Using the binary
Simply go to the [release section](https://github.com/fabienfoerster/stow/releases) and download the binary corresponding to your platform

### Building from source

Simply clone the repo
````git
git clone https://github.com/fabienfoerster/stow.git
````

And build the project
````go
go install
````

Enjoy ʕ”̣̫Ɂ

