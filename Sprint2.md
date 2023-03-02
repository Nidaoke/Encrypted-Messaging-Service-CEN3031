# Sprint 2 Video:
    https://youtu.be/jGfwMNA987s
# Progress
    Significant progress has been made in furthering the development of our web application. 
    All leftover issues from Sprint1 have been completed, and our frontend and backend are fully integrated with each-other.
    There is successful routing to another page on the button.
    From the client-side, users can create an account, posting to the backend and updating the database.
    Additionally, the frontend can query the backend for a list of user accounts and their respective data.
    The User component was created to assist in the front/back end to make sure the posting and getting works. 
    We have also begun constructing unit tests and documentation for both our frontend and backend to ensure that everything is in working order.
    Finally, we have begun making progress towards secure account authentication, password/account retrieval, and the login process for our clients.

# Unit Tests
## Backend
#### Get and Post method testing:
    One unit test has been made to verify the integrity of both Backend methods (getting and posting accounts).
    The test creates a temporary database, and verifies our post method with an insertion of some user.
    It then tests the get method by getting the list of users from the database, and ensures both methods
    are working correctly by comparing the last user's information with the data we inputted to the post method.
    Finally it refreshes the temp database for further testing purposes.
## Frontend
#### Cypress Test:
    The Cypress Test implemented here is as if the user just started to sign up and is entering their username, email, password, and submits it.
    This is a simple test that just makes sure things are able to be typed in and button is able to be pressed. It also tests if the button routes
    to another page (the profile component) and to see if it displays. While it is supposed to store to the backend, this test does not test that. 
#### Login Component:
    The unit test is done with Cypress Component Testing, because after the Cypress Test, we were more familiar with how everything worked. Cypress
    also interferes with the Angular built in unit testing (Jasmine/Karma). The unit test performed on the Login Component checks
    to make sure it mounts correctly, the username input is valid, and to make sure the button is visible. The button had disappeared oddly during
    a few moments of coding and thus, this was actually important. 

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
