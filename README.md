# TODO-API-Service
The todo api service acts as a wrapper around the **Google Calendar Api**. It exposes endpoints which allow it to behave like a Todo interface. On calling these endpoints with appropriate values, CRUD operations can be performed on users calendars.
The list of exposed endpoints are:
- **GET events**
- **CREATE event**
- **UPDATE event**
- **DELETE event**

The important files are:
1. golangci-lint-config.yaml contains the configuration for the golangci-lint.
2. .circleci contains the circleci configuration
3. main.go file is in ./todoAPI/main.go

<p align="right">
  <img src="./images/logo.png" alt="Shukra in Spirit" width="150">
  <p align="right">Learn. Grow. Prosper.</p>
</p>

How to set up dev env:
1. Deploy a docker container of golang:latest
2. Load Git credentials
3. Clone
4. Setup remote from vs code
5. Run main.go file