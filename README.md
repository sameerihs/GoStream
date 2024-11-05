# 🔄GoStream: Real-Time WAV to FLAC Audio Conversion Service

Welcome to **GoStream**! GoStream is a backend service that enables real-time streaming of WAV audio to FLAC conversion through WebSocket communication. Designed with efficient concurrency and powered by `ffmpeg`, this service is optimized for handling audio conversions with speed and accuracy. GoStream is hosted on an AWS EC2 instance, ensuring high availability and scalability for real-world applications.

> 🌐 **Live Demo**: [http://54.87.57.157/](#)   *(link to client website)*

---

## Features

- **Real-Time Audio Conversion**: Stream WAV files to be converted instantly to FLAC format.
- **WebSocket Communication**: Experience seamless two-way communication, ideal for real-time applications.
- **Concurrency**: Leverages Go’s concurrency model to handle multiple audio streams simultaneously, ensuring performance and efficiency.
- **Modular Design**:
  - **Core Service Module**: Manages audio streaming input and output efficiently.
  - **Conversion Logic Module**: Utilizes `ffmpeg` to convert WAV audio data into FLAC format with precision.
  - **WebSocket Communication Module**: Handles real-time streaming to and from connected clients.
  - **Router Module**: Sets up API routing, making it easy to extend and integrate with different frameworks (like Fiber, Gin, or Beego).
  - **Deployment Module**: Configured for AWS EC2 deployment to provide a stable and scalable service.
- **Powered by ffmpeg**: Utilizes `ffmpeg` for high-quality audio processing and conversion.
- **Scalable and Robust**: Hosted on an AWS EC2 instance, making it suitable for production-grade applications.

---

## Requirements

- **Go** (Version 1.23.1 or later)
- **ffmpeg** (Version 7.1-full_build)
- **AWS EC2 Instance** (for hosting)

---

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/sameerihs/GoStream.git
   cd GoStream

2. **Install Go dependencies**:
   ``` bash
   go mod tidy

## Usage

1. **Run the Server**:
   ```bash
   go run main.go

2. **Access the WebSocket Endpoint**:
   ``` bash
   The server runs on port 8080, and the WebSocket endpoint can be accessed at ws://<AWS-EC2-IP>:8080/ws.
3. **Client-Side Applications**:
   ```bash
   The client website (linked above) allows users to upload WAV files directly and receive FLAC conversions in real-time.
