import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import {
  combineLatest,
  map,
  Observable,
  of,
  startWith,
  switchMap,
  tap,
} from 'rxjs';
import { Message, ProfileUser} from 'src/app/models/Users';
//ChatsService, UsersService

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
})

export class HomeComponent implements OnInit {
  searchControl = new FormControl('');
  messageControl = new FormControl('');
  chatListControl = new FormControl('');


  constructor(
  ) {}

  ngOnInit(): void {

  }

  createChat(user: ProfileUser) {

  }

  sendMessage() {

  }

  scrollToBottom() {

  }
}
