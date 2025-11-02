# 🌡️ Real-Time Temperature Monitoring System (Go Channels Project)

## 🧠 Overview
This project simulates a **real-world temperature monitoring system** using **Go channels**.  
A background *sensor goroutine* generates random temperature readings and sends them through channels to a *main monitoring routine*.  
The main process:
- Receives sensor data,
- Detects slow responses,
- Times out after inactivity,
- Cancels and shuts down safely.

This project is part of **Module 2: Communication via Channels**, focusing purely on Go’s **channel fundamentals**.

---

## 🚀 Features & Concepts Covered

| Concept | Description |
|----------|-------------|
| **Channels (`make`)** | Used for sending and receiving temperature data |
| **Unbuffered Channels** | Used for `done` signal (stop communication) |
| **Buffered Channels** | Used for `tempChan` to store temperature readings |
| **Channel Direction** | Send-only (`chan<-`) and receive-only (`<-chan`) used in functions |
| **`select` Statement** | Handles multiple concurrent events |
| **Timeouts (`time.After`)** | Detects delays in sensor data |
| **Cancellation (`done` Channel)** | Stops sensor gracefully |
| **`close` Keyword** | Properly closes the temperature channel |
| **`range` Keyword** | (Optional) For iterating until channel closes |
| **Blocking Behavior** | Shown through unbuffered communication |

---

## 🧱 Project Structure

