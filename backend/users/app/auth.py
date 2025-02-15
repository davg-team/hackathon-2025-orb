# app/auth.py
from datetime import datetime, timedelta

from app.config import settings
from app.schemas import TokenData
from jose import JWTError, jwt
from passlib.context import CryptContext

# Создаём контекст для хэширования паролей
pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")


def verify_password(plain_password: str, hashed_password: str) -> bool:
    """
    Проверяет, соответствует ли введённый пароль хэшированному паролю.
    """
    return pwd_context.verify(plain_password, hashed_password)


def get_password_hash(password: str) -> str:
    """
    Хэширует пароль с помощью bcrypt.
    """
    return pwd_context.hash(password)


def create_access_token(data: dict, expires_delta: timedelta = None):
    """
    Генерирует JWT-токен с указанными данными, подписанный RSA-ключом.
    """
    to_encode = data.copy()
    expire = datetime.utcnow() + (
        expires_delta
        if expires_delta
        else timedelta(minutes=settings.ACCESS_TOKEN_EXPIRE_MINUTES)
    )
    to_encode.update({"exp": expire})

    # Чтение приватного RSA-ключа
    with open(settings.SECRET_KEY_PATH, "r") as key_file:
        private_key = key_file.read()

    encoded_jwt = jwt.encode(to_encode, private_key, algorithm=settings.ALGORITHM)
    return encoded_jwt


def decode_access_token(token: str):
    """
    Декодирует и проверяет JWT-токен, используя публичный RSA-ключ.
    """
    try:
        with open(settings.PUBLIC_KEY_PATH, "r") as key_file:
            public_key = key_file.read()
        payload = jwt.decode(token, public_key, algorithms=[settings.ALGORITHM])
        return payload
    except JWTError:
        return None
