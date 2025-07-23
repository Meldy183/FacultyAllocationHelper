# ✅ Automated Testing

Automated tests are a critical part of ensuring system reliability and maintainability. This project uses different testing frameworks and approaches for the frontend and backend, tailored to their respective technology stacks.

---

## 🧪 Frontend Testing (Next.js + React)

For testing the frontend application, which is built with **Next.js (React framework)**, we use:

### 🛠️ **Jest** as the testing framework
- Industry-standard tool for unit testing in React applications.
- Well-integrated with Next.js and supports a wide range of features including mocking, snapshot testing, and async tests.

### 🔌 **Jest + Next.js Compatibility**
- The project uses a custom Jest configuration (including necessary extensions/plugins) to ensure compatibility with the Next.js environment and features like dynamic routing.

📁 [Browse frontend unit tests](fah-frontend/__test__)

---

## ⚙️ Backend Testing (Go)

The backend is written in **Go**, and testing is based on the native Go toolchain. Several libraries were evaluated for richer test syntax and features, including:

- Ginkgo
- GoConvey
- Testify

However, the **standard `testing` package** was chosen for the following reasons:
- **Lightweight** and doesn't add external dependencies
- **Fast** and optimized for performance
- **Integrated** directly with Go toolchain (`go test`)
- **Well-documented** and commonly used in the Go ecosystem

---

## 🔄 Integration Testing

The backend also includes **integration tests**, which verify that components work together correctly. These tests are organized in a separate directory and typically simulate realistic flows across modules or services.

📁 [Browse backend integration tests](FAH-auth-service/tests/integration)

---