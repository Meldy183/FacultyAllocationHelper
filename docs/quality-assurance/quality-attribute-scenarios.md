# Quality-atribute scenarios

## Reliability

### Recoverability

- **Source**: A single user with ability to manage faculties/courses
- **Stimulus**: The user doing his regular job and making a accidental mistake using the Faculty Allocation Helper
- **Artifact**: The Facculty Allocation Helper application
- **Environment**: The Faculty Allocation Helper application in normal conditions
- **Response**: The application supports version control AND creates logs of every change in the database those will be avaliable to other users
- **Response measure**: The user can check the logs of last changes and return the data to one of the previous versions in case of a mistake

One of the most important things for our customer is the data preservation. There should be no way the data can disappear permanently. For this purpose the app is going to have version control and logs of all changes

## Interaction Capability

### Operability

- **Source**: A user responsible for assigning faculty or managing courses
- **Stimulus**: The user interacts with the Faculty Allocation Helper to complete a task (assigning a TA for instance)
- **Artifact**: The Faculty Allocation Helper user interface
- **Environment**: The application in normal operating conditions
- **Response**:  The interface minimizes the number of clicks and page transitions required to complete common tasks, Data is searchable, filterable, and sortable to reduce time spent manually scanning through lists
- **Response measure**: All major operations (faculty assignment, course editing) can be completed in less than 5 interactions

The main goal of our application is to simplify the allocation process and user interaction with the database

### Lernability

- **Source**: A new user of the system, such as an institute or faculty representative
- **Stimulus**: The user opens the Faculty Allocation Helper interface for the first time to perform a task
- **Artifact**: The Faculty Allocation Helper user interface
- **Environment**: Normal conditions, without prior training or documentation
- **Response**: The user intuitively understands how to use the system: what actions are available, where to find needed features, and how they work
- **Response measure**: The user can complete a basic task (like assigning a faculty member to a course) within ≤ 3 minutes without assistance and without making mistakes

The app is going to be used by different people. Most of them used to do the same work using excel before so the transition should be comfortable for them
