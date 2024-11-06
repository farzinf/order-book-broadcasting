# Order Book WebSocket Service

This project is a backend server developed in Golang that connects to Binance via WebSocket to obtain the real-time trading order book, calculates the average price, and broadcasts it to connected clients using WebSocket. A sample frontend UI built with SvelteJS can display this average price.

## Features

- Connects to Binance WebSocket to receive order book data in real-time.
- Computes the average price from the received order book data.
- Supports broadcasting the computed average price to a large number of clients (up to 10,000).
- Designed to be efficient with concurrent client handling using goroutines.

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/farzinf/order-book-broadcasting.git
   cd order-book-broadcasting/backend
   ```

2. Setup `Go Modules`:

   ```
   go mod init
   go mod tidy
   ```

3. Run the backend server:
   ```
   go run .
   ```

## Improvements

- **Error Handling and Logging:** Implement error tracing for better debugging support.
- **Parallel Data Processing:** Implement parallelism in processing data to improve performance.
- **Secure WebSocket Connection:** Implement origin checking to prevent CSRF attacks.
- **Add Unit Tests:** Implement comprehensive unit tests to ensure the reliability of the system.
