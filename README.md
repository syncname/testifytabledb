

Статья <a href="">  </a>
[Статья на хабре](https://habr.com/ru/companies/first/articles/771428/)
Для запуска проекта нужен установленный докер (для контейнера с postgres).

Последовательность запуска

1) make postgres-container (postgres-run если контейнер уже существует)
2) make postgres-create
3) make migration-up
4) make test





