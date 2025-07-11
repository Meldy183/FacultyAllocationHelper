## 🏗 Architecture

Architecture-related content is organized under [`docs/architecture/`](docs/architecture/). All diagrams are available both as images and editable [PlantUML](https://plantuml.com/) sources.

---

### 🧱 Static View

- 📌 **UML Component Diagram**  
  ![Static View Diagram](docs/architecture/static-view/ComponentUML.png)

- 📝 [View the PlantUML source](docs/architecture/static-view/StaticView.puml)

Our system follows a **modular monolith** architecture:

- The **backend** is a single Go application that handles all business logic and database interaction.
- The **frontend** is a separate Next.js service.
- **External authentication** is provided via a third-party SSO system (e.g., via OAuth).
- The backend exposes HTTP endpoints consumed by the frontend. There is no API Gateway.
- The backend also communicates directly with the **PostgreSQL** database.

🔁 This architecture ensures:
- **High cohesion** within services
- **Loose coupling** between frontend and backend
- **Improved maintainability** due to a clear division of concerns

---

### 🔄 Dynamic View

- 📌 **UML Sequence Diagram**  
  ![Dynamic View Diagram](docs/architecture/dynamic-view/DynamicView_SequenceDiagram.png)

### 🚀 Deployment View

- 📌 **UML Deployment Diagram**  
  ![Deployment View Diagram](docs/architecture/deployment-view/DeploymentUML.png)

- 📝 [View the PlantUML source](docs/architecture/deployment-view/DeploymentUML.puml)

The application is deployed on a remote server and uses **Docker containers** for isolation and scalability.

**Deployment structure:**
- `frontend`: Container for the Next.js web client
- `backend`: Container for the Go API
- `db`: Container for PostgreSQL

Each container communicates over the local Docker network. Only the frontend is exposed externally to users. Backend is accessed only through frontend or authenticated channels.

🧳 On the customer’s side, the system can be deployed via Docker Compose. No external cloud services or Kubernetes are required, making the setup lightweight and reproducible.

---

### ⚙️ Tech Stack

| Layer        | Technology            |
|--------------|------------------------|
| Frontend     | Next.js (React)        |
| Backend      | Go                     |
| Auth         | External SSO (OAuth)   |
| DB           | PostgreSQL             |
| DevOps       | Docker, GitLab CI      |
| Testing      | Jest (frontend), Go `testing` (backend) |

---