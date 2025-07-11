# 🤝 Contributing to Faculty Allocation Helper

Welocme here! Thank you for wanting to contribute to our project 🎉

---

## 📌 Common rules

- All changes must be linked to the **issue**
- All pull requests are **code reviewed**
- Any code must be **testable**, **consistent** in style and ** not violate** CI
- The issues are resolved by their assignees

## 📋 Kanban board

We use **Jira** as the main tools for planning. However, we have also got a board on GitLab

[🔗 Link to the board](https://gitlab.pg.innopolis.university/f.markin/fah/-/boards)

### Column policies

- **To Do**
  - The task is a priority and ready for implementation
  - All acceptance criteria are defined
  - Required subtasks or related issues added

- **In Progress**
  - The task has an assigned executor
  - The branch has been created, commits are linked to the issue

- **Done**
  - Merge Request accepted
  - Issue closed automatically via `Closes #<number>` in the PR description

## 🔁 Git Workflow

Мы используем **GitLab Flow**

### 📌 Main rules

- The issues are marked by the followwing lables: frontend, backend, UX/UI, devops
- Any issue can be assigned by any member at any moment

### 🏷️ Branching rules

- The branches are created and named after each microservice, like authorization, parsing or profile service. There are also separate branches for frontend and backend generally

### 💬 Commits

- Commits are made in free form by any member

### ✅ Pull Requests

- Pull requests can be created by any member for the purposes of requesting changing the functionality
- Any member's code can be reviewed by another member
- Any member can merge pull requests for other members

### 🔁 Gitgraph

See the diagram for the info

[gitgraph](/docs/development/git-workflow/gitgraph.png)

## 🔐 Secrets management

### Common rules

So far, the secrets are transmitted manually when configuring the server. They are not stored in the repository.


