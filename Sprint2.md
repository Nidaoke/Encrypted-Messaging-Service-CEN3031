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
