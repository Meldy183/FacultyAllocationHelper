# Quality Attribute Scenarios

This document describes key quality attributes critical to the success of the Faculty Allocation Helper. Each section outlines a sub-characteristic and provides detailed quality attribute scenarios.

---

## 🛠️ Reliability

### 🔁 Recoverability

- **Source**: A single user with permission to manage faculties and courses  
- **Stimulus**: The user, while performing routine work, accidentally makes an unintended change  
- **Artifact**: The Faculty Allocation Helper application  
- **Environment**: Normal operation conditions  
- **Response**: The system supports version control and maintains logs of all database changes, which are visible to other users  
- **Response Measure**: The user can inspect logs of recent changes and revert the system to a previous version if needed  

> 💬 _"Data preservation is a top priority for our customer. There should be no risk of data loss. To address this, the system includes full version control and logging mechanisms."_

---

## 🤝 Interaction Capability

### 🧩 Operability

- **Source**: A user responsible for managing course and faculty assignments  
- **Stimulus**: The user performs an operation (e.g., assigning a TA) using the Faculty Allocation Helper  
- **Artifact**: The user interface of the Faculty Allocation Helper  
- **Environment**: Regular working conditions  
- **Response**: The interface minimizes the number of clicks and page transitions required for common operations. Data is searchable, sortable, and filterable to improve accessibility and reduce manual effort  
- **Response Measure**: Core tasks such as assigning faculty or editing courses can be completed in no more than 5 interactions  

> 💬 _"The application's purpose is to streamline the process and reduce friction in user interaction with the system."_

---

### 🧠 Learnability

- **Source**: A new user, such as a faculty or institute representative  
- **Stimulus**: The user accesses the system for the first time to perform a task  
- **Artifact**: The Faculty Allocation Helper user interface  
- **Environment**: Standard operating environment, without prior training or documentation  
- **Response**: The interface is intuitive; users can immediately identify key actions and navigate through functionality without confusion  
- **Response Measure**: A user can complete a basic task (e.g., assigning a faculty member to a course) within 3 minutes without assistance or errors  

> 💬 _"The system should provide a smooth learning curve, especially since many users are transitioning from Excel workflows. Familiarity and ease of use are key."_

---