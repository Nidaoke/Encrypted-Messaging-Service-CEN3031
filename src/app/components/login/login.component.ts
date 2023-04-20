import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormControl, FormBuilder, FormGroup, Validators } from '@angular/forms';
//import { AuthService } from '../../services/auth.service';
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
    //private authService: AuthService,
    private http: HttpClient,
    private router: Router
    ) { }

  ngOnInit(): void {
    this.signupForm = this.createSignFormGroup();
    this.loginForm = this.createLogFormGroup();
  }

  //code largely from: https://github.com/Jon-Peppinck/social-posts
  createSignFormGroup(): FormGroup {
    return new FormGroup({
      username: new FormControl("", [Validators.required, Validators.minLength(2)]),
      email: new FormControl("", [Validators.required, Validators.email]),
      password: new FormControl("", [
        Validators.required,
        Validators.minLength(7),
      ]),
    });
  }

  createLogFormGroup(): FormGroup {
    return new FormGroup({
      username: new FormControl("", [Validators.required, Validators.minLength(2)]),
      password: new FormControl("", [
        Validators.required,
        Validators.minLength(7),
      ]),
    });
  }

  onSignUp(): void {
    //https://stackoverflow.com/questions/39698247/angular-2-form-serialization-into-json-format
    const url = 'http://localhost:9000/';

    let resource = JSON.stringify(this.signupForm.value);
    console.log('Add Button clicked: ' + resource);

    let formObj = this.signupForm.getRawValue();
    let serializedForm = JSON.stringify(formObj);

    if (this.signupForm.valid && this.signupForm.controls['username'].valid && this.signupForm.controls['email'].valid && this.signupForm.controls['password'].valid)
    {
      this.http.post(url, serializedForm).subscribe((data) => {
        console.log(data);}
      );

      this.router.navigate(['../profile']);
    }
    else
    {
      alert('Empty Item');
    }

    // this.authService.onSignUp(this.signupForm.value).subscribe((msg) => {
    //   console.log(msg);
    // });
  }

  onLogin(): void {
    //authentication


    //take username and password
    //make HTTP request to backend to find it
    //if found, load up that person's homepage
    //if not, prompt for Sign Up (alert form)

    const url = 'http://localhost:9000/login';
    const username = this.loginForm.get('username')!.value;
    const password = this.loginForm.get('password')!.value;

    this.http.post(url, { username, password }).subscribe(
      (response) => {
        // Successful authentication, redirect to the home page or user profile
        console.log('Authentication successful', response);
        this.router.navigate(['../profile']);
      },
      (error) => {
        // Failed authentication, show an error message
        console.log('Authentication failed', error);
        alert('Invalid username or password');
      }
    );
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
