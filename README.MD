<div align="center">
  
# RTU IT Lab <br> Требования по направлению Backend

<img height="300" alt="logo" src="https://avatars.githubusercontent.com/u/38423762?s=200&v=4">

</div> 


## Оглавление
- ### [Задание](#1)
- ### [Решение](#2)
- ### [Запуск кода](#3)
- ### [Уникальность нашего решения](#4)
- ### [Стек](#5)
- ### [Команда](#6)
- ### [Ссылки](#7)

## <a name="1"> Задание </a>

Вы начальник ГИБДД Гиперборейского городского округа. Только что завершилась очередная русо-ящерская война. В этой войне, к сожалению, русы потеряли ценнейшие технологии машинного обучения, которые в 21 веке стали называться якобы технологиями будущего. Для обеспечения безопасности на дорогах вами было принято решение о создании системы автоматической фиксации нарушений правил дорожного движения, однако работу нейронных сетей будут выполнять пленные ящеры - проверяющие специалисты или оценщики.

## <a name="2">Решение </a>

### Архетиктура решения

<div align="center">
<img height="300" alt="logo" src="assets/deploy.drawio.png">

**Развертывание модели**
</div> 



## <a name="3">Запуск кода </a>

### Последовательные шаги для запуска кода:
1. Склонируйте гит репозиторий;
```Bash
git clone https://github.com/BuldakovN/RRR-Video-Action-Classification.git
```
2. Скачайте веса для модели детекции https://drive.google.com/drive/folders/1iQRcneyDnb3q7LZ9wcGfkDt3W5zAKjwD?usp=sharing и разместите их в папках ```model/weights``` и ```modelStream/weights```

3. Запуск контейнеров и сервера стриминга:
```Bash
cd RRR-Video-Action-Classification
docker-compose build
docker-compose up

cd modelStream

pip3 install -r requirements.txt
python3 serverStream.py
```

## <a name="4">Уникальность нашего решения </a>

1. Возможность распознавания действий нескольких людей на одном кадре.

2. Работает в режиме реального времени.

3. Модель была предобучена на большом датасете.

4. Удобный веб-интерфейс для работы с моделью.

## <a name="5">Стек </a>
<div align="center">
  
  <img src="https://ezerus.com.au/wp-content/uploads/2019/05/postgresql-logo.png" title="PostgreSQL" alt="PostgreSQL" height="40"/>&nbsp;
  <img src="https://static-00.iconduck.com/assets.00/docker-icon-512x438-ga1hb37h.png" title="Docker" alt="Docker" height="40"/>&nbsp;

  <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaUm2uk_PwKIKZnQ-SqRyAxyFz7vWJPEgULg&usqp=CAU" alt="POSTMAN"  height="40"/>&nbsp;
  <img src="https://fronty.com/static/uploads/1.11-30.11/languages%20in%202022/go.png" title="GO" alt="GO" height="40"/>&nbsp;
  <img src="https://cdn.freebiesupply.com/logos/thumbs/2x/visual-studio-code-logo.png"  title="VScode" alt="VScode" height="40"/>
</div>

## <a name="7">Ссылки </a>

- [ссылка на веса модели детекции](https://drive.google.com/drive/folders/1iQRcneyDnb3q7LZ9wcGfkDt3W5zAKjwD?usp=sharing)&nbsp;
- [ссылка на скринкаст](https://drive.google.com/file/d/1Md4uNQFHtO_w9xxaigyc6ftbF0Tru6qk/view?usp=sharing)&nbsp;
