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
            18:32:56 Message broadcating handler started
            18:32:56 Server is running on port 50051...
            18:32:56 Server started on port 50051

    2. Start Clients 
        go run .\client\participants.go 'Name'

        Expected output:
            18:34:34 Connecting to server...
            18:34:34 Connected
            18:34:34 'Name' is trying to join the chat...
            18:34:34 Message stream established
            18:34:34 Joined the chat succesfully

            =Chit Chat=
            Write your message and press Enter
            Write 'exit' to leave
            ========

            > 
            [1760978074314] Server: Participant 'Name' joined Chit Chat at logical time 1760978074314
            > 18:34:34 Received Time: 1760978074314, From Server Content: Participant 'Name' joined Chit Chat at logical time 1760978074314

    2. Send message
        Hellow!

        Expected output:
            18:37:19 Send message: Hellow!
            > 
            [1760978239449] 'Name': Hellow!
            > 18:37:19 Received Time: 1760978239449, From 'Name' Content: Hellow!
    
    3. Leave the chat
        exit

        Expected output:
            18:38:39 Leaving the chat...
            [1760978319147] Server: Participant 'Name' left Chit Chat at logical time 1760978319147
            > 18:38:39 Received Time: 1760978319147, From Server Content: Participant 'Name' left Chit Chat at logical time 1760978319147
            18:38:39 Goodbye!

    4. Stop the server
        Ctrl+C 

        Expected output:
            18:40:20 Server shutting down...
            18:40:20 Message broadcasting handler stopped
            18:40:20 Server stopped