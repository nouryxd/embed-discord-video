# Embed Discord Video

Lets you watch a discord video in the browser instead of downloading it by default.

Simply paste the link to a discord video at the end of the url. 

Like this:  [https://dc.dank.pw/https://cdn.discordapp.com/attachments/1004882628140339281/1039659891310198845/namnam.mp4](https://dc.dank.pw/https://cdn.discordapp.com/attachments/1004882628140339281/1039659891310198845/namnam.mp4)

## Install
### With docker-compose.yml
Clone the repo:
```sh
$ git clone https://github.com/nouryxd/embed-discord-video.git
```

Change the `7352` port value in the `docker-compose.yml` to the port you want to run it on with nginx/caddy.
```yml
ports:
      - "127.0.0.1:7532:8080"
```

Then build it and see if it works.
```sh
$ make build
$ sudo docker compose build
$ sudo docker compose up 
```

Run it in the background.
```sh
$ sudo docker compose up -d
```

### Manually
Clone the repo:
```sh
$ git clone https://github.com/nouryxd/embed-discord-video.git
```
Change the `PORT` value in `cmd/main.go` to the port you want to run it on with nginx/caddy.
```go
var PORT = "8080"
```

Build, and run it.
```sh
$ make build
$ make run
```

Then add it to systemd as autostart or whatever, you know this probably better.
