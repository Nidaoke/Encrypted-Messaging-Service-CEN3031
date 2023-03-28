import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

//import { AuthGuard } from './services/auth.guard';
import { LoginComponent } from './components/login/login.component';
import { ProfileComponent } from './components/profile/profile.component';
import { ForgetComponent } from './components/forget/forget.component';

//https://www.youtube.com/watch?v=f3shwARuhEM&t=424s - guide on creating login
//https://blog.devgenius.io/angular-login-with-golang-and-mysql-92a59eec1249 - made the first page the profile page (will change later to login)
const routes: Routes = [
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'profile',
    component: ProfileComponent
  },
  {
    path: 'forget',
    component: ForgetComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
