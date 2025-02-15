# app/database.py
from app.config import settings
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

engine = create_engine(settings.DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)


def get_db():
    """
    Зависимость для получения сессии базы данных.
    """
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
