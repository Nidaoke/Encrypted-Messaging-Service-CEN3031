import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  signupUsers: any[] = [];

  signupObj: any = {
    username: '',
    email: '',
    password: ''
  };

  loginObj: any = {
    username: '',
    password: ''
  };

  constructor() {}

  ngOnInit(): void {
    const localData = localStorage.getItem('signupUsers');
    if(localData != null)
    {
      this.signupUsers = JSON.parse(localData);
    }
  }

  onSignUp() {
    this.signupUsers.push(this.signupObj);
    localStorage.setItem('signupUsers', JSON.stringify(this.signupUsers));

    this.signupObj = {
      username: '',
      email: '',
      password: ''
    };
  }
  //this needs fixing. It's displaying correctly
  onLogin() {
    const userExists = this.signupUsers.find(m => m.username == this.loginObj.username && m.password == this.loginObj.password);
    if(userExists != undefined)
    {
      alert('User login successfully');
    }
    else {
      alert('Wrong credentials');
    }
  }

}

// import { HttpClient } from '@angular/common/http';
// import { Component } from '@angular/core';
// import { Router } from '@angular/router';

// @Component({
//   selector: 'app-login',
//   templateUrl: './login.component.html',
//   styleUrls: ['./login.component.scss']
// })
// export class LoginComponent {
//   isLogin: boolean = false
//   registerFirstName: string | null = null
//   registerLastName: string | null = null
//   registerEmail: string | null = null
//   registerPassword: string | null = null
//   loginEmail: string | null = null
//   loginPassword: string | null = null
//   constructor(
//     private httpClient: HttpClient,
//     private router: Router
//   ){
//   }
//   register(){
//     console.log(this.registerFirstName, this.registerPassword)
//     this.httpClient.post('http://localhost:8080/register', {
//       firstName: this.registerFirstName,
//       lastName: this.registerLastName,
//       email: this.registerEmail,
//       password: this.registerPassword,
//     }).subscribe((response: any) => {
//       if(response){
//         localStorage.setItem('token', response.jwt)
//         this.router.navigate(['profile'])
//       }
//       this.registerFirstName = null
//       this.registerLastName = null
//       this.registerEmail = null
//       this.registerPassword = null
//     })
//   }
//   login(){
//     this.httpClient.post('http://localhost:8080/login', {
//       email: this.loginEmail,
//       password: this.loginPassword
//     }).subscribe((response: any) => {
//       if(response){
//         localStorage.setItem('token', response.jwt)
//         this.router.navigate(['profile'])
//       }
//       this.loginEmail = null
//       this.loginPassword = null
//     })
//   }
// }
