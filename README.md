There are three main steps to building a concurrent TCP server. Keep these three steps in mind:

1.  **Listen:** The server starts listening on a specific address and port to see if anyone connects.
2.  **Accept:** As soon as a client (user) requests a connection, the server accepts it.
3.  **Handle:** The server hands off that connection to a new **Goroutine** to handle it and returns to step 2 to wait for the next user.