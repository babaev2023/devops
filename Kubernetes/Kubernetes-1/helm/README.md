Добавление секретов в configmap
```yaml
db:
  user: "$DB_USER"
  password: "$DB_PASS"
var1: val1
```
`cat values.yaml | envsubst > values-injected.yaml && mv values-injected.yaml values.yaml`
