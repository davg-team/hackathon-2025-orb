# app/models.py
from sqlalchemy import Column, DateTime, Integer, String, func
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True, index=True)
    full_name = Column(String(255), nullable=False)  # ФИО
    organization = Column(String(255), nullable=False)  # Организация
    position = Column(String(255), nullable=False)  # Должность
    phone = Column(String(50), nullable=True)  # Номер телефона
    email = Column(
        String(255), unique=True, nullable=False, index=True
    )  # Электронная почта
    snils = Column(String(50), unique=True, nullable=False, index=True)  # СНИЛС
    role = Column(String(50), nullable=False)  # Роль: root, admin, user
    password_hash = Column(String(255), nullable=False)  # Хэш пароля
    created_at = Column(DateTime(timezone=True), server_default=func.now())
    updated_at = Column(DateTime(timezone=True), onupdate=func.now())
