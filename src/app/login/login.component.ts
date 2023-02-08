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

  onLogin() {
    const userExists = this.signupUsers.find(m => m.username == this.loginObj.username && m.password == this.loginObj.password);
    if(userExists != undefined)
    {
      alert('User login successfully');
    } else {
      alert('Wrong credentials');
    }
  }

}
