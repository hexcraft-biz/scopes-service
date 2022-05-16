# scopes-service

```bash
Testing Flow
$ cp ./.env.example ./.env
$ docker-compose -f dev.yml up --build -d
```

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
      - Desc : Scope Name
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
      - Desc : Resource Domain Name
    - resourceName
      - Required : False
      - Desc : Resource Name
    - name
      - Required : False
      - Desc : Scope Name
      - Format : `scope1` or `scope1|scope2`
    - type
      - Required : False
      - Desc : Scope Type
      - Allow : `public`, `private`
    - limit
      - Required : False
      - Desc : Limit
      - default : 20
    - offset
      - Required : False
      - Desc : Offset
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
      - Desc : Resource Domain Name
    - resourceName
      - Required : True
      - Desc : Resource Name
      - Format : `([a-zA-Z0-9\.\*]+\s?)+`
    - name
      - Required : True
      - Desc : Scope Name
    - type
      - Required : True
      - Desc : Scope Type
      - Allow : `public`, `private`
- Response
  - 200
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
      - Desc : Resource Domain Name
- Response
  - 204
  - 400
  - 401
  - 403
  - 404
  - 500
