# HomeCloudClient

HomeCloud is a implementation of Server-Client on LAN. The main goal is to have a server which stores up-to-date files which clients / devices can sync up with or submit changes. The Files are shared between all hosts.

## Client what type of application

What type of application should the client application be? It can be a CLI application that we pass arguments to the binary and send to the server or we could be running when Obsidian or any other chosen editor is running. The main benefits of CLI is that it's simpler to implement and use, that's why in this project we take a CLI based approach. The name for the application is __HomeDrop__

### CLI usage

How should CLI be used, what are the commands, parameters and arguemnts? I envision usage like this: `hdrop [-options] filepath`
The options could be:

- For HTTP methods __-g, -p, -u, -d__ (PUT gets -u for update to not cause confusion)
- For HTTP headers, In the future might want to add more support for some headers other than the essentials.
- Special commands parameters like __-gt__ (get tree) which gets servers Merkle tree.

### Client Work Flow

1. User runs a command
2. Client takes the arguemnts and writes a request
3. Client starts a connection to server and writes to it
4. Receive the response from the server and read it.
5. Depending on what the method was the client might perform one of several actions, update file, compare trees, etc.
