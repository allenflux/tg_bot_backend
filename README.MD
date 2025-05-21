
# 🤖 tg_bot_backend

`tg_bot_backend` is an open-source backend service built with Go using the GoFrame (gf) framework. It is designed to manage Telegram bots that help automate task assignments and permission management in group chats.

---

## 🌍 Features

- Assign tags to Telegram groups via bot commands
- Authorize specific users for task-related operations
- Interact with external services to retrieve group member info
- Secure and modular backend service with configuration flexibility

---

## 📦 Tech Stack

- **Language**: Go (Golang)
- **Framework**: GoFrame (gf)
- **Database**: MySQL
- **Cache**: Redis
- **External Service**: Python-based API (`tg_bot_assistant`)

---

## 🛠️ Configuration

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

### 🏃 Running the Service

To run the service in development mode:

```bash
gf run main.go
```

To build a production-ready binary:

```bash
gf build
```

Once built, you can run the compiled binary:

```bash
./tg_bot_backend
```

Make sure your `config.yaml` is in the correct location or passed via environment/config flags.

---

## 📖 Usage

Once the bot has been added to a Telegram group and is properly authorized, it supports the following commands:

- `/taggroup <tag>`  
  Tags the current group with a label.

- `/adduser <username>`  
  Grants permission to a user for executing certain tasks in the group.

- `/task <action>`  
  Executes a predefined task associated with the current group.

All commands are restricted to authorized users as configured in the backend logic.

---

## 🔐 Security

The project includes several built-in security features:

- Role-based access control for command usage
- Telegram user ID verification
- Secure integration with `tg_bot_assistant` via a trusted IP and endpoint
- Input validation and protection using the GoFrame framework
- Future support planned for JWT-based admin API authentication

---

## 🌍 Internationalization (i18n)

Internationalization support is planned in upcoming releases. The bot will detect user or group locale and provide localized responses accordingly.

---

## 🔗 Dependencies

This project depends on an external Python service:  
[**tg_bot_assistant**](https://github.com/allenflux/tg_bot_assistant)

This service is responsible for fetching authorized group members using the following API endpoint:

```text
http://8.210.149.123:5000/get_members?group_link=<your_group_link>
```

Make sure this service is running and accessible before using any features involving group membership validation.

---

## 🤝 Contributing

We welcome contributions! To get started:

```bash
# 1. Fork the repository
# 2. Create your feature branch
git checkout -b feature/your-feature-name

# 3. Make your changes and commit
git commit -am 'Add new feature'

# 4. Push to your fork
git push origin feature/your-feature-name

# 5. Open a Pull Request on GitHub
```

Please follow our coding guidelines and keep commits clear and well-documented.

---

## 🧭 Roadmap

Planned features and improvements:

- [ ] Admin dashboard (web UI)
- [ ] Multi-language (i18n) support
- [ ] Plugin system for dynamic task extensions
- [ ] Enhanced logging and observability
- [ ] Group-level analytics and metrics

---

## 📄 License

This project is released under the [MIT License](./LICENSE).

---

## 📬 Contact

Have feedback or want to collaborate?

- Open an issue on [GitHub](https://github.com/allenflux/tg_bot_backend/issues)
- Or reach out via email or Telegram (to be added)
