# Go Go Go

- [My Go Projects](#my-go-projects)
  - [My Go Project #5](#my-go-project-5)
  - [My Go Project #4](#my-go-project-4)
  - [My Go Project #3](#my-go-project-3)
  - [My Go Project #2](#my-go-project-2)
  - [My Go Project #1](#my-go-project-1)
- [My Go Apps 😀](#my-go-apps-)

## My Go Projects

### My Go Project #5

<!-- Statistics Canada -->

?: contributed to an Open Source AI Application

- Tool Replacement for Cost Optimization: Identified and replaced a deprecated Go-based testing tool that had shifted from open-source to a paid license model. Successfully integrated a free, open-source alternative with equivalent capabilities, ensuring continued test coverage and CI pipeline integrity without incurring additional cost.

- Authentication Architecture Update: Adapted to upstream changes in the app’s authentication system, transitioning from OpenID to OAuth 2.0. Updated Kubernetes (K8s) manifests, secrets, and ingress configurations accordingly. Modified Azure AD settings and redirect URIs to align with the new auth flow, ensuring seamless and secure user authentication under the new architecture.

### My Go Project #4

<!-- Hatfield -->

?: Created a web app to manage Dask Kubernetes Clusters and deployed to Azure with Golang REST API on the backend and VUE on the front-end

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

? : Built the Canadian Dental Care Plan enrollment web app for Health Canada

- Backend: Developed a full-stack web application using Golang and RESTful API services, implementing business logic and integrating with the database layer using GORM (Go ORM) for seamless data management. Implemented secure authentication and authorization mechanisms using JWT and middleware patterns in Go.
- Frontend: Utilized React to build responsive and dynamic user interfaces, incorporating modern features such as React Hooks and Context API to manage state effectively. Integrated Axios for API communication and Material-UI to deliver a clean, user-friendly design. Ensured code quality and scalability by adhering to best practices in component-based development, and conducted comprehensive unit testing using Go's testing package for the backend and Jest for the frontend.

### My Go Project #2

Thales Group

?: Spearheaded the design and implementation of Cloud Functions and CI/CD data pipelines within Google Cloud for a prominent multinational corporation.

- Cloud Function Development: Leveraged Golang to craft robust cloud functions. Notably, configured one cloud function to monitor incoming Pub/Sub messages, efficiently storing them in BigQuery.
- Additionally, engineered another cloud function with Golang to seamlessly archive BigQuery data to a GCS bucket in Avro format using Cloud Scheduler.
- Infrastructure as Code: Orchestrated all GCP infrastructure using Terraform scripts to parameterize all cloud functions. Automated infrastructure provisioning and enabled auto-scaling, ensuring seamless scalability and resource efficiency.
- CI/CD Pipeline Design: Architected and executed CI/CD pipelines with precision using Cloud Source Repository, Cloud Build, and GCP Container Registry integrated with Docker and YAML. Configured triggers for test pipelines and production pipeline, enabling streamlined development workflows. Implemented concurrent execution for multiple test pipelines, optimizing development velocity.
- Dataflow Pipeline Expertise: Collaborated on Dataflow pipelines utilizing Apache Beam SDK for Java and Maven, ensuring efficient data processing and analysis.
- GitLab: Established GitLab merge request templates, checklists, and pipelines, facilitating efficient code management and review processes. Maintained code repositories with diligence, ensuring version control and collaboration excellence.

### My Go Project #1

<!-- RBC -->

?: Built a new web application with APIs and deployed to Azure Cloud

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

[GO JSON Read & Write](GoJsonReadWrite/README.md)

[Go Vue WebSockets Chat App](GoVueWebSocketsChat/README.md)

[Go WebSockets](GoWebSocketsConsole/README.md)
