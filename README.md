# Receipt Processor Challenge

A lightweight REST API built in **Go (Golang)** that processes receipts and returns a points score based on specified rules. The application uses **Docker** for containerization, includes **Swagger documentation**, and is deployed using **Render** with CI/CD via **GitHub Actions**.

---

## Tech Stack

- **Backend:** Go (Golang) with Gorilla Mux
- **API Documentation:** Swagger (via `swaggo/swag`)
- **Containerization:** Docker with Docker Hub
- **CI/CD:** GitHub Actions with Docker Build and Render Deployment
- **Deployment:** Render (https://receipt-processor-challenge-latest.onrender.com)

---

## Key Features

1. **Process Receipts API** (`POST /receipts/process`)  
2. **Get Points API** (`GET /receipts/{id}/points`)  
3. **Health Check API** (`GET /health`)  
4. **Swagger Documentation** available at:  
5. [Swagger UI](https://receipt-processor-challenge-latest.onrender.com/swagger/index.html)

---

## Test Coverage

**Coverage Summary:**
- **Service Layer:** 100% Coverage  
- **Route Layer:** 100% Coverage
```bash
receipt-processor-challenge % go test ./... -cover                             

        github.com/Nivedithabp/receipt-processor-challenge              coverage: 0.0% of statements
        github.com/Nivedithabp/receipt-processor-challenge/docs         coverage: 0.0% of statements
ok      github.com/Nivedithabp/receipt-processor-challenge/routes       (cached)        coverage: 100.0% of statements
ok      github.com/Nivedithabp/receipt-processor-challenge/services     (cached)        coverage: 100.0% of statements
```

---

## Docker Instructions

### Run Locally Using Docker Image in Docker hub
```bash
# Pull Docker image from Docker Hub
docker pull --platform=linux/amd64 bpniveditha/receipt-processor-challenge:latest

# Run Docker container
docker run -p 8080:8080 bpniveditha/receipt-processor-challenge:latest

#Access API locally
http://localhost:8080/swagger/index.html
```
## Run Locally Using Go
```bash
# Clone the repository
git clone https://github.com/Nivedithabp/receipt-processor-challenge.git
cd receipt-processor-challenge

# Run the application
go mod tidy
go run main.go
```

##CI/CD Pipeline Overview
###GitHub Actions Workflow:

1. Run Unit Tests with Coverage.
2. Build and Push Docker Image to Docker Hub.
3. Trigger Deploy Hook in Render to Deploy Application.

##Render Deployment
###Live API Hosted on Render:

[Receipt Processor API](https://receipt-processor-challenge-latest.onrender.com/swagger/index.html)


## Screenshots

<img width="1145" alt="Screenshot 2025-03-28 at 8 05 32 AM" src="https://github.com/user-attachments/assets/ab396b5f-3ad7-4c4b-9171-d92ee7b3244f" />

<img width="519" alt="Screenshot 2025-03-28 at 8 04 14 AM" src="https://github.com/user-attachments/assets/9e84b628-5607-4a57-b396-1642a0c8952f" />

<img width="1440" alt="Screenshot 2025-03-28 at 8 02 47 AM" src="https://github.com/user-attachments/assets/37150990-2894-4b11-ba2e-c6fb58cc73b5" />


