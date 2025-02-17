# **Архитектура и технологии**

## **Архитектура**

Микросервис спроектирован на основе принципов модульности и низкой связанности, чтобы обеспечить масштабируемость, простоту тестирования и удобство разработки. Основными архитектурными элементами являются:

### **1. Паттерн Repository**

Паттерн *Repository* используется для отделения бизнес-логики от логики доступа к данным, создавая абстрактный уровень взаимодействия с хранилищем данных.

**Преимущества использования Repository:**

- **Абстракция доступа к данным**: Бизнес-логика остается независимой от способа хранения данных.
- **Удобство тестирования**: Позволяет заменить репозитории моками, упрощая тестирование.
- **Упрощение кода**: Вся логика работы с данными сосредоточена в репозиториях, что облегчает поддержку и расширение.

**Пример реализации репозитория:**

```python
class UserRepository(BaseRepository):
    async def get_by_id(self, user_id: int) -> User:
        # Логика получения пользователя из базы данных
        return await self.session.query(User).filter(User.id == user_id).first()

```

**Использование репозитория в сервисе:**

```python
class UserService:
    def __init__(self, uow: UoW):
        self.repository = uow.user_repo

    async def get_user_by_id(self, user_id: int):
        return await self.repository.get_by_id(user_id)

```

---

### **2. Сервисы**

Сервисы реализуют бизнес-логику микросервиса и используют репозитории для взаимодействия с базой данных. Они предоставляют методы, которые непосредственно решают задачи приложения, инкапсулируя всю логику в одном месте.

### **Примеры сервисов**:

1. **UserService**: Управление пользователями.
2. **RelationService**: Работа с отношениями пользователей и провайдеров.
3. **NotificationService**: Отправка и обработка уведомлений.
4. **RequestService**: Управление запросами на изменение ролей или данных.

**Пример реализации сервиса:**

```python
class UserService:
    def __init__(self, uow: UoW):
        self.repository = uow.user_repo

    async def create_user(self, user_data: dict) -> User:
        user = User(**user_data)
        await self.repository.add(user)
        return user

```

**Задачи сервисов:**

- Инкапсуляция сложной бизнес-логики.
- Валидация данных перед их обработкой.
- Обработка исключений, связанных с бизнес-правилами.

---

### **3. Unit of Work (UoW)**

Паттерн *Unit of Work* (UoW) отвечает за управление транзакциями и отслеживание изменений объектов, обеспечивая атомарность операций.

**Зачем нужен UoW:**

- **Управление транзакциями**: Позволяет объединять несколько операций в одну транзакцию.
- **Согласованность данных**: Гарантирует, что все операции либо завершатся успешно, либо будут отменены.
- **Удобство использования**: Упрощает написание кода для обработки нескольких связанных действий.

**Пример использования UoW:**

```python
async with await uow() as transaction:
    user_service = UserService(transaction)
    relation_service = RelationService(transaction)

    user = await user_service.create_user(user_data)
    await relation_service.create_relation(user.id, provider_data)

    await transaction.commit()

```
