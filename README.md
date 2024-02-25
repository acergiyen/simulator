# Simulator Application

This application creates simulations of basketball matches that start simultaneously.

## Used Libraries
 Gorm - An Object-Relational Mapping (ORM) library.
 Viper - Manages reading, writing, and managing configuration files.
 Counterfeiter - Enables automatic generation of fake objects (mocks).
 Ginkgo - Provides a structure to organize your tests and helps in creating readable and understandable test scenarios.
 Log - Utilized for logging functionalities.
  
## Run Locally
### From Docker

To build the image of the application created with the `make build` command.
To create the PostgreSQL image and run it in the same container with the `make setup` command.
After the application has started successfully, the results can be listed at `http://localhost:8080/`.


## Run Unit Test
Run unit tests with the `make unit test` command.

## CI/CD

The `github/workflow` has been utilized. 

