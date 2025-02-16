# app/routers/users.py
from app import crud, models, schemas
from app.database import get_db
from fastapi import APIRouter, Depends, HTTPException, Request, status
from sqlalchemy.orm import Session

router = APIRouter(prefix="/users", tags=["users"])


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


def get_current_superadmin(current_user: models.User = Depends(get_current_user)):
    """
    Проверяет, что текущий пользователь имеет роль 'суперадмин'.
    """
    if current_user.role.lower() != "root":
        raise HTTPException(
            status_code=status.HTTP_403_FORBIDDEN, detail="Недостаточно прав"
        )
    return current_user


@router.get("/", response_model=list[schemas.UserOut])
def read_users(
    skip: int = 0,
    limit: int = 100,
    db: Session = Depends(get_db),
):
    """
    Получение списка всех пользователей (доступно только суперадмину).
    """
    users = crud.get_users(db, skip=skip, limit=limit)
    return users


@router.get("/{user_id}", response_model=schemas.UserOut)
def read_user(
    user_id: int,
    db: Session = Depends(get_db),
    current_user: models.User = Depends(get_current_superadmin),
):
    """
    Получение информации о конкретном пользователе по ID (суперадмин).
    """
    user = crud.get_user(db, user_id)
    if not user:
        raise HTTPException(status_code=404, detail="Пользователь не найден")
    return user


@router.post("/", response_model=schemas.UserOut)
def create_new_user(
    user: schemas.UserCreate,
    db: Session = Depends(get_db),
    current_user: models.User = Depends(get_current_superadmin),
):
    """
    Создание нового пользователя (суперадмин).
    """
    if crud.get_user_by_email(db, user.email):
        raise HTTPException(
            status_code=400, detail="Пользователь с таким email уже существует"
        )
    if crud.get_user_by_snils(db, user.snils):
        raise HTTPException(
            status_code=400, detail="Пользователь с таким СНИЛС уже существует"
        )
    return crud.create_user(db, user)


@router.put("/{user_id}", response_model=schemas.UserOut)
def update_existing_user(
    user_id: int,
    updates: schemas.UserUpdate,
    db: Session = Depends(get_db),
    current_user: models.User = Depends(get_current_superadmin),
):
    """
    Обновление данных пользователя (суперадмин).
    """
    user = crud.get_user(db, user_id)
    if not user:
        raise HTTPException(status_code=404, detail="Пользователь не найден")
    updated_user = crud.update_user(db, user, updates)
    return updated_user


@router.delete("/{user_id}")
def delete_existing_user(
    user_id: int,
    db: Session = Depends(get_db),
    current_user: models.User = Depends(get_current_superadmin),
):
    """
    Удаление пользователя (суперадмин).
    """
    user = crud.get_user(db, user_id)
    if not user:
        raise HTTPException(status_code=404, detail="Пользователь не найден")
    crud.delete_user(db, user)
    return {"detail": "Пользователь удалён"}


@router.post("/{user_id}/reset_password")
def reset_user_password(
    user_id: int,
    db: Session = Depends(get_db),
    current_user: models.User = Depends(get_current_superadmin),
):
    """
    Сброс пароля пользователя: генерируется новый случайный пароль и возвращается в ответе (с суперадминскими правами).
    """
    user = crud.get_user(db, user_id)
    if not user:
        raise HTTPException(status_code=404, detail="Пользователь не найден")
    new_password = crud.reset_password(db, user)
    # При необходимости можно добавить логирование сброса пароля
    return {"detail": "Пароль сброшен", "new_password": new_password}
