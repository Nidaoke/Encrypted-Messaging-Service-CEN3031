import { HomeComponent } from "src/app/components/home/home.component"
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { MatAutocompleteModule } from "@angular/material/autocomplete";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { ChatService } from 'src/app/services/chat.service';

////////////////////this test is on hold due to modifications
// it('Check to see if autocomplete pops up and can be clicked', () => {
//   cy.mount(HomeComponent, {
//     imports: [HttpClientModule, MatAutocompleteModule, BrowserAnimationsModule],
//     declarations: [HomeComponent],
//     providers: [HttpClient],
//   })

//   cy.get('input[formcontrolname="searchControll"]').click()
//   .wait(1000)
//   cy.contains("Testerbaby").click()

// })

// it('If the chat can have something be typed', () => {
//   cy.mount(HomeComponent, {
//     imports: [HttpClientModule, MatAutocompleteModule, BrowserAnimationsModule],
//     declarations: [HomeComponent],
//     providers: [HttpClient, ChatService],
//   })

//   cy.get('input[formcontrolname="messageControll"]').type("Cypress says Hello")
// })

it('Tester1 types to Tester2 via enter key', () => {
  cy.mount(HomeComponent, {
    imports: [HttpClientModule, MatAutocompleteModule, BrowserAnimationsModule],
    declarations: [HomeComponent],
    providers: [HttpClient, ChatService],
  })

  cy.get('input[formcontrolname="messageControll"]')
  .type("Cypress says Hello")
  .trigger('keydown', {
    key: 'Enter',
  });

  //refresh to see message pop up
  cy.reload()
})

it('Tester1 types to Tester2 via button press', () => {
  cy.mount(HomeComponent, {
    imports: [HttpClientModule, MatAutocompleteModule, BrowserAnimationsModule],
    declarations: [HomeComponent],
    providers: [HttpClient, ChatService],
  })

  cy.get('mat-mdc-button-touch-target')
  .type("Cypress says Hello")
  .trigger('keydown', {
    key: 'Enter',
  });

  //refresh to see message pop up
  cy.reload()
})
