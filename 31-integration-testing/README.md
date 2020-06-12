## Пример использования docker-compose и godog для интеграционного тестирования микросервисов

Все команды выполняются из корня проекта.

### Запуск микросервисов
Поднимаем docker-compose
```shell script
$ make up
docker-compose up -d --build
Creating network "31-integration-testing_db" with driver "bridge"
Creating network "31-integration-testing_rabbit" with driver "bridge"
Building notification_service
...
Successfully built 8fc475a1a227
Successfully tagged godog_example_notification_service:latest
Building registration_service
...
Successfully built 69bcf25006a9
Successfully tagged 31-integration-testing_registration_service:latest
Creating 31-integration-testing_postgres_1 ... done
Creating 31-integration-testing_rabbit_1   ... done
Creating 31-integration-testing_notification_service_1 ... done
Creating 31-integration-testing_registration_service_1 ... done
```

Проверяем, что все сервисы поднялись.
```shell script
$ docker-compose ps
                    Name                                   Command               State                                             Ports
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
31-integration-testing_notification_service_1   /bin/notify_service              Up
31-integration-testing_postgres_1               docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
31-integration-testing_rabbit_1                 docker-entrypoint.sh rabbi ...   Up      15671/tcp, 0.0.0.0:15672->15672/tcp, 25672/tcp, 4369/tcp, 5671/tcp, 0.0.0.0:5672->5672/tcp
31-integration-testing_registration_service_1   /bin/reg_service                 Up      0.0.0.0:8088->8088/tcp```
```

Проверяем доступность сервиса регистрации
```shell script
$ curl http://localhost:8088/
OK
```

Регистрируем пользователя
```shell script
$ curl -d '{"first_name":"otus", "email":"otus@otus.ru", "age": 27}' -H "Content-Type: application/json" -X POST http://localhost:8088/api/v1/registration
```

Проверяем, что в базе появился пользователь
```shell script
$ docker-compose exec postgres psql -U test -d exampledb -c "select * from users;"
 first_name |    email     | age
------------+--------------+-----
 otus       | otus@otus.ru |  27
(1 row)
```

Проверяем, что было опубликовано событие о новой регистрации
http://127.0.0.1:15672/#/queues/%2F/ToNotificationService
<img src="https://github.com/OtusGolang/webinars_practical_part/raw/master/31-integration-testing/assets/user_reg_event.png" width="600">

**Теперь у нас есть возможность писать тесты и дебажить их локально,
так как вся инфраструктура поднята в Docker, а необходимые порты пробросаны на host.**

Необходимо только помнить, куда тесты ходят - на localhost или во внутреннюю сеть докера.

Останавливаем docker-compose
```shell script
$ make down
docker-compose down
Stopping 31-integration-testing_registration_service_1 ... done
Stopping 31-integration-testing_notification_service_1 ... done
Stopping 31-integration-testing_rabbit_1               ... done
Stopping 31-integration-testing_postgres_1             ... done
Removing 31-integration-testing_registration_service_1 ... done
Removing 31-integration-testing_notification_service_1 ... done
Removing 31-integration-testing_rabbit_1               ... done
Removing 31-integration-testing_postgres_1             ... done
Removing network 31-integration-testing_db
Removing network 31-integration-testing_rabbit
```

### Интеграционное тестирование
После разработки и отладки тестов проверяем их работу из контейнера.

Запускаем тесты
```bash
$ make test
...
Creating network "31-integration-testing_db" with driver "bridge"
Creating network "31-integration-testing_rabbit" with driver "bridge"
Building notify_service
...
Successfully built c457d5c9c188
Successfully tagged 31-integration-testing_notify_service:latest
Building reg_service
...
Successfully built ee011c27e6be
Successfully tagged 31-integration-testing_reg_service:latest
Building integration_tests
...
Successfully built ab1eab529321
Successfully tagged 31-integration-testing_integration_tests:latest

Creating 31-integration-testing_postgres_1 ... done
Creating 31-integration-testing_rabbit_1   ... done
Creating 31-integration-testing_notify_service_1 ... done
Creating 31-integration-testing_reg_service_1    ... done
Creating 31-integration-testing_integration_tests_1 ... done

Starting 31-integration-testing_postgres_1 ... done
Starting 31-integration-testing_rabbit_1   ... done
Starting 31-integration-testing_notify_service_1 ... done
Starting 31-integration-testing_reg_service_1    ... done

2020/03/27 21:19:33 wait 5s for service availability...

...... 6

2 scenarios (2 passed)
6 steps (6 passed)
3.0521812s
testing: warning: no tests to run
PASS
ok      godog_example/integration_tests 8.096s

Stopping 31-integration-testing_reg_service_1    ... done
Stopping 31-integration-testing_notify_service_1 ... done
Stopping 31-integration-testing_rabbit_1         ... done
Stopping 31-integration-testing_postgres_1       ... done

Removing 31-integration-testing_integration_tests_run_a462f5aa65f2 ... done
Removing 31-integration-testing_integration_tests_1                ... done
Removing 31-integration-testing_reg_service_1                      ... done
Removing 31-integration-testing_notify_service_1                   ... done
Removing 31-integration-testing_rabbit_1                           ... done
Removing 31-integration-testing_postgres_1                         ... done
Removing network 31-integration-testing_db
Removing network 31-integration-testing_rabbit

$ echo $?
0
```

В логе мы видим:
- поднимаются микросервисы;
- успешно выполняются 2 тестовых сценария из 6 шагов;
- контейнеры останавливаются и удаляются;
- **сам скрипт возвращает 0 и 1 в зависимости от статуса прохождения тестов**
(это важно, так пригодится нам в Continuous Integration).
