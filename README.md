# Ring Leader
## One-Stop Shop for Automatic resource management for organizations

_Started for the 2016 Gopher Gala_
Status: Not working. Look at [checklist](#Checklist)

### Motivations
- It's hard to keep track resources used by employees.
- It's hard to know when policies for resources.
- It's hard to know who is in charge of the resources.
- Easy to onboard and off board people for resources.

### About
- Ring Leader is meant to be an all-in-one solution for resource management.

### Running
- Only dependency is `docker` and `docker-compose`
- Run by `docker-compose up -d`
  - The API address can be found by typing `echo $(docker-machine ip $DOCKER_MACHINE_NAME):8080`
  - The RethinkDB UI can be found by typing `echo $(docker-machine ip $DOCKER_MACHINE_NAME):8081`

### Technology
- Redis for session management storage
- RethinkDB (trying for giggles) for subscription of live policies and resources to monitor
- Go Backend
- ? Frontend

### Checklist
#### Phase 1 (Must have for Gala)
- [x] Setup Docker-Compose for easy setup of servers
- [x] Basic Login Session
- [ ] Resources
  - [ ] Add Software License Resource
    - [x] Setup Software License Creation
    - [ ] Setup Software License Allocation
    - [ ] Setup Software License Deallocation
  - [ ] Add SSH Account Resource
    - [ ] Setup account creation
    - [ ] Setup account creation
    - [ ] Setup password expiration _maybe_
- [ ] Approval workflow
  - [ ] Setup Email for approvals
- [ ] Monitorer
  - [ ] Subscribe to resource and policy tables
  - [ ] Alert users about policy violations for resources
    - [ ] Set up E-mail alerts
- [ ] Policies
  - [ ] Capacity Warning
  - [ ] Time Expiration
- [ ] Web App UI
  - [ ] Login
  - [ ] CRUD Policies
  - [ ] CRUD Resources
- [ ] Tests

#### Phase 2 (Nice to have for Gala)
- [ ] Move Monitorer to it's own app instead of goroutine.
  - [ ] Setup in docker-compose script
- [ ] Refactor code to make more DRY

#### Phase 3 (Post Gala)
- [ ] Slack integration for monitorer
- [ ] Other Resources
- [ ] CLI Tool
