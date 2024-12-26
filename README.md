# Beego Cat Image Voting App
## Project Overview
This project implements a web application inspired by [**The Cat API**](https://thecatapi.com/). It allows users to fetch and vote on cat images dynamically, providing an interactive and responsive experience.  

---

## Features

- Fetch Dynamic Cat Images: Fetches cat images in real time using The Cat API.
- Voting System: Enables users to upvote or downvote their favorite images.
- Asynchronous API Calls: Optimized with Go channels for efficient data fetching.
- Responsive UI: Built with Vanilla JavaScript, CSS, and Bootstrap for smooth interactions.
- Backend: Developed using the Beego framework for robust and scalable functionality.
- Configuration Handling: Manages API keys and settings using Beego config.
- Cross-Platform Support: Compatible with Linux and Windows environments.
- High Test Coverage: Includes unit tests with coverage to ensure reliability.


## Prerequisites
Before running this project, ensure the following are installed on your system:
- Go (Golang): Version 1.18 or higher.
- Beego Framework: Installed globally [Beego](https://beego.wiki/docs/install/install/).
- Git: For cloning the repository.


## Installation

1. **Step 1: Install Go:**
- Download and install Go from [Official](https://go.dev/dl/).
- Check Installation:
 ```bash
   go version
```

2. **Step 2: Install Beego:**
    Beego is the framework used for this project, and Bee CLI is a development tool.
    - Open terminal and Run: 
    ```bash
    go get github.com/beego/beego/v2@latest
    ```

    Ensure your GOPATH is set correctly.

## GOPATH SetUp
1. **For Linux:**
    ```bash
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
    source ~/.bashrc
    go mod tidy #Version Compatibility
    ```

2. **For Windows:**
     Set GOPATH and add it to your system PATH:
    1. Open Environment Variables in System Properties.
    2. Add a new variable:
        - Variable Name: GOPATH
        - Variable Value: C:\Users\<YourUsername>\go or in which workspace you like, setup path of that workspace
    3. Edit the Path variable and add: %GOPATH%\bin Verify your setup:

     ```bash
    echo $GOPATH   # For Linux/MacOS
    echo $env:GOPATH   # For Windows PowerShell
    echo $env:Path 
    ```
To verify installation:

```bash
    bee version
```

If face any version compatibality use:
```bash
    go mod tidy
```


## Clone Repository
```bash
    https://github.com/Muntasir-Ayan/beego-project.git
    cd beego-project
```
- Run the appliation:
```bash
    bee run
``` 
- This will run on this url: http://localhost:8080/custom
- You can see all favourites images: http://localhost:8080/custom/favourites
- You can see all voted images: http://localhost:8080/custom/votes

## Unit Testing
Run unit tests to ensure the application's reliability:
1. Run Test files: 
```bash
    go test ./tests -v
``` 
2. Generate coverage report: 
```bash
    go test -coverprofile coverage.out ./...
    go tool cover -html coverage.out
``` 

## Cat-Api key Configuration:
you can Generate your api-key from [thecatapi.com](https://thecatapi.com/)

- use your api key at  `conf/app.conf` and update `catapi_key="give your key"`