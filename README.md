Payment Gateway Microservice
============================

This project implements a Payment Gateway Microservice that integrates with multiple payment gateways, handles deposit and withdrawal operations, and supports asynchronous transaction status updates using Kafka.

Table of Contents
-----------------

1.  [High-Level System Architecture](#system-design)
2.  [Request Flow](#request-flow)
3.  [Technology Stack](#technology-stack)
4.  [Setup Instructions](#setup-instructions)
    1.  [Pre-requisites](#pre-requisites)
    2.  [Clone the Repository](#clone-the-repository)
    3.  [Environment Variables](#environment-variables)
    4.  [Running the Application](#running-the-application)
5.  [Database Setup](#database-setup)
6.  [API Endpoints](#api-endpoints)
7.  [Kafka Integration](#kafka-integration)
8.  [Testing](#testing)

* * * * *

System Design
-------------

### 1\. High-Level System Architecture

The system is designed with the following key components:

-   **API Gateway**: Receives deposit and withdrawal requests from clients.
-   **Payment Gateway Interface**: Abstracts the interaction with different payment gateways using the **Strategy Pattern**.
-   **Transaction Manager**: Coordinates transactions, handles callbacks, and manages polling for transaction status updates.
-   **Background Worker (Polling Engine)**: Periodically checks the status of transactions that require polling.
-   **Database (PostgreSQL)**: Stores transaction details, statuses, and history.
-   **Message Queue (Kafka)**: Asynchronously processes callbacks and other background tasks.

* * * * *

Request Flow
------------

1.  **Client Request to API Gateway**

    -   The client initiates a deposit or withdrawal request. The system selects the gateway.
2.  **API Gateway Request Validation**

    -   The API Gateway validates the request (e.g., transaction_id, amount, currency). If validation fails, an error is returned.
3.  **Forward to Transaction Manager**

    -   The validated request is forwarded to the Transaction Manager.
4.  **Best Gateway Selection**

    -   The Transaction Manager selects the best gateway based on current load, gateway health, response times, and priority.
5.  **Request Formatting for Selected Gateway**

    -   The request is formatted as per the selected gateway's requirements (JSON for Gateway A, SOAP/XML for Gateway B) with a callback URL.
6.  **Transaction Storage**

    -   The transaction is stored in the database with a pending status and the gateway ID.
7.  **Client Acknowledgment**

    -   The API Gateway acknowledges the request with a pending status, transaction ID, and currency used.
8.  **Callback Handling (Callback-Supported Gateways)**

    -   The gateway sends a callback once the transaction is processed, and the Transaction Manager updates the status in the database.
9.  **Polling (For Non-Callback Gateways)**

    -   The Background Worker polls the gateway for pending transactions, retrieving the status and updating the database.

* * * * *

Technology Stack
----------------

-   **Go**: Main programming language.
-   **Kafka**: Message queue for asynchronous transaction status updates.
-   **PostgreSQL**: Database for transaction storage.
-   **Docker**: Containerization.
-   **Zookeeper**: Kafka coordination.

* * * * *

Setup Instructions
------------------

### Pre-requisites

-   **Go 1.16+**
-   **Docker & Docker Compose**
-   **Git**

### Clone the Repository

`git clone https://github.com/your-username/go-payment-gateway.git
cd go-payment-gateway`

### Environment Variables

Create a `.env` file in the root directory with the following:

`DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=pasword
DB_NAME=payment_db
KAFKA_BROKERS=kafka:9092
SERVER_PORT=8080`

### Running the Application

To run the application, use Docker Compose:

`docker-compose up --build`

This will start PostgreSQL, Kafka, and the Go application.

* * * * *

Database Setup
--------------

To initialize the database, migrations will run automatically.

* * * * *

API Endpoints
-------------

-   **Deposit**: `POST /transaction/deposit`

`{
  "amount": 100.0,
  "currency": "USD",
  "customer_id": "cust_01"
}`

-   **Withdrawal**: `POST /transaction/withdraw`

`{
  "amount": 50.0,
  "currency": "USD",
  "customer_id": "cust_02"
}`

* * * * *

Kafka Integration
-----------------

-   **Producer**: Sends transaction status updates to the `transaction-status-updates` topic.
-   **Consumer**: Listens for status updates and updates the database.

* * * * *

Testing
-------

-   **Postman** or **curl** can be used for testing API endpoints.
-   Check logs using:

`docker-compose logs -f app`
