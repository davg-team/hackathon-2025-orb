# app/crud.py
import random
import string

from app import models, schemas
from app.auth import get_password_hash
from sqlalchemy.orm import Session


def get_user_by_email(db: Session, email: str):
    return db.query(models.User).filter(models.User.email == email).first()


def get_user_by_snils(db: Session, snils: str):
    return db.query(models.User).filter(models.User.snils == snils).first()


def get_user(db: Session, user_id: int):
    return db.query(models.User).filter(models.User.id == user_id).first()


def get_users(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.User).offset(skip).limit(limit).all()


def create_user(db: Session, user: schemas.UserCreate):
    hashed_password = get_password_hash(user.password)
    db_user = models.User(
        full_name=user.full_name,
        organization=user.organization,
        position=user.position,
        phone=user.phone,
        email=user.email,
        snils=user.snils,
        role=user.role,
        password_hash=hashed_password,
    )
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user


def update_user(db: Session, db_user: models.User, updates: schemas.UserUpdate):
    if updates.full_name is not None:
        db_user.full_name = updates.full_name
    if updates.organization is not None:
        db_user.organization = updates.organization
    if updates.position is not None:
        db_user.position = updates.position
    if updates.phone is not None:
        db_user.phone = updates.phone
    if updates.email is not None:
        db_user.email = updates.email
    if updates.snils is not None:
        db_user.snils = updates.snils
    if updates.role is not None:
        db_user.role = updates.role
    if updates.password is not None:
        db_user.password_hash = get_password_hash(updates.password)
    db.commit()
    db.refresh(db_user)
    return db_user


def delete_user(db: Session, db_user: models.User):
    db.delete(db_user)
    db.commit()


def reset_password(db: Session, db_user: models.User, length: int = 12):
    """
    Сбрасывает пароль пользователя, генерируя новый случайный пароль.
    """
    new_password = "".join(
        random.choices(string.ascii_letters + string.digits, k=length)
    )
    db_user.password_hash = get_password_hash(new_password)
    db.commit()
    return new_password
