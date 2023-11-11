# scopes-service

```bash
Set up env file
$ cp ./.env.example ./.env
# Don't forget change your mysql connection info at .env file.
```

- Standalone Testing Flow
```bash
$ docker-compose -f dev.yml up --build -d
```
- Integration testing
  - Please flollow [hexc-deploy](https://github.com/hexcraft-biz/hexc-deploy) README.md step.

## End Point

### GET /healthCheck/v1/ping
- Params
  - None
- Response
  - 200

### GET /scopes/v1/scopes/:name
- Params
  - Uri
    - name
      - Required : True
- Response
  - 200
  - 400
  - 401
  - 403
  - 404
  - 500

### GET /scopes/v1/scopes
- Params
  - QueryString
    - resourceDomainName
      - Required : False
    - resourceName
      - Required : False
    - name
      - Required : False
      - Format : `scope1` or `scope1|scope2`
    - type
      - Required : False
      - Allowed : `public`, `private`
    - limit
      - Required : False
      - default : 20
    - offset
      - Required : False
      - default : 0
- Response
  - 200
  - 400
  - 401
  - 403
  - 500

### POST /scopes/v1/scopes
- Params
  - Headers
    - Content-Type : application/json
  - Body
    - resourceDomainName
      - Required : True
    - resourceName
      - Required : True
      - Format : `([a-zA-Z0-9\.\*]+\s?)+`
    - name
      - Required : True
    - type
      - Required : True
      - Allowed : `public`, `private`
- Response
  - 201
  - 400
  - 401
  - 403
  - 409
  - 500

### DELETE /scopes/v1/scopes
- Params
  - QueryString
    - resourceDomainName
      - Required : True
- Response
  - 204
  - 400
  - 401
  - 403
  - 404
  - 500
