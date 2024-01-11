# GoLang Simple API Project

This is a simple GoLang API project that demonstrates basic API functionality.

## Getting Started

### Prerequisites

Make sure you have Go installed on your machine. If not, you can download and install it from [https://golang.org/dl/](https://golang.org/dl/).

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Tergeltengis/todogo.git
    ```

2. Change into the project directory:

    ```bash
    cd todogo
    ```

3. Run the application:

    ```bash
    go run .
    ```

The API server will start running at `http://localhost:8080`.

## API Endpoints

### 1. Get All Tasks

- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Get a simple hello message.
- **Response:**
    ```json
    [
      {
        "id": "1",
        "title": "go to grociery"
      },
      {
        "id": "2",
        "title": "go to barber"
      },
      {
        "id": "3",
        "title": "get the keys"
      }
]
    ```
