<div align="center">
  
# RTU IT Lab <br> Требования по направлению Backend

<img height="300" alt="logo" src="https://avatars.githubusercontent.com/u/38423762?s=200&v=4">

</div> 


## Оглавление
- ### [Задание](#1)
- ### [Решение](#2)
- ### [Запуск кода](#3)
- ### [Стек](#4)
- ### [Контакты](#5)
- ### [Ссылки](#6)

## <a name="1"> Задание </a>

Вы начальник ГИБДД Гиперборейского городского округа. Только что завершилась очередная русо-ящерская война. В этой войне, к сожалению, русы потеряли ценнейшие технологии машинного обучения, которые в 21 веке стали называться якобы технологиями будущего. Для обеспечения безопасности на дорогах вами было принято решение о создании системы автоматической фиксации нарушений правил дорожного движения, однако работу нейронных сетей будут выполнять пленные ящеры - проверяющие специалисты или оценщики.

## <a name="2">Решение </a>

### Архетиктура решения

<div align="center">
<img height="300" alt="logo" src="DemoCameraClientPy\photo\lab.png">
</div> 
### Идея решения

Данную задачу я решил сделать через REST API, где ключевыми технологиями были СУБД PostgreSQL и Golang(без фреймворков). Я уважаю MVC подход, но в 'best practices' по Golang использовалось немного видоизмененная архитектура, поэтому  у меня получилось что-то среднее. Также, чтобы меньше гонять данные между API и БД, было принято решение написать большую часть логики в самой СУБД на PLpgSQL. Поэтому на Golang отсалась задача аутентификации валидирования данных и формирования запросов в telegram API и в БД. Что, я считаю положительно повлияло на стабильность и скорость.





## <a name="3">Запуск кода </a>

### Последовательные шаги для запуска кода:
1. Склонируйте гит репозиторий;
```Bash
git clone https://github.com/Fact0RR/RTULab.git
```

2. Запуск контейнеров и сервера стриминга:

```Bash
cd RTULab
docker-compose build
docker-compose up
```

3. Дальше требуется подключиться к контейнеру под названием main_server и обновиться до последней версии миграции


```Bash
goose -dir migration postgres "host=db dbname=test_db sslmode=disable user=root password=ваш_пароль" up
```

## <a name="4">Стек </a>
<div align="center">

  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/8/83/Telegram_2019_Logo.svg/2048px-Telegram_2019_Logo.svg.png" title="Telegram" alt="Telegram" height="40"/>&nbsp;
  <img src="https://ezerus.com.au/wp-content/uploads/2019/05/postgresql-logo.png" title="PostgreSQL" alt="PostgreSQL" height="40"/>&nbsp;
  <img src="https://static-00.iconduck.com/assets.00/docker-icon-512x438-ga1hb37h.png" title="Docker" alt="Docker" height="40"/>&nbsp;

  <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaUm2uk_PwKIKZnQ-SqRyAxyFz7vWJPEgULg&usqp=CAU" alt="POSTMAN"  height="40"/>&nbsp;
  <img src="https://fronty.com/static/uploads/1.11-30.11/languages%20in%202022/go.png" title="GO" alt="GO" height="40"/>&nbsp;
  <img src="https://cdn.freebiesupply.com/logos/thumbs/2x/visual-studio-code-logo.png"  title="VScode" alt="VScode" height="40"/>
</div>

## <a name="5">Контакты</a>

- <h4><img align="center" height="25" src="https://user-images.githubusercontent.com/51875349/198863127-837491f2-b57f-4c75-9840-6a4b01236c7a.png">: @facT0RR, Маслов Денис - Fullstack-developer</h3>

## <a name="6">Ссылки </a>

- [ссылка на скринкаст]()&nbsp;
- [ссылка на swagger](https://github.com/Fact0RR/RTULab/blob/master/openapi.yaml)&nbsp;
