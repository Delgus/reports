[![codecov](https://codecov.io/gh/Delgus/reports/branch/master/graph/badge.svg)](https://codecov.io/gh/Delgus/reports)
[![Go Report Card](https://goreportcard.com/badge/github.com/delgus/reports)](https://goreportcard.com/report/github.com/delgus/reports)

# Пример API формирования отчетов на POSTGRES + GO


### Запуск демо версии локально c docker-compose

```bash
docker-compose up --build
```

На [localhost:8080](http://localhost:8080) будет доступно приложение

### Разработка

Скомпилировать бинарник 
```bash
make build 
```

Для запуска тестов понадобится утилита go-acc `github.com/ory/go-acc`
```bash
go get github.com/ory/go-acc
```

Форматирование кода всего проекта:
```bash
make fmt
```

Запуск линтеров:
```bash
make lint
```

Запуск тестов:
```bash
make test
```

### Переменные среды 

```env
PG_PORT=5432
PG_HOST=pghost
PG_USER=postgres
PG_PASSWORD=123456
PG_DBNAME=postgres
APP_HOST=
APP_PORT=80
```

## API  

Выбор библиотеки для формирования excel обоснован бенчмарками [тут](https://github.com/Delgus/xls-test)

### Report1

#### JSON

Endpoint: `GET /r1/json`

Response:  
```json
{
  "categories": [
    {
      "name": "Пиццы",
      "products": [
        {
          "name": "4сыра",
          "count": 3,
          "cost_sum": "451.38",
          "sell_sum": "1350.80"
        },
        {
          "name": "Мясное Плато",
          "count": 6,
          "cost_sum": "901.03",
          "sell_sum": "2850.59"
        }
      ],
      "count": 9,
      "cost_sum": "1352.41",
      "sell_sum": "4201.38"
    },
    {
      "name": "Супы",
      "products": [
        {
          "name": "Борщ",
          "count": 3,
          "cost_sum": "90.99",
          "sell_sum": "300.29"
        },
        {
          "name": "Харчо",
          "count": 3,
          "cost_sum": "60.51",
          "sell_sum": "200.59"
        }
      ],
      "count": 6,
      "cost_sum": "151.50",
      "sell_sum": "500.88"
    }
  ],
  "count": 15,
  "cost_sum": "1503.91",
  "sell_sum": "4702.26"
}
```

#### XLSX

Endpoint: `GET /r1/xlsx`

### Report2

#### JSON

Endpoint: `GET /r2/json`

Response:  
```json
{
  "categories": [
    {
      "name": "Пиццы",
      "products": [
        {
          "name": "4сыра",
          "count": 3,
          "cost_sum": "451.38",
          "sell_sum": "1350.80"
        },
        {
          "name": "Мясное Плато",
          "count": 6,
          "cost_sum": "901.03",
          "sell_sum": "2850.59"
        }
      ],
      "count": 9,
      "cost_sum": "1352.41",
      "sell_sum": "4201.38"
    },
    {
      "name": "Супы",
      "products": [
        {
          "name": "Борщ",
          "count": 3,
          "cost_sum": "90.99",
          "sell_sum": "300.29"
        },
        {
          "name": "Харчо",
          "count": 3,
          "cost_sum": "60.51",
          "sell_sum": "200.59"
        }
      ],
      "count": 6,
      "cost_sum": "151.50",
      "sell_sum": "500.88"
    }
  ],
  "count": 15,
  "cost_sum": "1503.91",
  "sell_sum": "4702.26"
}
```

#### XLSX

Endpoint: `GET /r2/xlsx`
