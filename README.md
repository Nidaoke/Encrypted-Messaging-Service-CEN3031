# Encrypt-Me
Encrypt-Me is a messaging web application secured by end-to-end encryption up to industry standard cryptographic processes. This is the final project for the software engineering course (CEN3031) at the University of Florida.

## Project Description
Our project is an instant messaging platform where users will be able to create accounts, add friends, and send them messages. All passwords and messages will be SHA-256 encrypted within our database to ensure security. Users will be able to view their friends and see the conversations they've had with them. Features such as removing/blocking friends and deleting messages will also be available. The thematic details are to be decided.

## Technical Details
This application uses **Angular** for the front-end and **Go** for the backend. All data will be stored persistently on a **MySQL** database.

## Members
Our team is composed of four members. [Jonathan Cunningham](https://github.com/Nidaoke) and [Marcus Lugrand](https://github.com/marcuslugrand) will be working on the back-end, while [Henry Rivero-Vera](https://github.com/henryriverovera) and [Karen Liu](https://github.com/KareO2) will be working on the front-end. For the purpose of the course, our group on Canvas is Group 6.

# Set-up
Required installments: Angular 13 for front, Go for back. 
(May also need to install CLI if not done so to run Angular.)

## Step 1
  Open the project in the terminal with the correct directory and boot up the backend first. 
  -> cd web (and press tab. This should bring you to the right section)
  In the terminal, execute the command 'go run .' to boot up in localhost:9000

## Step 2
  Once the backend is loaded, open a new terminal to the project and execute command 'ng serve' to start the project at localhost:4200

## Routes
  * **/home** Goes to the message home page 
  * **/profile** Goes to profile page for personal account editing
  * **/forget** Goes to the forget-password page to reset password and authentification
