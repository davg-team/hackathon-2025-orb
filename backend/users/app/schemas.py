# app/schemas.py
from typing import Optional

from pydantic import BaseModel


class UserBase(BaseModel):
    full_name: str
    organization: str
    position: str
    phone: Optional[str] = None
    email: str
    snils: str
    role: str  # допустимые значения: root, admin, user


class UserCreate(UserBase):
    password: str


class UserUpdate(BaseModel):
    full_name: Optional[str]
    organization: Optional[str]
    position: Optional[str]
    phone: Optional[str]
    email: Optional[str]
    snils: Optional[str]
    role: Optional[str]
    password: Optional[str]


class UserOut(UserBase):
    id: int

    class Config:
        orm_mode = True


class Token(BaseModel):
    access_token: str
    token_type: str = "bearer"


class TokenData(BaseModel):
    user_id: int
    email: str
    role: str


class Login(BaseModel):
    username: str  # может быть email или snils
    password: str
