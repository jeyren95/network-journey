# Network journey

Application that allows you to view the journey of your network request, documenting the geolocations of the IP hops, before eventually reaching the destination server.

<b>Note: This project is currently a WIP and only works on localhost. Contributions are welcome!</b>

## Current features

- `traceroute` is used to retrieve the IP hops
- https://ip-api.com/ is used to retrieve the geolocations of the IP hops
- Leaflet is used to display the geolocations of the IP hops

## Limitations

`traceroute` is used to retrieve the IP addresses so this project will probably not work for windows users. However, there is a similar command for windows, that can be implemented in the future

## Tech stack

- Go
- Vanilla typescript (no reliance on frameworks to purposely make it more challenging)

## Contributing
Fork the repo, branch out from it, and make a PR whenever you are ready!

## How to start

### Cloning the repo

```
git clone https://github.com/jeyren95/network-journey.git
```

### Starting backend

```
cd network-journey
go run main.go
```

### Starting frontend

```
cd network-journey/web
npm install
npm run start
```

## Future possible implementations
- Cater to windows users
- DNS resolution
