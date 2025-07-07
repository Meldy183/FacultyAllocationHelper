# Faculty Allocation Helper



## Usage

To test and touch it you should register or sign in.
1) To sign up use innopolis mail, e.g. m.pezer@innopolis.university
2) The password should be 9 or more symbols - Not necessarily the same as in your innopolis account
3) If you are already signed up in OUR web application - use existing mail and pass - it is checked by the system of authentication

## Architecture

### Static view

- UML Component Diagram: ![Static View Diagram](docs/architecture/static-view/ComponentUML.png)
- [PlantUML source](docs/architecture/static-view/StaticView.puml)

The architecture is based on microservices. Each of them uses its own database. Communication between the services is implemented via HTTP calls. This approach simplifies testing and replacing separate components while making changes.

### Dynamic view

- UML Sequence Diagram: ![Dynamic View Diagram](docs/architecture/dynamic-view/UserCaseDiagram.jpeg)

This diagram shows interactions between different user types and services in our system

### Deployment view

## Development

### Kanban board

Link to our Jira board: [https://facultyhelper.atlassian.net/jira/software/projects/SCRUM/boards/1]

- **To do** - [The task is assigned to somebody]
- **In progress** - [The work on a task is started]
- **Done** - [The task is completed]

### Git workflow

The chosen workflow is gitflow

- The issues are marked by the followwing lables: [frontend, backend, UX/UI, devops]
- Any issue can be assigned by any member at any moment
- The branches are created and named after each microservice, like authorization, parsing or profile service. There are also separate branches for frontend and backend generally
- Commits are made in free form by any member
-  Pull requests can be created by any member for the purposes of requesting changing the functionality
- Any member's code can be reviewed by another one
- Any member can merge pull requests for other members
- The issues are resolved by their assignees

#### Gitgraph diagram

[gitgrapgh](docs/development/git-workflow/gitgraph.png)

### Secrets management

So far, the secrets are transmitted manually when configuring the server. They are not stored in the repository and the main branch is protected.

### Automated testing

For unit testing, Jest was chosen as the main tool. It is widely
used for testing React applications, and Next.js which is based on React is well com-
patible with it. An extension for Jest was used for proper work with Next.js.

Beyond the standard testing package in Go, Ginkgo, GoConvey and Testify were con-
sidered. However, the Go standard package was chosen as the only one because it is
lightweight, fast, and tightly integrated with the Go toolchain (go test).

[The integration tests in the repo](FAH-auth-service/tests/integration)
[The unit tests in the repo](fah-frontend/__test__)

### Continious integration

[The CI workflow file](fah-frontend/.gitlab-ci.yml)
[Link to where all CI workflow runs can be seen](https://gitlab.pg.innopolis.university/f.markin/fah/-/pipelines)
ESlint is included as a static analyzer but not launched

## Quality

The quality characteristics are discussed with the customer

### Reliability

#### Recoverability

It is important to save the data somewhere so that it doesn't disappear in case of any error
### Flexibility

#### Adaptability

The product should be able to launch in any invironment.

### Performance efficiency

#### Capacity

Our product should have a limit to which the working process meets the requirements
