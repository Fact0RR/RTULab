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

### Архетиктура модели

<div align="center"><img height="200" alt="модель" src="assets/model_white.png"></div>

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
  <img src="https://github.com/devicons/devicon/blob/master/icons/python/python-original-wordmark.svg" title="Python" alt="Python" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/pytorch/pytorch-original.svg" title="Pytorch" alt="Pytorch" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/numpy/numpy-original.svg" title="Numpy" alt="Numpy" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/opencv/opencv-original.svg" title="OpenCV" alt="OpenCV" height="40"/>&nbsp;
  
  <img src="https://pjreddie.com/media/image/yologo_2.png" title="Yolo" alt="Yolo" height="40"/>&nbsp;
  <img src="https://spacenil.com/tutorial/public/uploads/categories/categories_1599665107.png" title="JS" alt="JS"  height="40"/>&nbsp;
  <img src="https://fronty.com/static/uploads/1.11-30.11/languages%20in%202022/go.png" title="GO" alt="GO" height="40"/>&nbsp;
  <img src="https://upload.wikimedia.org/wikipedia/commons/3/3c/Flask_logo.svg"  title="Flask" alt="Flask" height="40"/>
</div>

## <a name="7">Ссылки </a>

- [ссылка на веса модели детекции](https://drive.google.com/drive/folders/1iQRcneyDnb3q7LZ9wcGfkDt3W5zAKjwD?usp=sharing)&nbsp;
- [ссылка на скринкаст](https://drive.google.com/file/d/1Md4uNQFHtO_w9xxaigyc6ftbF0Tru6qk/view?usp=sharing)&nbsp;
