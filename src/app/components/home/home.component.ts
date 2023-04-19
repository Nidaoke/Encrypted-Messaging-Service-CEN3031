import { Component, OnInit, ViewChild, ElementRef, ChangeDetectorRef } from '@angular/core';
import { FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';


import { ChatService } from '../../services/chat.service';
import { ChatMessage } from '../../services/chat-message';

interface chatData {
  msg: string;
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
})

export class HomeComponent implements OnInit {
  searchControl = new FormControl('');
  messageControl = new FormControl('');
  chatListControl = new FormControl('');
  chatMessages!: ChatMessage[];
  friends: any[] = []; // initialize friendsList as an empty array

  //for testing purposes this is base profile
  currentUser: string = 'Tester1';

  constructor(
    private http: HttpClient,
    private router: Router,
    private chatService: ChatService,
    private cdr: ChangeDetectorRef
  ) {}

  //retreive from back to show up in chat
  ngOnInit():void {
    //friends to show up on side
    //this.getFriends();

    //chat to show up
    this.getChat();
  }

  //retreive meesages from backend
  getChat(){
    this.chatService.getChatMessages().subscribe(
      (messages: ChatMessage[]) => {
        this.chatMessages = messages;
      },
      (error: any) => {
        console.log(error);
      }
    );
  }

  // //retreives friends from backend
  // getFriends(){
  //   this.http.get<any[]>(`/friends/${this.currentUser}`).subscribe((response) => {
  //     this.friends = response;
  //   });
  // }

  //sending the messages to the backend
  sendMessage() {
    const message = this.messageControl.value;
    const url = 'http://localhost:9000/messages';

    this.http.post(url, {
      from: this.currentUser,
      to: "Tester2",    //later change to friend
      message: JSON.stringify(message)
    }).subscribe(() =>{
      this.messageControl.reset();
      this.cdr.detectChanges(); // detect changes to update chat interface
    });
  }

}
