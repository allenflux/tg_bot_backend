# tg_bot_backend

**tg_bot_backend** is an open-source backend service built with the [GoFrame (GF)](https://goframe.org/) framework. It is designed to manage Telegram bots with powerful group tagging and task delegation capabilities. This backend interacts with Telegram bots to support group organization and user-level task management within each group.

## 🌍 Overview

This project targets developers and community managers seeking to automate operational tasks inside Telegram groups through a structured bot-driven approach. With `tg_bot_backend`, you can:

- Tag and categorize Telegram groups for better organization.
- Grant authorized users permission to execute tasks in specific groups.
- Integrate with external tools or scripts to streamline group workflows.

It’s optimized for performance, scalability, and ease of integration into larger bot-based systems.

## 🚀 Features

- Group tagging and identification via Telegram bot commands
- Role-based access and task delegation for group members
- Fast MySQL-based persistent storage and Redis caching
- Easily extendable API structure for new task types
- Seamless integration with external Python services

## 📦 Tech Stack

- **Language**: Go (with GoFrame framework)
- **Database**: MySQL
- **Cache**: Redis
- **Integration**: Telegram Bot API
- **Deployment**: Docker-compatible

## 🔗 Dependencies

This project depends on the external Python project [**tg_bot_assistant**](https://github.com/yourusername/tg_bot_assistant), which provides Telegram group member data via a RESTful API.

> 🔗 API endpoint used:
> http://8.210.149.123:5000/get_members?group_link=<your_group_link>


Ensure this service is running and reachable before using features related to group member queries.

## 🛠️ Getting Started

### Prerequisites

- Go 1.18+ and [GoFrame CLI](https://goframe.org/pages/viewpage.action?pageId=1114315)
- MySQL database instance
- Redis instance
- A valid Telegram Bot Token

### Installation

```bash
git clone https://github.com/allenflux/tg_bot_backend.git
cd tg_bot_backend
gf run main.go
```  

### Configuration

Edit the configuration file `config/config.yaml` with the following content:

```yaml
server:
  address: ":8822"
  dumpRouterMap: true
  routeOverWrite: true

database:
  default:
    link: "mysql:root:123456@tcp(127.0.0.1:3306)/tg_bot?charset=utf8mb4&parseTime=true&loc=Local"
    debug: true

redis:
  default:
    address: 127.0.0.1:6379

tg_bot_assistant:
  address: http://8.210.149.123:5000/get_members?group_link=
```