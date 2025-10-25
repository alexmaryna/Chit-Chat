Chit Chat
How to run the program
    Installation
    1. Clone the repository
        git clone <>
        cd Chit-Chat
    2. Install dependencies
        go mod download

    Running the system
    1. Start the server
        go run .\server\chitChatService.go

        Expected output:
            15:57:27 Server started on port 50051

    2. Start Clients 
        go run .\client\participants.go 'Name'

        Expected output:
            18:34:34 Connecting to server...
            [LogTime] Server: Participant 'Name' joined Chit Chat

            =Chit Chat=
            Write your message and press Enter
            Write 'exit' to leave

    2. Send message
        > Hellow!

        Expected output:
            [LogTime] 'Name': Hellow!
    
    3. Leave the chat
        exit

        Expected output:
            [LogTime] Server: Participant 'Name' left Chit Chat

    4. Stop the server
        Ctrl+C 

        Expected output:
            18:40:20 Server shutting down...
            18:40:20 Server stopped