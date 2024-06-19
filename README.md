<h1>Запуск:</h1>

<p>Задать имя хоста:</p>

```console
go run main.go host [HOST_NAME]
```
<p>Получить список DNS серверов:</p>

```console
go run main.go dns
```
<p>Добавить DNS:</p>

```console
go run main.go dns -a [DNS_IP]
```
<p>Удалить DNS:</p>

```console
go run main.go dns -d [DNS_IP]
```
Это клиент, сервер находится в репозитории по [ссылке](https://github.com/TaratinIvan/yadro-test-server)