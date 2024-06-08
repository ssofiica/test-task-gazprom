# Тестовое задание

ER-диаграмма
```mermaid
erDiagram
    BIRTHDAY_SUBSCRIBING||--o{ USER : subscribes
    BIRTHDAY_SUBSCRIBING||--o{ USER : has-birthday
    BIRTHDAY_SUBSCRIBING {
        int birthday_user_id PK, FK
        int subscribing_user_id PK, FK
    } 
    USER{
        int id PK
        text name
        text surname
        text email
        text password
        date birthday
    }
```

