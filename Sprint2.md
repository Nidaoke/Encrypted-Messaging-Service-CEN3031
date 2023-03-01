# Progress
    Significant progress has been made in furthering the development of our web application. 
    All leftover issues from Sprint1 have been completed, and our frontend and backend are fully integrated with each-other.
    From the client-side, users can create an account, posting to the backend and updating the database.
    Additionally, the frontend can query the backend for a list of user accounts and their respective data.
    We have also begun constructing unit tests and documentation for both our frontend and backend to ensure that everything is in working order.
    Finally, we have begun making progress towards secure account authentication, and the login process for our clients.

# Unit Tests
## Backend
#### Get and Post method testing:
    One unit test has been made to verify the integrity of both Backend methods (getting and posting accounts).
    The test creates a temporary database, and verifies our post method with an insertion of some user.
    It then tests the get method by getting the list of users from the database, and ensures both methods
    are working correctly by comparing the last user's information with the data we inputted to the post method.
    Finally it refreshes the temp database for further testing purposes.
## Frontend
#### Say something here:
    Say something here

# Backend API Documentation
## GetAccounts
    Request:
        localhost:9000/
        
    HTTP Type:
        GET
        
    Input:
        Null
        
    Output:
        JSON Array of user accounts in the following format:
            int 'id', string 'username', string 'password', string 'email'
        HTTP 200
        
## PostAccount
    Request:
        localhost:9000/
        
    HTTP Type:
        POST
        
    Input:
        JSON field in the following format:
            string 'username', string 'password', string 'email'
        
    Output:
        JSON field in the following format:
            int 'id', string 'username', string 'password', string 'email'
        HTTP 201
