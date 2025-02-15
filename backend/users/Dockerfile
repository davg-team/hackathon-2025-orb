# Используем официальный образ Python
FROM python:3.13-slim

# # Устанавливаем рабочую директорию внутри контейнера
# WORKDIR /app

# Копируем файл зависимостей и устанавливаем их
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Копируем исходный код проекта
COPY . .

# Открываем порт 8000
EXPOSE 8000

# Запускаем entrypoint скрипт
ENTRYPOINT ["/entrypoint.sh"]