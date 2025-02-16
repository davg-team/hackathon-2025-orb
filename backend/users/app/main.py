# app/main.py
from app import models
from app.database import engine
from app.routers import auth as auth_router
from app.routers import users as users_router
from fastapi import FastAPI

# Для разработки можно создавать таблицы автоматически.
# В продакшене рекомендуется использовать Alembic для миграций.
models.Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="Микросервис авторизации и управления пользователями",
    description="API для аутентификации, выдачи JWT-токенов, OAuth2-интеграции и управления пользователями",
    version="1.0.0",
)

app.include_router(auth_router.router)
app.include_router(users_router.router)
