# Frensmatria

Welcome to **Frensmatria** – an innovative peer-to-peer (P2P) network dedicated to exploring and expanding gematria knowledge through collaborative nodes and a unified relay system.

![Gematria Lookup](images/lookup.png)

Currently in development

## Getting Started

### Starting the P2P Service

To begin using the P2P tools, run the `main.go` file located in the `relay` folder.

```bash
go run main.go
```

#### Commands and Configuration

Here are the main parameters you can use to customize the execution:

- **relay**: Specifies the address of the relay to connect to. The default setting is `localhost:9090`.

- **node**: This is the ID generated by the relay upon connection. Use this ID to connect with another node in the network.

- **port**: Defines the local server port. The default is set to `6969`.

- **not-update**: Controls whether the database should be updated when the service starts. By default it will update.

#### Example Usage

```bash
go run main.go -relay "localhost:9090" -node "<generated ID>" -port "6969" -not-update
```

This example runs the service on `localhost` at port 8080, using the generated ID and updating the database at startup.

## Key Features

- **Node Communication**: Enables nodes to connect and share information.
- **SDP Relay System**: Facilitates sharing of SDP descriptors for P2P connectivity.
- **Connection Management**: Supports multiple connections across the network.
- **Event Handling**: Manages real-time updates and communication events.
- **Gematria Calculations**:
  - Synx
  - Alphanumeric Qabbalah
- **Web Interface**:
  - Interactive gematria calculator
  - Recent results display
- **Database Integration**: Stores gematria results for future access.
- **Node Information Sharing**: Allows for the exchange of events and data across nodes.

## Roadmap

Here's a list of completed and planned features:

- [x] **Node Communication**
- [x] **Relay for SDP Sharing**
- [x] **Multi-Connection Management**
- [x] **Event Handler**
- [x] **Gematria Calculations** (Synx and Alphanumeric Qabbalah)
- [x] **Web Interface** (for calculating and viewing gematria results)
- [ ] **Simple Chat** *(under consideration; may be removed)*
- [x] **Database Integration**
- [ ] **Authentication** (for secure access between nodes)
- [ ] **Additional Relays** (to enhance P2P connectivity)
- [ ] **Error Correction** with Solomon Codes

## Upcoming Enhancements

- Improvements to relay functionalities
- Expanded system for adding and sharing node information

[normie.webm](https://github.com/user-attachments/assets/b3b53278-b24c-48ba-8cdb-d1e5f1a99379)
