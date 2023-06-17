test cases
```shell

curl 127.0.0.1:7777/ping

curl -XPOST 127.0.0.1:7777/api/User/Create -d '{"name": "abc", "password":"123", "public_key":"public_key"}'

curl -XPOST 127.0.0.1:7777/api/User/Describe -H "username:abc"  -H "password:123" -d '{"name": "abc"}'

```