

マイグレーションの作成
```sh
migrate create -ext sql -dir infra/migrations -seq create_cars_table
```

マイグレーションの実行
```sh
migrate -database "postgres://test_user:test_password@db/test_db?sslmode=disable" -path ./infra/migrations up
```
