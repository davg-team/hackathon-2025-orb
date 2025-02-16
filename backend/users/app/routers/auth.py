# app/routers/auth.py
import random
import string
from datetime import timedelta

import requests
from app import crud, models, schemas
from app.auth import create_access_token, verify_password
from app.config import settings
from app.database import get_db
from fastapi import APIRouter, Depends, HTTPException, Request, Response, status
from fastapi.responses import RedirectResponse
from sqlalchemy.orm import Session

router = APIRouter(prefix="/auth", tags=["auth"])


def get_current_user(request: Request, db: Session = Depends(get_db)) -> models.User:
    """
    Извлекает текущего пользователя из cookie access_token.
    """
    token = request.cookies.get("access_token")
    if not token:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Пользователь не аутентифицирован",
        )
    from app.auth import decode_access_token

    payload = decode_access_token(token)
    if not payload:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED, detail="Неверный токен"
        )
    user = crud.get_user(db, payload.get("user_id"))
    if not user:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED, detail="Пользователь не найден"
        )
    return user


@router.post("/login", response_model=schemas.Token)
def login_for_access_token(
    response: Response, credentials: schemas.Login, db: Session = Depends(get_db)
):
    """
    Авторизация по логину и паролю. При успешном входе возвращается JWT-токен и устанавливается httpOnly cookie.
    """
    # Поиск пользователя сначала по email, затем по SNILS
    user = crud.get_user_by_email(db, credentials.username)
    if not user:
        user = crud.get_user_by_snils(db, credentials.username)
    if not user or not verify_password(credentials.password, user.password_hash):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Неверное имя пользователя или пароль",
        )

    token_data = {
        "user_id": user.id,
        "email": user.email,
        "role": user.role,
        "snils": user.snils,
        "full_name": user.full_name,
        "organization": user.organization,
        "position": user.position,
        "phone": user.phone,
    }
    access_token = create_access_token(
        data=token_data,
        expires_delta=timedelta(minutes=settings.ACCESS_TOKEN_EXPIRE_MINUTES),
    )
    response.set_cookie(key="access_token", value=access_token, httponly=False)
    response.headers["Location"] = "/"
    response.status_code = status.HTTP_302_FOUND
    return schemas.Token(access_token=access_token)


@router.get("/oauth/callback", response_model=schemas.Token)
def oauth_login(
    code: str, response: Response, reqests: Request, db: Session = Depends(get_db)
):
    """
    OAuth2 авторизация через внешний сервис. Обменивает code на access_token, получает данные пользователя и выдает JWT.
    """
    # Обмен кода на токен
    headers = {
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 YaBrowser/23.1.1.359 (beta) Yowser/2.5 Safari/537.36"
    }
    token_response = requests.post(
        settings.OAUTH_TOKEN_URL,
        data={
            "grant_type": "authorization_code",
            "code": code,
            "client_id": settings.OAUTH_CLIENT_ID,
            "client_secret": settings.OAUTH_CLIENT_SECRET,
            # "redirect_uri": "https://fsp-platform.ru/callback/auth/return/rsaag",
            "redirect_uri": "https://hackathon-8.orb.ru/auth/oauth/callback",
        },
        headers=headers,
    )

    if token_response.status_code != 200:
        print(token_response.text)
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Ошибка обмена OAuth токена",
        )

    token_json = token_response.json()
    print(token_json)
    oauth_access_token = token_json.get("access_token")
    if not oauth_access_token:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Не получен access_token от OAuth провайдера",
        )

    # Получение информации о пользователе
    user_info_response = requests.get(
        settings.OAUTH_USERINFO_URL,
        params={"scope": settings.OAUTH_SCOPES},
        headers={
            "Authorization": f"Bearer {oauth_access_token}",
            "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 YaBrowser/23.1.1.359 (beta) Yowser/2.5 Safari/537.36",
        },
    )
    print(user_info_response.text)
    if user_info_response.status_code != 200:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Не удалось получить информацию о пользователе от OAuth провайдера",
        )

    user_info = user_info_response.json()
    # Предполагается, что структура ответа: { "user": { "esia_snils": ..., "email": ..., "full_name": ..., ... } }
    user_data = user_info.get("user")
    if not user_data or "esia_snils" not in user_data:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Некорректные данные пользователя от OAuth провайдера",
        )

    # Если пользователь отсутствует в БД, создаём его
    print(user_data)
    user = crud.get_user_by_snils(db, user_data["esia_snils"])
    if not user:
        new_user = schemas.UserCreate(
            full_name=user_data.get("full_name", "Unknown"),
            organization=user_data.get("organization", "Unknown"),
            position=user_data.get("position", "Unknown"),
            phone=user_data.get("phone", None),
            email=user_data.get("email", user_data["esia_snils"]),  # запасной вариант
            snils=user_data["esia_snils"],
            role="user",
            password="".join(
                random.choices(string.ascii_letters + string.digits, k=12)
            ),
        )
        user = crud.create_user(db, new_user)

    token_data = {
        "user_id": user.id,
        "email": user.email,
        "role": user.role,
        "snils": user.snils,
        "full_name": user.full_name,
        "organization": user.organization,
        "position": user.position,
        "phone": user.phone,
    }
    jwt_token = create_access_token(
        data=token_data,
        expires_delta=timedelta(minutes=settings.ACCESS_TOKEN_EXPIRE_MINUTES),
    )
    response.set_cookie(key="access_token", value=jwt_token, httponly=False)
    # редирект на главную
    response.headers["Location"] = "/"
    response.status_code = status.HTTP_302_FOUND
    return schemas.Token(access_token=jwt_token)


@router.post("/logout")
def logout(response: Response):
    """
    Выход из системы. Удаляет httpOnly cookie с JWT-токеном.
    """
    response.delete_cookie(key="access_token")
    return {"detail": "Вы успешно вышли из системы"}


@router.get("/me", response_model=schemas.UserOut)
def read_users_me(current_user: models.User = Depends(get_current_user)):
    """
    Получение информации о текущем пользователе.
    """
    return current_user


@router.get("/login_as_root", response_model=schemas.Token)
def login_as_root(
    response: Response,
    code: str,
    db: Session = Depends(get_db),
):
    """
    Авторизация под другим пользователем (только с секретным кодом).
    """

    if code != settings.ROOT_LOGIN_CODE:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Неверный код авторизации",
        )

    user = crud.get_user_by_email(db, "admin")
    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Пользователь не найден",
        )

    token_data = {
        "user_id": user.id,
        "email": user.email,
        "role": user.role,
        "snils": user.snils,
        "full_name": user.full_name,
        "organization": user.organization,
        "position": user.position,
        "phone": user.phone,
    }

    access_token = create_access_token(
        data=token_data,
        expires_delta=timedelta(minutes=settings.ACCESS_TOKEN_EXPIRE_MINUTES),
    )

    response.set_cookie(key="access_token", value=access_token, httponly=False)
    return schemas.Token(access_token=access_token)
