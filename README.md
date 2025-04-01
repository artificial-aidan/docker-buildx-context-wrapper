# docker-buildx-env-wrapper

A drop in docker replacement that will force a context argument to be added to the buildx command. Used to get around vscode not allowing builder selection for dev container building

## usage

Install (`default` context)

`go install .`

Install (`different` context)

`go install -ldflags="-X 'main.Context=different'" .`

Configure vscode setting

```json
    "docker.dockerPath": "docker-buildx-context-wrapper",
```
