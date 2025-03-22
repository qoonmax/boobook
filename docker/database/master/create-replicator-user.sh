#!/bin/bash
set -e

echo "Создание пользователя репликатора..."

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<EOF
CREATE ROLE replicator WITH LOGIN REPLICATION PASSWORD 'replicator';
EOF

echo "Пользователь репликатора успешно создан."
