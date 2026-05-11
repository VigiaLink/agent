# agent

Remote camera relay agent for distributed AI-powered monitoring systems.

Agent is a lightweight service designed to securely connect local IP cameras to a remote monitoring server. It enables centralized video processing, AI analytics, and real-time monitoring without requiring direct public access to cameras or local networks.

The agent runs inside the client's local network, captures RTSP streams from IP cameras, and securely relays them to a remote server using low-latency streaming protocols.

---

# Features

- RTSP camera ingestion
- Remote video relay
- Low-latency streaming
- Automatic reconnection
- Lightweight resource usage
- Works behind NAT and CGNAT
- Multi-camera support
- Secure remote connectivity
- Designed for distributed AI surveillance systems

---

# Architecture

```text
IP Cameras
     ↓
   Agent
     ↓
Secure Remote Transport
     ↓
Central Monitoring Server
     ↓
AI Processing + Dashboard
