#!/bin/sh
set -e

# Применяем миграции
echo "Применяем миграции Alembic..."
alembic upgrade head

# Запускаем приложение
echo "Запуск uvicorn..."
exec uvicorn app.main:app --host 0.0.0.0 --port 8000
