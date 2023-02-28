import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormControl, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  //strict type checking so we bypass with !
  signupForm!: FormGroup;
  loginForm!: FormGroup;

  //http: HttpClient
  constructor(
    private formbuilder: FormBuilder,
    private authService: AuthService,
     private router: Router
    ) { }

  ngOnInit(): void {
    this.signupForm = this.createSignFormGroup();
    this.loginForm = this.createLogFormGroup();
  }

  //code largely from: https://github.com/Jon-Peppinck/social-posts
  createSignFormGroup(): FormGroup {
    return new FormGroup({
      name: new FormControl("", [Validators.required, Validators.minLength(2)]),
      email: new FormControl("", [Validators.required, Validators.email]),
      password: new FormControl("", [
        Validators.required,
        Validators.minLength(7),
      ]),
    });
  }

  createLogFormGroup(): FormGroup {
    return new FormGroup({
      name: new FormControl("", [Validators.required, Validators.minLength(2)]),
      password: new FormControl("", [
        Validators.required,
        Validators.minLength(7),
      ]),
    });
  }

  onSignUp(): void {
    this.authService.onSignUp(this.signupForm.value).subscribe((msg) => {
      console.log(msg);
    });
  }

  onLogin(): void {
    //authentication


    //take username and password
    //make HTTP request to backend to find it
    //if found, load up that person's homepage
    //if not, prompt for Sign Up (alert form)

//---------------------------------------------------------------------------
    //SPRINT 1 CODE:
    // const userExists = this.signupUsers.find(m => m.username == this.loginObj.username && m.password == this.loginObj.password);
    // if(userExists != undefined)
    // {
    //   alert('User login successfully');
    // }
    // else {
    //   alert('Wrong credentials');
    // }
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
