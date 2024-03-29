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
    //take username and password
    //make HTTP request to backend to find it
    //if found, load up that person's homepage
    //if not, prompt for Sign Up (alert form)

    const url = 'http://localhost:9000/checklogin';
    let formObj = this.loginForm.getRawValue();
    let serializedForm = JSON.stringify(formObj);

    this.http.post(url, { serializedForm }).subscribe(
      (response: Object) => {
        // Cast the response to an object with a value property
        const responseData = response as { value: string };
        // Check the response string and handle it accordingly
        const responseString = responseData.value;

        // If the response is 'goodpassword', authentication is successful and redirect to home page
        if (responseString === 'goodpassword') {
          console.log('Authentication successful', response);
          this.router.navigate(['/home']);
        }
        // If the response is 'baduser' or 'badpassword', authentication failed and prompt for invalid credentials
        else if (responseString === 'baduser') {
          console.log('Authentication failed: invalid usename', response);
          alert('Invalid usename');
        }
        else if ( responseString === 'badpassword') {
          console.log('Authentication failed: invalid password', response);
          alert('Invalid password');
        }
        else {
          // If the response is neither 'goodpassword', 'baduser' nor 'badpassword', show an error message
          console.log('Authentication failed: unknown response', response);
          alert('Unknown response');
        }
      },
      (error) => {
        // Failed authentication, show an error message
        console.log('Authentication failed', error);
        alert('Cannot make it to back');
      }
    );
  }

}
