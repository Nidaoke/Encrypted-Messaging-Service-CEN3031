import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './auth.guard';
import { LoginComponent } from './login/login.component';
import { ProfileComponent } from './profile/profile.component';

//https://www.youtube.com/watch?v=f3shwARuhEM&t=424s - guide on creating login
//https://blog.devgenius.io/angular-login-with-golang-and-mysql-92a59eec1249 - made the first page the profile page (will change later to login)
const routes: Routes = [
  { path: "", component: LoginComponent },
  { path: "profile", redirectTo: '/profile', pathMatch: 'full'},
  {
    path:'profile',
    component: ProfileComponent,
    canActivate: [AuthGuard]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
