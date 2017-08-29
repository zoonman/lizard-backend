
```bash
curl -H "Content-Type: application/json" -X POST \
    -d '{"email":"blackhole@zoonman.com","password":"xyz"}' \
    http://127.0.0.1:10123/api/v1/users
```


```bash
curl -H "Content-Type: application/json" -X POST \
    -d '{"bmi":1.05}' \
    http://127.0.0.1:10123/api/v1/users/3/records
```


```bash
curl http://127.0.0.1:10123/api/v1/users
```

```bash
curl http://127.0.0.1:10123/api/v1/users/3/records
```
