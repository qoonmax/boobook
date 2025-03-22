set -e

source .env

# Проверка на наличие переменной окружения $DATABASE_PASSWORD и $REPLICATION_USER
if [ -z "$DATABASE_PASSWORD" ]; then
  echo "Ошибка: Переменная DATABASE_PASSWORD не установлена"
  exit 1
fi

if [ -z "$REPLICATION_USER" ]; then
  echo "Ошибка: Переменная REPLICATION_USER не установлена"
  exit 1
fi

if [ -z "$REPLICATION_PASSWORD" ]; then
  echo "Ошибка: Переменная REPLICATION_PASSWORD не установлена"
  exit 1
fi

# Устанавливаем переменную окружения для пароля
export PGPASSWORD=$DATABASE_PASSWORD

# Останавливаем контейнеры
docker-compose down postgres_master postgres_replica_1 postgres_replica_2

# Удаляем дамп данных
rm -rf ./database/data/master-dump
# Удаляем данные реплик
rm -rf ./database/data/replica_1
rm -rf ./database/data/replica_2

sleep 2

# Запускаем контейнер с мастером
docker-compose up -d postgres_master

# Проверяем готовность контейнера postgres_master
echo "Ожидание готовности postgres_master..."
until docker-compose exec -T postgres_master pg_isready -U postgres > /dev/null 2>&1; do
  sleep 2
done

# Создаем дамп данных
docker-compose exec -T postgres_master pg_basebackup -D /var/lib/postgresql/master-dump -U replicator -v -P --wal-method=stream

# Копируем данные из дампа
cp -r ./database/data/master-dump ./database/data/replica_1
cp -r ./database/data/master-dump ./database/data/replica_2

# Создаем файл standby.signal для реплики
touch ./database/data/replica_1/standby.signal
touch ./database/data/replica_2/standby.signal

# Изменяем конфигурацию реплики
if [[ "$OSTYPE" == "darwin"* ]]; then
  # Для macOS
  sed -i '' "s|^#primary_conninfo = .*|primary_conninfo = 'host=postgres_master port=5432 user=${REPLICATION_USER} password=${REPLICATION_PASSWORD} application_name=postgres_replica_1'|" ./database/data/replica_1/postgresql.conf
  sed -i '' "s|^#primary_conninfo = .*|primary_conninfo = 'host=postgres_master port=5432 user=${REPLICATION_USER} password=${REPLICATION_PASSWORD} application_name=postgres_replica_2'|" ./database/data/replica_2/postgresql.conf
else
  # Для Linux
  sed -i "s|^#primary_conninfo = .*|primary_conninfo = 'host=postgres_master port=5432 user=${REPLICATION_USER} password=${REPLICATION_PASSWORD} application_name=postgres_replica_1'|" ./database/data/replica_1/postgresql.conf
  sed -i "s|^#primary_conninfo = .*|primary_conninfo = 'host=postgres_master port=5432 user=${REPLICATION_USER} password=${REPLICATION_PASSWORD} application_name=postgres_replica_2'|" ./database/data/replica_2/postgresql.conf
fi

# Запускаем контейнер с репликой
docker-compose up -d postgres_replica_1 postgres_replica_2