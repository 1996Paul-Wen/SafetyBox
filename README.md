# WEB APP Test Cases
```shell

# 测试web app
curl 127.0.0.1:7777/ping

# 创建用户
curl -XPOST 127.0.0.1:7777/api/User/Create -d '{"name": "abcd", "password":"123", "public_key":"-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAnpOYPt0YePUtUzCLTx5ywmqYKsZfJmtN4as5keTWFvD88cE5tT3C\nR7X4itbfGEzi1ITp3+KlMiH6GkMMZtrafodwmPKcdvB9Fc0rWk2NnrmGMBQrcN+J\ngpVbL6oQjgZn7vv3PyHxldg41UPPQsGJTyMCFvNNz4QTHglPmUctJou+uQTBp5YW\npVKtgugDkGx0GhUsRNN3DRSFVtGMSqNO/U0DwbLnJEvJ3jzVtQkqDrMnqyixIbls\nH0jSbS6/judWWroXjzzKzj6IdE3NwH7PbM4vF5mEHW8IHrpHUzTMbohl85rfa/Dr\njVOXV2Kcv8iHbXefNtOITDvpGYkRRSfyzQIDAQAB\n-----END RSA PUBLIC KEY-----"}'

# 查询指定用户详情
curl -XPOST 127.0.0.1:7777/api/User/Describe -H "username:abc"  -H "password:123" -d '{"name": "abc"}'

# 用户存入档案
curl -XPOST 127.0.0.1:7777/api/SafetyData/Create  -H "username:abcd"  -H "password:123" -d '{"archive_key": "key", "archive_value":"value"}'

# 用户查询档案
curl -XPOST 127.0.0.1:7777/api/SafetyData/List  -H "username:abc"  -H "password:123" -d '{"archive_key": "test_key"}'
```