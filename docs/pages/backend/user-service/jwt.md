# **Работа с JWT Токенами**

## **Что такое JWT?**

JWT (JSON Web Token) — это компактный, URL-безопасный способ представления претензий (claims) между двумя сторонами, такими как клиент и сервер. В контексте микросервиса JWT токен используется для авторизации пользователей. Токен подписывается секретным ключом и может быть проверен при каждом запросе к серверу.

JWT токен состоит из трех частей:

1. **Header (Заголовок)**: Содержит информацию о типе токена (обычно это `JWT`) и алгоритме подписи (например, HMAC SHA256 или RSA).
2. **Payload (Полезная нагрузка)**: Содержит утверждения или *claims* — информацию о пользователе и метаданные, такие как идентификатор пользователя, его роль, время выпуска и время истечения токена.
3. **Signature (Подпись)**: Для создания подписи используется секретный ключ или приватный ключ (в случае асимметричной криптографии). Подпись необходима для проверки подлинности токена.

**Пример структуры JWT:**

```txt
Header.Payload.Signature
```

---

## **Пример содержимого JWT токена**

Пример закодированного токена с полезной нагрузкой:

```json
{
  "aud": "platform",
  "exp": 1733552183.4535291,
  "fsp_id": "0",
  "iat": 1733508983.45353,
  "iss": "platform",
  "role": "superadmin",
  "sub": "01939d2a-c96d-7f54-8a89-5a0faf6119ec"
}

```

### **Описание полей:**

- **sub (subject)** — Идентификатор пользователя, уникальный для каждого пользователя.
- **role** — Роль пользователя в системе (например, `admin`, `user`, `superadmin`).
- **iat (issued at)** — Время, когда токен был выдан.
- **exp (expiration time)** — Время истечения срока действия токена.
- **aud (audience)** — Для кого предназначен токен, может быть использован для ограничения доступа.
- **fsp_id** — Дополнительный идентификатор, если необходимо использовать в системе, может быть специфичен для конкретного приложения.
- **iss (issuer)** — Идентификатор, указывающий на систему, которая выдала токен.

