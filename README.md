# Boobook - социальная сеть, домашняя работа курса Otus Highload Architecture

![Image alt](https://github.com/qoonmax/boobook/blob/main/cover.jpg)

### Локальная установка
1. Склонировать репозиторий;
2. Выполнить ``cp .env.example .env``;
3. Запустить ``make to-dev``;
4. Запустить ``make run``;
5. Запустить ``/bin/bash /home/qoonmax/GolandProjects/boobook/docker/database/replica/sync-data.sh`` 
чтобы синхронизировать данные между мастером и репликой;

### TODO список доработок:
1. [x] Инжектить логгер в контроллеры, убрать глобальный;
2. [ ] В сервисы передавать DTO, или примитивы, а не объекты реквестов;
3. [ ] Добавить ограничение времени выполнение запроса;