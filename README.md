# edwardofclt/server-config

This is how I configure my homelab. I don't use anything other than Docker Compose to manage everything.

## What comes baked in

- Caddy
- Plex
- Overseer
- Deleterr
- Tautulli
- nzbGet
- Sonarr
- Radarr
- Minecraft
- PiHole

## Getting Started

1. Create a `/plex` directory in the root of your file system
2. Clone this repository into `/plex`
3. (Optional) Update the `pihole-updater` .env file. ([See the README for more info](./pihole/local-dns-updater/README.md))
4. Update the `DOMAIN` environment variable in `plex.service` with your domain.
5. Copy `plex.service` to `/etc/systemd/system/plex.service`
6. Run `systemctl daemon-reload`
7. Run `systemclt enable --now plex`

## Optional Setup

Update your router to use your server's IP address as its DNS server. This will allow PiHole to begin filtering out ads AND keeping your traffic internal to your LAN while you're on you're on the network.