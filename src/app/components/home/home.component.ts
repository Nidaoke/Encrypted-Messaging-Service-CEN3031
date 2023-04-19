import { Component, OnInit } from '@angular/core';
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

  //backend url
  url = 'http://localhost:9000/messages';

  chatMessages!: ChatMessage[];

  constructor(
    private http: HttpClient,
    private router: Router,
    private chatService: ChatService
  ) {}

  //retreive from back to show up in chat
  ngOnInit():void {
    this.chatService.getChatMessages().subscribe(
      (messages: ChatMessage[]) => {
        this.chatMessages = messages;
      },
      (error: any) => {
        console.log(error);
      }
    );
  }


  //sending the messages to the backend
  //tester1 to tester 2 only so far - later need to modify to user to friend
  sendMessage() {
    const message = this.messageControl.value;
    const url = 'http://localhost:9000/messages';

    this.http.post(url, {
      from: "Tester1",
      to: "Tester2",
      message: JSON.stringify(message)
    }).subscribe(() => this.messageControl.setValue(''));
  }

}
