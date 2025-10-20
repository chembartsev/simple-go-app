# simple-go-app

Предварительные требования
Развернута виртуальная машина Ubuntu с установленными Docker и Docker Compose

Шаги развертывания:
1. Клонировать репозиторий:
   git clone https://github.com/chembartsev/simple-go-app.git
   cd simple-go-app

2. Запустить все сервисы с помощью docker-compose
   docker-compose up -d --build

3. Проверить доступ: 
Приложение: http://your-vm-ip
Grafana: http://your-vm-ip:3000 (admin/admin123)
Prometheus: http://your-vm-ip:9090

4. Настроить дашборд в Grafana, добавив источником данных Prometheus

*******Мои действия при выполнении задания*******
1. Нашел простое приложение https://github.com/madlopt/simple-go-app/tree/main
2. Сделал fork в свой репозиторий 
3. Клонировал на виртуальную машину ubuntu
4. Изменил\создал файлы:
dockerfile 
docker-compose.yml 
main go 
prometheus.yml
5. Развернул на виртуальной машине с помощью docker-compose
6. Проверил доступность и работу
http://localhost/
http://localhost:3000
http://localhost/health
7. Сделал build с помощью action в github, создав перед этим файл ci.yml
8. В Grafana вручную создал простой дашборд, проверяющий работу приложения посредством Prometheus   

