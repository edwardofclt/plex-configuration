# local-dns-updater

This basically scrapes your Docker containers for any Caddy label-based configurations. If detected, it takes the identified hostnames and compares them to the local DNS in pihole. It reconciles the differences so you can keep your traffic local when you're in your LAN.

## Getting Started

1. Create a copy of the .env.example file:
    ```
    cp .env.example .env
    ```

1. Update the values:

    | Key             | Required | Description                                                                                                                                                 |
    | --------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
    | PASSWORD        | Yes      | Generate this value in pihole                                                                                                                               |
    | PIHOLE_HOST     | Yes      | Hostname your pihole is running on                                                                                                                          |
    | TARGET_IP       | Yes      | What is the IP of the server your container is running on?                                                                                                  |
    | MONITOR_DOMAINS | No       | Comma-delimited list of domains. If this value is set, any values for the domain you list will be cleaned up if not found within your caddy configurations. |

1. Run the stack

    ```
    docker compose up -d
    ```