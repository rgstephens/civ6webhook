Civ 6 Notifcation Server in Go

## YAML Config File

`config.yml`

```yaml
users:
  - name: greg
    notify-method: sms
    sms: 206-304-1111
    email: greg@aol.com
  - name: joey
    notify-method: sms
    sms: 404-323-4344
    email: joey@gmail.com
```

## Run

```sh
nohup ./civ6/civ6webhook.linux_amd64  &
```

## ToDo

- [x] Read config file with user id and endpoint id
- [ ] SMS endpoint support via email, [example](https://github.com/acmacalister/sms), [emersion module](https://github.com/emersion/go-smtp#client), [gmail gist](https://gist.github.com/jpillora/cb46d183eca0710d909a)

## Build for another architecture

```sh
go tool dist list
GOOS=linux GOARCH=amd64 go build -o civ6webhook.linux_amd64 civ6webhook.go 
GOOS=linux GOARCH=arm go build -o civ6webhook.linux_arm civ6webhook.go 
GOOS=linux GOARCH=arm64 go build -o civ6webhook.linux_arm64 civ6webhook.go 
```

```sh
scp civ6webhook.linux_amd64 user@host:./civ6
```