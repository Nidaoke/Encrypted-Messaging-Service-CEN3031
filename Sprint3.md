# Sprint 3 Video:
    
# Progress
    From Sprint 2, the signup/login page is now more refined. 
    There is a Forget Password feature now implemented but will need further refinement. 
    A homepage that doubles as a chatroom is created and messages are now able to be sent. Additionally, the messages are able to be stored in
    the backend.
    The goal for next sprint is to be able to send/store longer messages, refine the visuals of various components, connect/route everything
    together, and hopefully get to the authentification during sign up.

# Unit Tests
## Backend
#### (UNIT) testing:
    (TBC)
## Frontend
#### Cypress Tests:
    Building off of Sprint 2's Cypress Test, this one checks to see if a user is able to login and send a message. It tests the updated login
    as well as the new homepage chat. Tester1 and Tester2 are dummy accounts made for this purpose.
#### Login Component:
    (TBC)
#### Forget Component:
    (TBC)
#### Home Component:
    This unit test is testing to see a string of max 200 characters is able to be typed and sent to a friend.

# Backend API Documentation
## GetAccounts
    Description:
        Returns a JSON array of all user accounts in the format detailed in Output. This includes their usernames, passwords (encoded), and emails.
        The list is returned in the order that the accounts were created (first created account is at index 0 of the array). This can be used to 
        verify if a user entered the correct password, or to check if a username has already been taken at account creation, or if friend requests
        and messages are going to a user that already exists.
        
    Request:
        localhost:9000/
        
    HTTP Type:
        GET
        
    Input:
        Null - As this is a GET request, there is no input to be provided.
        
    Output:
        JSON Array of user accounts in the following format:
            string 'username', string 'password', string 'email'
            Here, username represents the username of the created user, password represents the password of the created user, and email represents
            the email of the created user.
        HTTP 200
        
    Usage:
        For this method, the client should make an HTTP GET request to port 9000 of the server with no additional path (currently, this is at
        localhost:9000/). The client should be prepared to receive a list in the form of a JSON array, with each element containing three 
        attributes, one string called 'username, one string called 'password', and one string called 'email'.
        
## PostAccount
    Description:
        Allows for the addition of an account to the list of all accounts in the database. Essentially, this is used for account creation for the
        purpose of users, but can also be used for testing purposes (like testing CURL methods in the backend). When this POST request is made,
        a username, password, and email must all be supplied, or else an error will be thrown from the backend. This method will also return the
        JSON field representing the account that was just created.
        
    Request:
        localhost:9000/
        
    HTTP Type:
        POST
        
    Input:
        JSON field in the following format:
            string 'username', string 'password', string 'email'
            Here, username represents the username of the created user, password represents the password of the created user, and email represents
            the email of the created user.
        
    Output:
        JSON field in the following format:
            string 'username', string 'password', string 'email'
            Here, username represents the username of the created user, password represents the password of the created user, and email represents
            the email of the created user.
        HTTP 201
        
    Usage: 
        For this method, the client should make an HTTP POST request to port 9000 of the server with no additional path (currently, this is at
        localhost:9000/), whilst passing along a JSON field representing the parameters discussed in the Input section (must include a username,
        password, and email). The client can be prepared to receive a field representing the account that was just created, but this is not 
        necessary (may be helpful if numeric IDs are assigned to accounts). 
        
## GetMessages
    Description:
        Returns a JSON array of all messages in the format detailed in Output. This includes the usernames of the senders and recipients, the 
        message contents, and a numerical ID representing each message. The list is returned in the order that the messages were sent (first ever
        message to be sent will be at index 0), and the IDs are in the order that the messages were sent (first ever message to be sent will have
        an ID of 0). This can be used to see how many messages have been sent on the platform, for seeing how many messages a user has sent, for
        seeing how many messages two users have sent each other, and for displaying a users messages between them and their friends to said user.
        
    Request:
        localhost:9000/messages
        
    HTTP Type:
        GET
        
    Input:
        Null - As this is a GET request, there is no input to be provided.
        
    Output:
        JSON Array of all users' messages in the following format:
            int 'id', string 'from', string 'to', string 'message'
            Here, the ID represents the numerical id (order) of the message, from represents the sender of the message, to represents the 
            recipient of the message, and message represents the contents of the message. Note that 'message' has a maximum of 200 characters
            for now.
        HTTP 200
        
    Usage:
        For this method, the client should make an HTTP GET request to port 9000 of the server with an additional path of /messages (currently,
        this is at localhost:9000/messages). The client should be prepared to receive a list in the form of a JSON array, with each element
        containing four attributes, one integer called 'id', one string called 'from', one string called 'to', and one string called 'message'.
        
## GetMessage1
    Description:
        Returns a JSON array of all messages that were either sent by or received by some user in the format detailed in Output. It is a subset
        of the GetMessages function where either the 'from' field or the 'to' field are equal in string value to the name supplied in the path
        (see Request). The data returned includes the usernames of the senders and recipients (one of which must be the user supplied in the
        path), the message contents of these messages, and the numerical IDs representing each of these messages (in global order - the order
        found in GetMessages). The list is returned in the order that the messages were sent (first ever message sent by or to the user
        specified will be returned first. Likewise, this message will also have the lowest ID of the returned list of messages). This can be used
        as a shortcut to user checking from the GetMessages function, and is ultimately used to see a list of all the messages of some user, and 
        could be used (E.G) to see how popular some user is.
        
    Request:
        localhost:9000/messages/name
        Where name is the username of some user. E.g. localhost:9000/messages/jdikon37beforestbonanza1300and47touchdown would be the path to send
        a GET request to to obtain a list of all the messages that either had a 'from' value equal to "jdikon37beforestbonanza1300and47touchdown"
        or a 'to' value equal to "jdikon37beforestbonanza1300and47touchdown". Note that if both the 'from' value equals
        "jdikon37beforestbonanza1300and47touchdown" and 'to' value equals "jdikon37beforestbonanza1300and47touchdown", then the message will still
        be returned in the list requested by the GET request.
        
    HTTP Type:
        GET
        
    Input:
        Null - As this is a GET request, there is no input to be provided (the username is supplied in the path).
        
    Output:
        JSON Array of all users' messages where either the 'from' field or the 'to' field are equivalent to the user specified in the path in the 
        following format:
            int 'id', string 'from', string 'to', string 'message'
            Here, the ID represents the numerical id (order) of the message, from represents the sender of the message, to represents the 
            recipient of the message, and message represents the contents of the message. Note that 'message' has a maximum of 200 characters
            for now. Also note that either the 'from' field or the 'to' field must be equivalent to the user specified in the path.
        HTTP 200
        
    Usage:
        For this method, the client should make an HTTP GET request to port 9000 of the server with an additional path of /messages/name, where name
        is the username of the user that the client wants to collect the messages of (either sent by or received by) (currently, this is at 
        localhost:9000/messages/name where name is the username of the user that the client wants to collect the messages of (either sent by or
        received by)). The client should be prepared to receive a list in the form of a JSON array, with each element containing four attributes,
        one integer called 'id', one string called 'from', one string called 'to', and one string called 'message'. Note here that either the 'from'
        field or the 'to' field will be equivalent to the username specified in the path.
        
## GetMessage2
    Description:
        Returns a JSON array of all messages that were either sent by some user (hereby referred to as user1) and receieved by some user (hereby
        referred to as user2), or sent by user2 and receieved by user1, essentially a list of all messages that were sent between two users in
        the format detailed in Output. It is a subset of the GetMessages function where the 'from' field and the 'to' field match the two users
        specified in the path (see Request) in either order. The data returned includes the usernames of the senders and recipients (which should
        both match the usernames supplied in the path), the message contents of the messages, and the numerical IDs representing each of these 
        messages (in global order - the order found in GetMessages). The list is returned in the order that the messages were sent (first ever
        message sent between the two users specified in the path will be returned first. Likewise, this message will also have the lowest ID of the
        returned list of messages). This can be used as a shortcut to user checking from the GetMessages function, and is ultimately used to see
        a list of all the messages between two users, and could be used E.g. to see how much communication some user has had with one of their
        friends.
        
    Request:
        localhost:9000/messages/name1/name2
        Where name1 is the username of some user and name2 is the username of some other user.
        E.g. localhost:9000/messages/jdikon37beforestbonanza1300and47touchdown/jdikon37beforestbonanza1300and47fieldgoal would be the path to send
        a GET request to obtain a list of all the messages that either had a 'from' value equal to "jdikon37beforestbonanza1300and47touchdown" and
        a 'to' value equal to "jdikon37beforestbonanza1300and47fieldgoal" or a 'from' value equal to "jdikon37beforestbonanza1300and47fieldgoal" and
        a 'to' value equal to jdikon37beforestbonanza1300and47touchdown".
        
    HTTP Type:
        GET
        
    Input:
        Null - As this is a GET request, there is no input to be provided (the username is supplied in the path).
        
    Output:
        JSON array of all users' messages where either the 'from' field is equivalent to the first username specified in the path and the 'to'
        field is equivalent to the second username specified in the path, or the 'from' field is equivalent to the second username specified in
        the path and the 'to' field is equivalent to the first username specified in the path in the following format:
            int 'id', string 'from', string 'to', string 'message'
            Here, the ID represents the numerical id (order) of the message, from represents the sender of the message, to represents the 
            recipient of the message, and message represents the contents of the message. Note that 'message' has a maximum of 200 characters
            for now. Also note that both the 'from' field and the 'to' field must be equivalent to the users specified in the path, in either order.
        HTTP 200
        
    Usage:
        For this method, the client should make an HTTP GET request to port 9000 of the server with an additional path of /messages/name1/name2,
        where name1 and name2 are the usernames of two users that the client wants to collect the messages that were sent between said two (in
        which the order of who sent the message does not matter) (currently, this is at localhost:9000/messages/name1/name2 where name1 and name2
        are the usernames of two users that the client wants to collect the messages that were sent between said two (in which the order of who
        sent the message does not matter)). The client should be prepared to receive a list in the form of a JSON array, with each element
        containing four attributes, one integer called 'id', one string called 'from', one string called 'to', and one strng called 'message'. Note
        here that both the 'from' field and the 'to' field will match the usernames supplied in the path, in either order.
