# Go Go Go

- [My Go Projects](#my-go-projects)
  - [My Go Project #5](#my-go-project-5)
  - [My Go Project #4](#my-go-project-4)
  - [My Go Project #3](#my-go-project-3)
  - [My Go Project #2](#my-go-project-2)
  - [My Go Project #1](#my-go-project-1)
- [My Go Apps 😀](#my-go-apps-)
  - [text](#text)
  - [GO JSON Read \& Write](#go-json-read--write)
  - [Go Vue WebSockets Chat App](#go-vue-websockets-chat-app)
  - [Go WebSockets](#go-websockets)
- [📌 **Popular Go API Frameworks**](#-popular-go-api-frameworks)
  - [✅ **1️⃣ Gin**](#-1️⃣-gin)
  - [✅ **2️⃣ Echo**](#-2️⃣-echo)
  - [✅ **3️⃣ Fiber**](#-3️⃣-fiber)
  - [✅ **4️⃣ Chi**](#-4️⃣-chi)
  - [✅ **5️⃣ Go-kit**](#-5️⃣-go-kit)
  - [✅ **6️⃣ gRPC**](#-6️⃣-grpc)
- [✅ **Practical Tip**](#-practical-tip)
- [Go Testing frameworks](#go-testing-frameworks)
  - [✅ **1️⃣ Built-in `testing` package**](#-1️⃣-built-in-testing-package)
  - [✅ **2️⃣ `testify`**](#-2️⃣-testify)
  - [✅ **3️⃣ `gomock`**](#-3️⃣-gomock)
  - [✅ **4️⃣ `GoConvey`**](#-4️⃣-goconvey)
  - [✅ **5️⃣ `Ginkgo` + `Gomega`**](#-5️⃣-ginkgo--gomega)
  - [✅ **6️⃣ Integration Testing Tools**](#-6️⃣-integration-testing-tools)
  - [✅ **Key points:**](#-key-points)

## My Go Projects

### My Go Project #5

<!-- Statistics Canada -->

Contributed to an Open Source AI Application

- Tool Replacement for Cost Optimization: Identified and replaced a deprecated Go-based testing tool that had shifted from open-source to a paid license model. Successfully integrated a free, open-source alternative with equivalent capabilities, ensuring continued test coverage and CI pipeline integrity without incurring additional cost.

- Authentication Architecture Update: Adapted to upstream changes in the app's authentication system, transitioning from OpenID to OAuth 2.0. Updated Kubernetes (K8s) manifests, secrets, and ingress configurations accordingly. Modified Azure AD settings and redirect URIs to align with the new auth flow, ensuring seamless and secure user authentication under the new architecture.

### My Go Project #4

<!-- Hatfield -->

Created a web app to manage Dask Kubernetes Clusters and deployed to Azure with Golang REST API on the backend and VUE on the front-end

- Backend in Golang: Developed a robust backend using Golang and the Gin framework for creating lightweight, high-performance REST APIs to manage Dask Kubernetes Clusters. These APIs allowed for dynamic provisioning, scaling, and monitoring of clusters on Azure Kubernetes Service (AKS).
  - Implemented cluster management functions such as creation, scaling, and deletion.
  - Used Gorm to interact with the database for storing and retrieving cluster state and configurations.
  - Integrated Swaggo for automated API documentation, ensuring seamless developer onboarding and maintenance.
- Testing & Mocking: Employed Gomock for unit testing, creating mocks of interfaces for testing the behavior of cluster management logic, to ensure code quality and functionality. This improved test coverage and reliability.
- Cluster Monitoring: Integrated Kubernetes tools like Kubectl and Kubernetes Dashboard for real-time monitoring of Dask clusters. This allowed proactive cluster management and resource scaling based on performance metrics.
- Frontend with Vue.js and Vuetify: The user interface was built using Vue.js with Vuetify for a responsive and intuitive experience. This frontend allowed users to visualize and manage Kubernetes clusters efficiently, providing real-time updates on cluster status and usage.
- Containerization and Deployment: Both the backend API and frontend UI were containerized using Docker. The container images were stored in Azure ACR and deployed to Azure AKS. This enabled seamless, scalable deployment across environments.
- Exposed the Golang API to a public IP, enabling external access to the application.
- Postman for API Testing: Thoroughly tested the APIs using Postman, ensuring that all endpoints functioned correctly for cluster management operations like create, scale, delete, and monitor.

### My Go Project #3

<!-- Sunlife -->

Built the Canadian Dental Care Plan enrollment web app for Health Canada

- Backend: Developed a full-stack web application using Golang and RESTful API services, implementing business logic and integrating with the database layer using GORM (Go ORM) for seamless data management. Implemented secure authentication and authorization mechanisms using JWT and middleware patterns in Go.
- Frontend: Utilized React to build responsive and dynamic user interfaces, incorporating modern features such as React Hooks and Context API to manage state effectively. Integrated Axios for API communication and Material-UI to deliver a clean, user-friendly design. Ensured code quality and scalability by adhering to best practices in component-based development, and conducted comprehensive unit testing using Go's testing package for the backend and Jest for the frontend.

### My Go Project #2

<!-- Thales Group -->

Spearheaded the design and implementation of Cloud Functions and CI/CD data pipelines within Google Cloud for a prominent multinational corporation.

- Cloud Function Development: Leveraged Golang to craft robust cloud functions. Notably, configured one cloud function to monitor incoming Pub/Sub messages, efficiently storing them in BigQuery.
- Additionally, engineered another cloud function with Golang to seamlessly archive BigQuery data to a GCS bucket in Avro format using Cloud Scheduler.
- Infrastructure as Code: Orchestrated all GCP infrastructure using Terraform scripts to parameterize all cloud functions. Automated infrastructure provisioning and enabled auto-scaling, ensuring seamless scalability and resource efficiency.
- CI/CD Pipeline Design: Architected and executed CI/CD pipelines with precision using Cloud Source Repository, Cloud Build, and GCP Container Registry integrated with Docker and YAML. Configured triggers for test pipelines and production pipeline, enabling streamlined development workflows. Implemented concurrent execution for multiple test pipelines, optimizing development velocity.
- Dataflow Pipeline Expertise: Collaborated on Dataflow pipelines utilizing Apache Beam SDK for Java and Maven, ensuring efficient data processing and analysis.
- GitLab: Established GitLab merge request templates, checklists, and pipelines, facilitating efficient code management and review processes. Maintained code repositories with diligence, ensuring version control and collaboration excellence.

### My Go Project #1

<!-- RBC -->

Built a new web application with APIs and deployed to Azure Cloud

- Developed and deployed a cloud-native web application to Microsoft Azure, leveraging modern DevOps practices and microservices architecture.
- Designed and implemented RESTful APIs using Golang, integrating with Azure SQL, Redis, and Docker Compose for local orchestration.
- Built and automated CI/CD pipelines using Jenkins and shell scripts to deploy application components to Azure Cloud environments.
- Performed unit and integration testing using testify and automated test execution as part of the deployment pipeline.
- Onboarded APIs to the Apigee API Gateway, ensuring secure and scalable API management.
- Monitored application health and performance using Splunk and Dynatrace, providing actionable insights for performance tuning and issue resolution.
- Followed Infrastructure as Code (IaC) practices to provision infrastructure and support deployment automation with Terraform.
- Collaborated with cross-functional teams to adopt DevOps best practices, shorten release cycles, and enhance delivery pipelines.

<!--
### My Go Project #2

Developed end-to-end solutions for managing Dask Kubernetes Clusters, leveraging Full Stack REST APIs and VUE applications deployed on Azure

- Engineered RESTful APIs using Golang for efficient cluster management, ensuring robust functionality and scalability.
- Containerized backend APIs and UI applications using Docker, enabling seamless deployment and scalability.
- Orchestrated deployment processes using Azure ACR repository and Azure AKS, ensuring efficient management and scaling of resources.
  Exposed APIs securely to the public internet through Azure AKS, ensuring accessibility while maintaining security standards. -->

<!-- Hatfield -->
<!-- ### 2: Created Full Stack REST APIs (backend) and VUE (front-end) apps to manage Dask Kubernetes Clusters and deployed to Azure

- Coded the functions to manage Dask Kubernetes Clusters with GoLang and Python.
- Monitored the Clusters with Kubectl and dashboard.
- Programmed REST APIs with GoLang.
- Built the UI with VUE and Vuetify.
- Containerized the backend API app and UI app with Docker, managed with Azure ACR repository, deployed to Azure AKS and then exposed the API to a public IP. -->

<!-- Thales Group -->
<!-- ### 1: Designed and implemented Cloud functions and CI/CD data pipelines in Google Cloud for a large multinational company

- Programmed cloud functions in Golang and Python.
- Configured one cloud function to monitor incoming Pub/Sub messages then save to BigQuery, another cloud function to archive data from BigQuery to GCS bucket in avro format with Cloud Schedule.
- Created unit tests with go mock framework and integration tests for test coverage.
- Built a stress testing multi-threading tool for performance benchmark and adjusted Cloud Function configuration.
- Managed all GCP infrastructure as code with Terraform scripts, parameterized all cloud functions, automated everything and enabled auto-scaling.
- Designed and built the CI/CD pipelines with Cloud Source Repository, Cloud Build and GCP Container Registry with Docker and YAML; set up the triggers for the test pipelines and prod pipeline; configured multiple test pipelines can run concurrently.
- Handed over high quality detailed runbook documentation. -->

<!-- ### My Go Project #1

Led the design and implementation of Cloud Functions and CI/CD data pipelines in Google Cloud for a large multinational company

- Developed and programmed two robust cloud functions in Golang to automate critical processes: Configured one cloud function to monitor incoming Pub/Sub messages and seamlessly integrate them into BigQuery for real-time data analysis; Engineered another cloud function to efficiently archive data from BigQuery to a Google Cloud Storage (GCS) bucket in Avro format using Cloud Schedule, ensuring data integrity and accessibility.
- Implemented comprehensive unit tests with gomock and integration tests with testify to ensure the reliability and robustness of Golang codebase.
- Engineered a stress testing multi-threading tool to benchmark performance and fine-tune Cloud Function configurations for optimal scalability and efficiency.
- Managed all GCP infrastructure as code using Terraform scripts, enabling seamless deployment, scaling, and management of cloud resources. Parameterized all cloud functions for flexibility and automation.
- Architected and built CI/CD pipelines utilizing Cloud Source Repository, Cloud Build, and GCP Container Registry with Docker and YAML configurations.
- Configured triggers for test pipelines and production pipeline, ensuring seamless integration and continuous delivery of code changes.
- Implemented concurrent test pipelines for improved efficiency and faster feedback loops.- -->

## My Go Apps 😀

### [Go Concurrency Demo: Bulk Job Processor](go-ng-cc-app/README.md)

### [GO JSON Read & Write](GoJsonReadWrite/README.md)

### [Go Vue WebSockets Chat App](GoVueWebSocketsChat/README.md)

### [Go WebSockets](GoWebSocketsConsole/README.md)

## 📌 **Popular Go API Frameworks**

### ✅ **1️⃣ Gin**

- **Most popular** lightweight HTTP web framework.
- High performance — small overhead.
- Has powerful routing, middleware support, JSON validation.
- Commonly used for REST APIs.
- You already use Gin — so highlight that.

👉 _Example_: _“I used Gin to build REST APIs for cluster management — lightweight and fast.”_

---

### ✅ **2️⃣ Echo**

- Another fast, minimalist framework.
- Similar to Gin but slightly different syntax and middleware handling.
- Great for large REST APIs with versioning.
- Supports middleware chaining, JWT auth, CORS.

---

### ✅ **3️⃣ Fiber**

- Inspired by Express.js (Node.js).
- Focus on performance — built on top of Fasthttp.
- Very intuitive for people coming from JavaScript/Node background.

---

### ✅ **4️⃣ Chi**

- Lightweight, idiomatic router.
- Good for building modular, composable HTTP services.
- Built around Go's net/http — very standard library-friendly.

---

### ✅ **5️⃣ Go-kit**

- Not a web framework — more of a **toolkit for microservices**.
- Provides support for transport layers (HTTP, gRPC), logging, circuit breakers.
- Good when you need robust, production-grade services.

---

### ✅ **6️⃣ gRPC**

- Technically not a “framework” but a **protocol**.
- For high-performance APIs with strong contracts — uses Protocol Buffers.
- Ideal for internal microservices — faster than JSON REST.

<!-- ### 📌 **How to talk about it in the interview**

When asked _“How do you build APIs in Go?”_, you can say:

> **“For REST APIs, I prefer Gin because it's lightweight, fast, and has excellent middleware support — I've used it with Gorm for database interactions and Swaggo for auto-generating API docs.**
>
> **For RPC-style internal APIs, I'd use gRPC with Protocol Buffers for type safety and performance.**
>
> **If the system requires microservices features like circuit breakers, I'd look at integrating Go-kit.”** -->

---

## ✅ **Practical Tip**

- **How to choose REST vs gRPC**:

  - REST for external APIs: human-readable, easy testing.
  - gRPC for internal services: better performance, strong type contracts.

## Go Testing frameworks

### ✅ **1️⃣ Built-in `testing` package**

- Every Go project uses this — it's the **standard** unit testing library.
- Works with `go test` — supports benchmarks, example tests.
- No external dependency — good for simple tests.

---

### ✅ **2️⃣ `testify`**

- The **most popular 3rd-party testing toolkit**.
- Provides:

  - **`assert`**: better assertions (e.g., `assert.Equal`, `assert.NoError`).
  - **`require`**: fail immediately.
  - **`mock`**: mocking interfaces and dependencies.

- Makes test code cleaner and more readable.

👉 _Example: You use `testify` — mention that to show best practices._

---

### ✅ **3️⃣ `gomock`**

- The **official mocking framework** by Google.
- Works well when you have interfaces — auto-generates mocks.
- Useful for **unit tests** when isolating components (e.g., DB calls, external services).

👉 _Example: You used `gomock` to mock Kubernetes cluster managers — perfect example for them!_

---

### ✅ **4️⃣ `GoConvey`**

- Adds a **web UI** for test output — useful for big suites.
- Good for **BDD-style** testing (e.g., `So()`, `Convey()` syntax).

---

### ✅ **5️⃣ `Ginkgo` + `Gomega`**

- Popular for **BDD testing** in larger Go projects.
- Expressive syntax for complex scenarios.
- Often used for integration and E2E tests, e.g., Kubernetes controllers, cloud-native apps.

---

### ✅ **6️⃣ Integration Testing Tools**

- **`httptest`** (built-in) — test HTTP handlers and API endpoints.
- **Postman/Newman** — for manual or automated API tests outside the code.
- **Docker Compose** — often used to spin up test DBs/services.

<!-- ---

## 📌 **How to talk about it in your interview**

If they ask **“How do you test Go code?”**, you can say:

> **“I follow TDD as much as possible, starting with the built-in `testing` package. For richer assertions and clearer test output, I use `testify`.**
>
> **When I need mocks — like when isolating an API handler from a DB call — I use `gomock` to generate interface mocks.**
>
> **I also use `httptest` for integration testing of HTTP handlers, and sometimes Postman for end-to-end API tests.”** -->

---

### ✅ **Key points:**

- **Unit tests** → `testing`, `testify`, `gomock`
- **Integration tests** → `httptest`, Postman, Docker Compose
- **Mocking** → `gomock`, `testify/mock`
- **BDD/E2E** → `Ginkgo`, `Gomega`
