# Agent

Remote camera relay agent for distributed AI-powered monitoring systems.

Agent is a lightweight service that securely connects local IP cameras to a centralized remote monitoring server. It enables remote video streaming, centralized AI processing, event detection, and real-time monitoring without requiring direct public access to cameras or local networks.

The agent runs inside the client's local network, captures RTSP streams from IP cameras, and securely relays them to a remote monitoring server using low-latency streaming protocols.

---

# Features

- RTSP camera ingestion
- Remote video relay
- Low-latency streaming
- Automatic reconnection
- Lightweight resource usage
- Multi-camera support
- Secure remote connectivity
- NAT and CGNAT friendly
- Designed for distributed AI surveillance systems
- Centralized monitoring support
- Real-time stream forwarding
- Fault-tolerant streaming pipeline

---

# Architecture

```text
┌─────────────────┐
│   IP Cameras    │
└────────┬────────┘
         │ RTSP
         ▼
┌─────────────────┐
│      Agent      │
│ Local Network   │
└────────┬────────┘
         │ Secure Stream Relay
         ▼
┌─────────────────┐
│ Remote Server   │
│ AI Processing   │
│ Dashboard       │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Mobile / Web UI │
└─────────────────┘
```

---

# Why Agent Exists

Most surveillance cameras operate inside private local networks and are not directly accessible over the internet.

Traditional approaches often require:
- Public IP addresses
- Port forwarding
- VPN setup
- Expensive cloud infrastructure
- Complex network configurations

Monitra Agent solves this problem by securely relaying local camera streams to a centralized monitoring server while remaining lightweight, simple, and infrastructure-efficient.

---

# Core Concepts

## Local Camera Access

The agent runs inside the same local network as the cameras and accesses RTSP streams directly.

Example:

```text
rtsp://192.168.0.20:554/stream1
```

---

## Remote Relay

The agent securely forwards video streams to a centralized remote server.

This enables:
- Remote live viewing
- AI processing
- Event detection
- Alert generation
- Centralized management

---

## Distributed Monitoring

Monitra Agent enables monitoring cameras across multiple remote locations from a single centralized system.

---

# Use Cases

- Retail stores
- Warehouses
- Small businesses
- Offices
- Industrial facilities
- Condominiums
- Distributed surveillance systems
- Smart monitoring systems
- Remote AI video analytics

---

# Tech Stack

## Core

- Go
- FFmpeg
- SRT
- RTSP

## Streaming

- MediaMTX
- WebRTC (planned)

## Connectivity

- Tailscale (planned)
- WireGuard (planned)

---

# Planned Features

- ONVIF auto-discovery
- Automatic camera provisioning
- Local configuration interface
- Multi-camera management
- Stream health monitoring
- Adaptive bitrate
- Hardware acceleration
- Automatic reconnection
- Secure encrypted transport
- WebRTC support
- OTA updates
- Docker support
- Multi-platform builds
- Event forwarding
- Local buffering
- Stream recording support

---

# Development Roadmap

## Phase 1

- RTSP ingestion
- Stream relay
- SRT transport
- Remote server connectivity

## Phase 2

- Multi-camera support
- Automatic reconnection
- Configuration management
- Local service mode

## Phase 3

- ONVIF discovery
- Local web UI
- Monitoring metrics
- Health checks

## Phase 4

- Hardware acceleration
- WebRTC support
- Adaptive streaming
- Edge-side optimizations

---

# Repository Structure

```text
agent/
├── cmd/
├── internal/
├── configs/
├── scripts/
├── ffmpeg/
├── streams/
├── network/
├── utils/
├── main.go
└── README.md
```

---

# Goals

Monitra Agent is designed to be:

- Lightweight
- Reliable
- Scalable
- Easy to deploy
- Infrastructure-efficient
- NAT-friendly
- Production-ready

The project focuses on simplicity, reliability, and centralized AI-powered monitoring.

---

# Status

Early development / proof of concept.

---

# License

MIT License

Copyright (c) 2026 VigiaLink

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND.
