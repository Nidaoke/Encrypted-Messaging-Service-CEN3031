# Sprint 4 Video:
    
# Progress
    From Sprint 2, the signup/login page is now more refined. 
    There is a Forget Password feature now implemented but will need further refinement. 
    A homepage that doubles as a chatroom is created and messages are now able to be sent. There is a search bar that will later be able to autofill a
    user's friends for selection. Additionally, the messages are able to be stored in the backend successfully. Tester1 and Tester2 are dummy accounts 
    made for this purpose.
    The backend-database has been setup to allow for friend requests to be sent between users, and relationships established between them.
    This allows for users to target their messages and view them between specific friends that they have mutually added.
    Additionally, we've added some API functionality to ease the burdon of user authentication and verification from the frontend - more details can
    be seen in the CheckLogin function in the API documentation.
    The goal for next sprint is to be able to send/store longer messages, refine the visuals of various components, connect/route everything
    together, and hopefully get to the authentification during sign up.

# Unit Tests
## Backend
#### GetRequests:
    
#### GetRequestsFrom:
    
#### GetRequestsTo:
    
#### PostRequests:
    
#### GetFriends:
    
#### PostFriend:
    
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

## GetRequests:
    Description:
        Returns a JSON array of all the requests in the format detailed in Output. This includes the usernames of the senders and recipients,
        as well as a numerical ID representing each request. The list is returned in the order that the requests were sent (first ever sent request
        to be sent will be at index 0), and the IDs are all in the order that the requests were sent (first ever request to be sent will have an ID
        of 0). This can be used to see how many requests have been sent on the platform, for seeing how many requests a user has sent, for seeing
        how many unaccepted requests a user has, and for measuring the request to friendship (accepted requests) ratio.
        
    Request:
        localhost:9000/requests
        
    HTTP Type:
        GET
        
    Input:
        Null - As this is a GET request, there is no input to be provided.
        
    Output:
        JSON Array of all users' requests in the following format:
            int 'id', string 'sentfrom', string'sentto
            Here, the ID represents the numerical id (order) of the request, sentfrom represents the sender of the request, sentto represents
            the recipient of the request.
        HTTP 200

    Usage:
        For this method, the client should make an HTTP GET request to port 9000 of the server with an additional path of /requests (currently,
        this is at localhost:9000/requests). The client should be prepared to receive a list in the form of a JSON array, with each element
        containing three attributes, one integer called 'id', one string called 'sentfrom', and one string called 'sentto'.

## GetRequestsFrom:
    Description:
        Returns a JSON array of all the requests that were sent from one specific user in the format detailed in Output (more specifically,
        their usernames). This includes the usernames of the users who the requests were sent to. The list does not include an ID, though
        the IDs of the requests can be reverse tracked from GetRequests. The list is however still returned in the order that the requests
        were sent due to how the database is stored (first ever request sent by the specified user will be at index 0). This can be used
        to display to a user how many outgoing friend requests they have, and to whom those requests were sent. It can also be used to calculate
        the ratio between a user's outgoing and ingoing requests, or to see how many requests they have sent in relation to how many friends they
        have or how many of their sent requests have been accepted.
        
    Request:
        localhost:9000/requests/from/name
        Where name is the username of some user.
        E.g. localhost:9000/requests/from/jdikon37beforestbonanza1300and47touchdown would be the path to send a GET request to obtain a list of all
        the requests that have a 'sentfrom' value equal to "jdikon37beforestbonanza1300and47touchdown".
        
    HTTP Type:
        GET
        
    Input:
        Null - As this is a GET request, there is no input to be provided (the username is supplied in the path).
        
    Output:
        JSON array of all users' requests where the 'sentfrom' field is equivalent to the username specified in the path in the following format:
            string 'name'
            Here, 'name' represents the username of the user for whom the request was sent out for. Note here that 'name' has a maximum of 64
            characters, which is in-line with the maximum character field choice during user account creation. 
        HTTP 200
        
    Usage:
        For this method, the client should make an HTTP GET request to port 9000 of the server with an additional path of /requests/from/name,
        where name is the username of the user that the client wants to collect the requests that were sent from said user (currently, this is at
        localhost:9000/requests/name where name is the username of the user that the client wants to collect the requests that were sent from said
        user). The client should be prepared to receive a list in the form of a JSON array, where each element contains only one single attribute,
        one string called 'name' representing the username of the user for whom each request was sent to.
        
## GetRequestsTo:
    Description:
        Returns a JSON array of all the requests that were sent to one specific user in the format detailed in Output (more specifically,
        their usernames). This includes the usernames of the users who the requests were sent from. The list does not include an ID, though
        the IDs of the requests can be reverse tracked from GetRequests. The list is however still returned in the order that the requests
        were sent due to how the database is stored (first ever request sent to the specified user will be at index 0). This can be used
        to display to a user how many incoming friend requests they have, and from whom those requests were sent. It can also be used to calculate
        the ratio between a user's incoming and outgoing requests, or to see how many requests they have been sent in relation to how many friends
        they have or how many of their received requests have yet to be accepted.
        
    Request:
        localhost:9000/requests/to/name
        Where name is the username of some user.
        E.g. localhost:9000/requests/to/jdikon37beforestbonanza1300and47touchdown would be the path to send a GET request to obtain a list of all
        the requests that have a 'sentto' value equal to "jdikon37beforestbonanza1300and47touchdown".
        
    HTTP Type:
        GET
        
    Input:
        Null - As this is a GET request, there is no input to be provided (the username is supplied in the path).
        
    Output:
        JSON array of all users' requests where the 'sentto' field is equivalent to the username specified in the path in the following format:
            string 'name'
            Here, 'name' represents the username of the user from whom the request was sent. Note here that 'name' has a maximum of 64
            characters, which is in-line with the maximum character field choice during user account creation. 
        HTTP 200
        
    Usage:
        For this method, the client should make an HTTP GET request to port 9000 of the server with an additional path of /requests/to/name,
        where name is the username of the user that the client wants to collect the requests that were sent to said user (currently, this is at
        localhost:9000/requests/name where name is the username of the user that the client wants to collect the requests that were sent to said
        user). The client should be prepared to receive a list in the form of a JSON array, where each element contains only one single attribute,
        one string called 'name' representing the username of the user from whom each request was sent.

## PostRequests:
    Description:
        Allows for the addition of a request to the list of all requests in the database. Essentially, this is used to update the database when
        one user on the platform sends out a request to another user of the platform, but can also be used for testing purposes (like testing CURL
        methods in the backend). Then this POST request is made, two usernames representing both the sender and the receiver of the request,
        both of which must be supplied, or else an error will be thrown from the backend. This method will also return the JSON field representing
        the request that was just created.
        
    Request:
        localhost:9000/requests
        
    HTTP Type:
        POST
        
    Input:
        JSON field in the following format:
            string 'sentfrom', string 'sentto'
            Here, sentfrom represents the username of the user who is sending the request, and sentto represents the username of the username who is
            receiving the request. Note here that each username must be capped to 64 characters, which should not pose a problem as user accounts
            have a maximum username length of 64 characters.
        
    Output:
        JSON field in the following format:
            int 'id', string 'sentfrom', string 'sentto'
            Here, id represents the numerical id of this request, sentfrom represents the username of the user who is sending the request, sentto
            represents the username of the user who is receiving the request. Note here that each username must be capped to 64 characters, which
            should not pose a problem as user accounts have a maximum username length of 64 characters.
        HTTP 201
        
    Usage: 
        For this method, the client should make an HTTP POST request to port 9000 of the server with an additional path of /requests (currently
        this is at localhost:9000/requests), whilst passing along a JSON field representing the parameters discussed in the Input section (must
        include the usernames of both the sender and the receiver of the request, each of which are capped at 64 characters). The client can
        be prepared to receive a field representing the request that was just created, but this is not necessary.

## GetFriends:
    Description:
        Returns a JSON array of all the friend-relationships in the format detailed in Output. This includes the usernames of each of the friends
        in the pair, as well as a numerical ID representing each friendship. The list is returned in the order that the friend requests were accepted
        (first ever accepted request to be sent will be at index 0), and the IDs are all in the order that the friend requests were accepted
        (first ever request to be accepted will have an ID of 0). This can be used to see how many requests have been accepted on the platform, for
        seeing how many friends a user has, and for seeing who some user is friended to.
        
    Request:
        localhost:9000/friends
        
    HTTP Type:
        GET
        
    Input:
        Null - As this is a GET request, there is no input to be provided.
        
    Output:
        JSON Array of all users' friendships in the following format:
            int 'id', string 'user1', string 'user2'
            Here, the ID represents the numerical id (order) of the friend-relationship, user1 and user2 each represent the usernames of the users
            in the friendship, in ambiguous order-pairing.
        HTTP 200

    Usage:
        For this method, the client should make an HTTP GET request to port 9000 of the server with an additional path of /friends (currently,
        this is at localhost:9000/friends). The client should be prepared to receive a list in the form of a JSON array, with each element
        containing three attributes, one integer called 'id', one string called 'user1', and one string called 'user2'.

## PostFriend:
    Description:
        Allows for the addition of a friend-relationship to the list of all friendship pairs in the database. Essentially, this is used to update
        the database when one user on the platform accepts a request from another user of the platform, but can also be used for testing purposes
        (like testing CURL methods in the backend). Then this POST request is made, two usernames representing the pair of friends that have
        accepted each other's requests, both of which must be supplied, or else an error will be thrown from the backend. This method will also return
        the JSON field representing the friendship relationship pair that was just created.
        
    Request:
        localhost:9000/friends
        
    HTTP Type:
        POST
        
    Input:
        JSON field in the following format:
            string 'user1', string 'user2'
            Here, user1 and user2 both represents the usernames of the two usernames who have accepted each other's friend requests (in ambiguous
            ordering). Note here that each username must be capped to 64 characters, which should not pose a problem as user accounts
            have a maximum username length of 64 characters.
        
    Output:
        JSON field in the following format:
            int 'id', string 'user1', string 'user2'
            Here, id represents the numerical id of this friendship relationship pair, and user1 and user2 both represent the usernames of the two users
            who have accepted each other's friend requests. Note here that each username must be capped to 64 characters, which should not pose a problem
            as user accounts have a maximum username length of 64 characters.
        HTTP 201
        
    Usage: 
        For this method, the client should make an HTTP POST request to port 9000 of the server with an additional path of /friends (currently
        this is at localhost:9000/friends), whilst passing along a JSON field representing the parameters discussed in the Input section (must
        include the usernames of both of the users represented by the new friendship relationship, each of which are capped at 64 characters).
        The client can be prepared to receive a field representing the friendship relationship pair that was just formed, but this is not necessary.

## CheckLogin:
    Description:
        Allows for the verification of some user's account status, particularly that of whether or not their account already exists, and if the
        inputted password is the correct password to the user's account. Essentially, this is used for two main purposes. Firstly, during account
        creation, the client needs to verify that the account being created does not already exist (the username has not been used yet), so this
        function can be used for that purpose. Secondly, during the login-stage, the client needs to verify that the correct password has been used
        and this function can be used to verify that as well. It is done with a POST request containing the username and password to be verified, 
        and returns a string containing values detailed in Output.
        
    Request:
        localhost:9000/checklogin
        
    HTTP Type:
        POST
        
    Input:
        JSON field in the following format:
            string 'username', string 'password'
            Here, username represents the username of the account that is to be verified, and password represents the password of the account to
            be verified.
        
    Output:
        JSON field in the following format:
            string 'value'
            Here, value represents the specific value of the returned string. It can have one of three different types: "baduser", "badpassword", and 
            "goodpassword". "baduser" means that the inputted username does not have a correlated account in the database system. "badpassword" means that
            the inputted password does not match the account that correlates to the matching username. "goodpassword" means that the user's account with the
            matching username has been found, and the inputted password matches their chosen password.
        HTTP 201
        
    Usage: 
        For this method, the client should make an HTTP POST request to port 9000 of the server with an additional path of /checklogin (currently
        this is at localhost:9000/checklogin), whilst passing along a JSON field representing the parameters discussed in the Input section (must
        include the username and password of the account that the client is attempting to verify). The client should be prepared to receieve a field
        representing the response to the account verification, particularly that of a JSON field with a 'value' attribute that stores some string
        describing the results of the verification. This can be used for both account creation and logging in.
