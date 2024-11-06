# Order Book WebSocket Service

This monorepo contains both the frontend and backend components of a WebSocket-based service designed to connect to Binance, obtain the real-time trading order book data, calculate the average price, and broadcast it to connected clients. The frontend is built using SvelteJS, and the backend is developed in Golang.

## Project Overview

- **Backend:** Connects to Binance via WebSocket, processes the order book data in real-time, calculates the average order book price, and broadcasts this information to up to 10,000 clients.
- **Frontend:** Connects to the backend WebSocket service to receive the average price in real-time and displays it using a simple and intuitive user interface built with SvelteJS.
