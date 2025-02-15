# app/config.py
import os

from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    DATABASE_URL: str = os.getenv("DATABASE_URL")
    SECRET_KEY_PATH: str = "./keys/private.pem"
    PUBLIC_KEY_PATH: str = "./keys/public.pem"
    ALGORITHM: str = "RS256"
    ACCESS_TOKEN_EXPIRE_MINUTES: int = 60  # время жизни токена в минутах
    OAUTH_AUTHORIZE_URL: str = "https://lk.orb.ru/oauth/authorize"
    OAUTH_TOKEN_URL: str = "https://lk.orb.ru/oauth/token"
    OAUTH_USERINFO_URL: str = "https://lk.orb.ru/api/get_user"
    OAUTH_SCOPES: str = "rsaag_id+email"
    OAUTH_CLIENT_ID: str
    OAUTH_CLIENT_SECRET: str

    class Config:
        env_file = ".env"


settings = Settings()
