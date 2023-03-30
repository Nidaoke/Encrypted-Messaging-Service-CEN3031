# Sprint 3 Video:
    
# Progress
    From Sprint 2, the signup/login page is now more refined. 
    There is a Forget Password feature now implemented but will need further refinement. 
    A homepage that doubles as a chatroom is created and messages are now able to be sent. There is a search bar that will later be able to autofill a
    user's friends for selection. Additionally, the messages are able to be stored in the backend successfully. Tester1 and Tester2 are dummy accounts 
    made for this purpose.
    The goal for next sprint is to be able to send/store longer messages, refine the visuals of various components, connect/route everything
    together, and hopefully get to the authentification during sign up.

# Unit Tests
## Backend
#### GetAccounts:
    The GetAccounts unit tests have been broken up into multiple tests. The first, TestGetAccountsCount, verifies that the number of results returned
    from GetAccounts is equal to the number of accounts in the backend's database. The second, TestGetAccountsContent, verifies that the
    GetAccounts function returns the proper data types, and that the data filled within the form is accurate (to the last account added).
#### PostAccount:
    Additionally, the PostAccount tests have been broken up as well. The first, TestPostAccountCount, verifies that the number of accounts stored
    in the database increases by one when PostAccount is called (and does not return an error), essentially verifying that an account was added.
    The second, TestPostAccountContent, verifies that the data added to the database is accurately representative of the input data given
    when PostAccount is called (that the username, password, and email all match).
#### GetMessages:
    The first of the GetMessages tests is TestGetMessagesCount, which verifies that the number of messages in the database is equal to the number
    of items returned by the GetMessages function. The second, TestGetMessagesContent, verifies that the GetMessages function returns the
    proper data types, and that the data filled within the form is accurate (to the last message added).
#### PostMessage:
    somethingsomething
#### GetMessages1:
    The first of the GetMessages1 tests is TestGetMessages1Count, which verifies that the GetMessages1 function returns a list with the same size
    as the number of messages in the database with the specified user as either the sender or receiver. The second of the GetMessages1 tests is
    TestGetMessages1Content, which verifies that the GetMessages1 function returns the information of the last message sent by or two the specified
    user is correct. Specifically, it makes sure that the message contents of the last message by or to that user in the database equals the 
    message contents of the last item returned by the function.
#### GetMessages2:
    All of the same tests for GetMessages1 exist for GetMessages2; that is, TestGetMessages2Count and TestGetMessages2Content. The only difference
    is that these tests specifically verify the count and content of the messages in the database vs those returned by the GetMessages2Function
    based on the two users specified by the function call path, as opposed to just the one user specified in GetMessages1.
## Frontend
#### Cypress Tests:
    spec: Building off of Sprint 2's Cypress Test, this one checks to see if a user is able to sign up and be redirected to a Patrick profile page. 
    forget: This e2e test will slowly be filled out as we go on. Its end goal is to be able to send an authentification email and to change password successfully.
        As of right now, since the "forget password" option in underneath the "Login" section, we are just trying to see if the login section is able to be pulled up.
#### Login Component:
    Checks that the login component page provides the 3 input boxes for user signup (username, email, and password) and the 2 input boxes for user login (username and password). It also checks that the signup and login buttons are provided.
#### Forget Component:
    Checks that the forget component page provides the 2 input boxes (username and email) for user password recovery and that recover button is provided to make the request.
#### Profile Component:
    Checks that the profile component page provides the 3 buttons to either edit the profile, logout, or send a message.
#### Home Component:
    The first unit test is testing to see if the search bar is able to be clicked and if the autocomplete of "Testerbaby" would pop up. This is to make sure
        for future user friends it would pop their names up.
    The second unit test is testing to see if the input for the chat is able to be typed in. We had ran into the issue of not being able to click into the input.
    

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
        
## PostMessage
    Description:
        Allows for the addition of a message to the list of all messages in the database. Essentially, this is used to update the database when
        one user on the platform send a message to another user of the platform, but can also be used for testing purposed (like testing CURL
        methods in the backend). Then this POST request is made, two usernames representing both the sender and the receiver of the message,
        and the message contents must all be supplied, or else an error will be thrown from the backend. This method will also return the JSON
        field representing the message that was just created.
        
    Request:
        localhost:9000/messages
        
    HTTP Type:
        POST
        
    Input:
        JSON field in the following format:
            string 'from', string 'to', string 'message'
            Here, from represents the username of the user who is sending the message, to represents the username of the user who is receiving the
            message, and message represents the contents of the message being sent. Note that the message here must be capped to 200 characters or
            fewer.
        
    Output:
        JSON field in the following format:
            int 'id', string 'username', string 'password', string 'email'
            Here, id represents the numerical id of this message, from represents the username of the user who is sending the message, to represents
            the username of the user who is receiving the message, and message represents the contents of the message being sent. Note that the
            message here must be capped to 200 characters or fewer.
        HTTP 201
        
    Usage: 
        For this method, the client should make an HTTP POST request to port 9000 of the server with an additional path of /messages (currently
        this is at localhost:9000/messages), whilst passing along a JSON field representing the parameters discussed in the Input section (must
        include the usernames of both the sender and the receiver and the message contents, capped at 200 characters). The client can be prepared
        to receive a field representing the message that was just created, but this is not necessary.
