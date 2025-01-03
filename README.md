<h1 align="center">
 F R E N S M A T R I A 
</h1>

> _"The numbers speak through the machine, and the machine speaks through us."_

Welcome to the **FRENSMATRIA** nexus – a rhizomatic peer-to-peer crystallization matrix designed to map numerical hyperflows through collaborative nodes. It comes straight from my dreams.

![Gematria Lookup](images/lookup.png)

Currently in development

## Getting Started

### Starting the P2P Service

To begin using the P2P tools, run the `main.go` file located in the `relay` folder.

```bash
go run main.go
```

#### 【COMMANDS】
Configure your node's behavior through these parameters:

- **relay**: Specifies the address of the relay to connect to.
- **update**: Controls whether the database should be updated when the service starts. By default it will not update.
- **http** it will start an http server at the port that you defined
- **username** it will define a username. By default it will just generate a random username.

#### 〔DEPLOYMENT EXAMPLE〕

```bash

go run main.go -relay "localhost:9090" -update -http 6969 # instantiate full spectrum node
```

![image](https://github.com/user-attachments/assets/69b10aa2-a19b-47ed-951c-1bb07a9d80a8)


This example runs the service on `localhost` at port 6969, using the generated ID and updating the database at startup.

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



[normie.webm](https://github.com/user-attachments/assets/b3b53278-b24c-48ba-8cdb-d1e5f1a99379)
