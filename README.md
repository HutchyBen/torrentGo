# torrentGo
This tool is created to download torrents for several sources (not several rn). It is made to be simple to use and doesn't require a BitTorrent client as it as it uses [a built in BitTorrent client](https://github.com/anacrolix/torrent) 

## Setup
The only needed thing you will need to do is to build the project with `go build`. It should then be ready to use.

## Usage
To download a torrent from 1337x it is very simple
`torrentGo 1337x ubuntu`
You should then be presented by a menu to select a torrent
```
? Choose a torrent:  [Use arrows to move, type to filter]
> Ubuntu MATE 16.04.2 [MATE][armhf][img.xz][Uzerus]2 By Uzerus
  A Practical Guide to Ubuntu Linux, 3rd Edition (PDF) By bookflare
  Ubuntu Unleashed 2019 Edition By spy1984
  Ubuntu 16.04.1 LTS Desktop 64-bit By SeuPirate
  Ubuntu 20.04 User Guide By putrosllio
  Udemy - The Complete Ubuntu Linux Server Administration Course1 By tutsgalaxy
  Ubuntu 16.10 LTS [Yakkety Yak][Unity][x32 i386][Desktop][ISO][Uzerus] By Uzerus
```
You can then select a torrent of which it will be downloaded

## Roadmap
- [ ] Add support for more sources
- [ ] Learn how go modules actually work
- [x] Add more info to torrents
- [ ] Add support for more pages to 1337x
