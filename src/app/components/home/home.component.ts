import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

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

  url = 'http://localhost:9000/messages';

  chatList: chatData[] = [];

  constructor(
    private http: HttpClient,
    private router: Router
  ) {}

  getData() {
    this.http.get<chatData[]>(this.url).subscribe((data: chatData[]) => {
      this.chatList = data;
    });
  }

  ngOnInit(): void {
  //   this.http.get(this.url).subscribe((result) => {
  //     this.apps = result as Applications[];
  // }
  }


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
