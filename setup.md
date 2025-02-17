### 🛠 **Пошаговая инструкция по развертыванию проекта "Книга Памяти Оренбургской области"**  

---

## 🔹 **1. Подготовка окружения**  

📌 **Требования:**  

- Linux / Windows / macOS  
- Установленные:  
  - **Git**  
  - **Docker**  
  - **Docker Compose**  

### **1.1 Установка Docker и Docker Compose**  

📌 **На Ubuntu/Debian:**  

```bash
sudo apt update && sudo apt install -y docker.io docker-compose
sudo systemctl enable docker --now
```

📌 **На macOS:**  

1. Установите [Docker Desktop](https://www.docker.com/products/docker-desktop/)  
2. После установки перезапустите систему.  

📌 **На Windows:**  

1. Установите [Docker Desktop](https://www.docker.com/products/docker-desktop/)  
2. Включите **WSL 2** в настройках Docker.  
3. Перезагрузите систему.  

---

## 🔹 **2. Клонирование репозитория**  

📌 Выполните команду:  

```bash
git clone https://github.com/davg-team/hackathon-2025-orb.git
cd hackathon-2025-orb
```

---

<!-- ## 🔹 **3. Аутентификация в Yandex Container Registry**  

Так как образы загружены в Yandex Container Registry, нужно выполнить вход:  

```bash
docker login cr.yandex
```

**Введите логин и пароль (или используйте токен).**  

--- -->

## 🔹 **3. Запуск проекта**  

📌 **Запуск всех контейнеров:**  

```bash
docker-compose up -d --build
```

⚡ Все сервисы запустятся в фоне.  

📌 **Проверка запущенных контейнеров:**  

```bash
docker-compose ps
```

📌 **Просмотр логов:**  

```bash
docker-compose logs -f
```

⏳ **Дождитесь полной загрузки сервисов.**

---


## 🔹 **Остановка проекта**  

📌 Остановка всех контейнеров:  

```bash
docker-compose down
```

📌 Полное удаление контейнеров и томов:  

```bash
docker-compose down -v
```
