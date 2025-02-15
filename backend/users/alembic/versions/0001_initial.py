"""initial migration: create users table and insert superadmin

Revision ID: 0001_initial
Revises:
Create Date: 2025-02-14 12:00:00.000000

"""

import sqlalchemy as sa
from alembic import op

# Импортируем passlib для хэширования пароля
from passlib.context import CryptContext
from sqlalchemy import DateTime, Integer, String, text

# revision identifiers, used by Alembic.
revision = "0001_initial"
down_revision = None
branch_labels = None
depends_on = None

# Создаём контекст для хэширования паролей
pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")


def upgrade():
    # Проверяем, существует ли таблица перед созданием
    conn = op.get_bind()
    result = conn.execute(text("SELECT to_regclass('public.users')"))
    table_exists = result.scalar() is not None

    if not table_exists:
        # Создаем таблицу users
        op.create_table(
            "users",
            sa.Column("id", sa.Integer, primary_key=True, autoincrement=True),
            sa.Column("full_name", sa.String(255), nullable=False),
            sa.Column("organization", sa.String(255), nullable=False),
            sa.Column("position", sa.String(255), nullable=False),
            sa.Column("phone", sa.String(50), nullable=True),
            sa.Column("email", sa.String(255), nullable=False, unique=True, index=True),
            sa.Column("snils", sa.String(50), nullable=False, unique=True, index=True),
            sa.Column("role", sa.String(50), nullable=False),
            sa.Column("password_hash", sa.String(255), nullable=False),
            sa.Column(
                "created_at",
                sa.DateTime(timezone=True),
                server_default=sa.func.now(),
            ),
            sa.Column("updated_at", sa.DateTime(timezone=True), onupdate=sa.func.now()),
        )

    # Проверяем, есть ли уже суперпользователь
    result = conn.execute(text("SELECT id FROM users WHERE email = 'admin'"))
    user_exists = result.scalar() is not None

    if not user_exists:
        # Вычисляем хэш для пароля "admin"
        password_hash = pwd_context.hash("admin")

        # Вставляем первую запись - суперадмина
        conn.execute(
            text(
                """
                INSERT INTO users (full_name, organization, position, phone, email, snils, role, password_hash)
                VALUES (:full_name, :organization, :position, :phone, :email, :snils, :role, :password_hash)
                """
            ),
            {
                "full_name": "Суперадмин",
                "organization": "Администрация",
                "position": "Суперадмин",
                "phone": "",
                "email": "admin",
                "snils": "000-000-000 00",
                "role": "root",
                "password_hash": password_hash,
            },
        )


def downgrade():
    op.drop_table("users")
