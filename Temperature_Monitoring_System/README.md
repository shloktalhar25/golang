🌡️ Real-Time Temperature Monitoring System (Go Channels Project)
🧠 Overview

This project simulates a real-world temperature monitoring system using Go channels.
A background sensor goroutine generates random temperature readings and sends them through channels to a main monitoring routine.
The main process:

Receives sensor data,

Detects slow responses,

Times out after inactivity,

Cancels and shuts down safely.

This project is part of Module 2: Communication via Channels, focusing purely on Go’s channel fundamentals.
