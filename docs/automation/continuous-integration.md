# 🔄 Continuous Integration (CI)

This document describes the Continuous Integration setup used in the **Faculty Allocation Helper (FAH)** project. The CI pipeline automates key development processes to improve code quality, ensure consistency, and catch issues early.

---

## 📂 CI Workflow File Location

The main GitLab CI configuration is stored in the following file:

[`fah-frontend/.gitlab-ci.yml`](fah-frontend/.gitlab-ci.yml)

This file defines the CI jobs triggered on every push or merge request.

---

## 🛠️ What’s Included in the CI Pipeline

Currently, the pipeline includes:

- **🧪 Test execution** – Automatically runs unit and integration tests (if present) for each commit
- **📦 Lint check configured with ESLint** – ESLint is integrated in the project as a static code analyzer
- **🔄 Trigger on push/MR** – The pipeline is launched for every push to the repository and merge request
- **🚨 Pipeline status visible per commit**

> ⚠️ **Note**: Although ESLint is configured in the project, it is **not yet launched** in the CI pipeline. This is a pending enhancement.

---

## ✅ Pipeline Visibility

You can monitor and inspect the status of all CI pipeline runs at the following link:

🔗 [GitLab CI Pipelines Dashboard](https://gitlab.pg.innopolis.university/f.markin/fah/-/pipelines)

Each entry shows:
- Commit hash
- Author
- Job status (passed/failed)
- Logs for debugging

---