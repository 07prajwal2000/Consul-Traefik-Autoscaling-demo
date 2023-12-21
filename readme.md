# Auto-Scale Microservices Demo
An application demonstrates auto-scaling system with the help of Consul by Hashicorp and Traefik (a load-balancer or reverse proxy) and Docker.

## Setup
- Step 1: Build the go `httpserver.go` file and output the binary as `testwebapp` name.
- Step 2: Build the Dockerfile using Docker.
- Step 3: Run the `docker-compose.yml` file.
- Step 4: Open the url `app.localhost:8080`. Here the response will get changed according to the request being routed to the servers.
- Step 5: Thank me 😁.

## Register a new service
- Method: PUT
- Url: `<consul url>/v1/agent/service/register`
- Body: 
    ```json
    {
        // if id not present in paylod, consul will create new service
        "id": "goapp3", // unique id
        "name": "Go Web App", // service name
        "address": "webapp3", // container internal addr
        "port": 3001, // port
        "check": { // health checks url
            "args": [
            "curl",
            "webapp3:3001"
            ],
            "interval": "20s" // time interval for health check
        },
        "tags": [ // imp: traefik prefix is required
            "goapp",
            "api",
            "traefik.enable=true",
            "traefik.http.routers.goapp.rule=Host(`app.localhost`)",
            "traefik.http.services.goapp.loadbalancer.server.port=3001"
        ]
    }
    ```
## De-Register existing service
- Method: Put
- Url: `<consul url>/v1/agent/service/deregister/<service_name>`

## De-Register existing instance from service
- Method: Put
- Url: `<consul url>/v1/agent/service/deregister/<instance_id>`