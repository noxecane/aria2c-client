# aria2c client
Bypass ST_CORP's block on torrents

## How to use

Make sure to install aria2c
`brew install aria2c`

`apt install aria2c`

Install the client
`go get github.com/noxecane/aria2c-client`

Start aria2c in a new terminal or screen
`aria2c --enable-rpc --rpc-listen-all=true --rpc-allow-origin-all`

Add downloads
`aria2c-client <url>`

## The plan
- [x] Build a client that can connect to aria2c on server and add downloads(including torrent)
- [ ] Download list through file. 
- [ ] Add support for configuration
- [ ] Add auto-sync config(for remote downloads)
- [ ] `fast`, `efficient` and `normal` mode
- [ ] User Interface **ge geun**
- [ ] Support for secure downloads over RPC
