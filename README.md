# 📦 Vladislove MCP Server

**Jira MCP Server** — это сервер MCP (Model Context Protocol), предназначенный для работы с Jira API. Он предоставляет LLM-доступ к инструментам для анализа, учета и управления трудозатратами, проектами и пользователями. Интегрируется с Claude, Open WebUI и другими LLM-клиентами, поддерживающими MCP.
---

## 📂 Структура

- `cmd/main.go` — точка входа, запуск HTTP или STDIO сервера
- `internal/handlers` — реализация prompts, tools и resources
- `internal/registry` — регистрация всего функционала
- `internal/clients` — обертка над API
- `internal/types` — типы входных данных

---

## 📅 Установка

### 1. Склонировать

```bash
git clone https://github.com/vladislovely/jira-mcp-server.git
cd jira-mcp-server
```

### 2. Описать `.env`
```bash
md .env.example .env
```

```env
JIRA_API_URL=https://your-jira-api-url
JIRA_USER=your-jira-email
JIRA_TOKEN=your-jwt-token
OPENAI_API_KEY=your-openai-key
```

### 3. Запуск

```bash
make docker-build
```

### Можно добавить в любой клиент
Пример конфигурации
```json
{
    "mcpServers": {
        "mcp-server": {
            "command": "docker",
            "args": [
                "run",
                "--rm",
                "-i",
                "-e", "JIRA_TOKEN=your-token",
                "-e", "JIRA_USER=your-jira-email",
                "-e", "JIRA_API_URL=your-api-url",
                "mcp-server",
                "-t", "stdio"
            ]
        }
    }
}
```

## 🏠 Компоненты docker-compose

| Сервис          | Назначение                     |
| --------------- | ------------------------------ |
| `mcp-server`    | Основной MCP сервер            |
| `mcp-inspector` | UI-отладка prompts/tools       |

---
## 🌐 Запуск MCP

### STDIO:

```bash
go run cmd/main.go
```

### HTTP:

```bash
go run cmd/main.go -t http
```
## 👤 Автор

Владислав Зворыгин — [vladislav.zvorygin147@gmail.com](mailto\:vladislav.zvorygin147@gmail.com)

