# Nacos Config

## api => API_GROUP

> dataid => api
> group => API_GROUP

```json
{
  "name": "api",
  "host": "your host",
  "port": 9900,
  "jwt": {
    "key": "you key"
  },
  "otel": {
    "endpoint": ":4317"
  },
  "auth_srv": {
    "name": "auth_srv"
  },
}
```

## auth_srv => AUTH_GROUP

> dataid => auth_srv
> group => AUTH_GROUP

```json
{
  "name": "auth_srv",
  "host": "your host",
  "mysql": {
    "host": "127.0.0.1",
    "port": 3306,
    "user": "root",
    "password": "123456",
    "db": "",
    "salt": ""
  },
  "otel": {
    "endpoint": ":4317"
  }
}
```
