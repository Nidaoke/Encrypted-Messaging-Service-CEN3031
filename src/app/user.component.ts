import { Component, Input } from '@angular/core';
import { Users, User } from './user';
import { UserService } from './user.service';

@Component({
  selector: 'user',
  templateUrl: './user.component.html',
  styles: [
    `
      h1 {
        font-family: Lato;
      }
    `,
  ],
})
export class UserComponent {
  users: Users;

  constructor(public userService: UserService) {
    this.users = {} as Users;
  }

  ngOnInit(): void {
    this.users.data = [];
    this.userService.getUsers().subscribe((response: any) => {
      debugger;
      for (let i = 0; i < response.length; i++){
        var user = {} as User;
        user.username = response[i].username;
        user.password = response[i].password;
        this.users.data.push(user);
      }
    });
  }
}
