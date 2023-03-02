import { LoginComponent } from "src/app/components/login/login.component"
import { HttpClient, HttpClientModule } from '@angular/common/http';

it('mounts, username types, and checks for a button that is visible', () => {
  cy.mount(LoginComponent, {
    imports: [HttpClientModule],
    declarations: [LoginComponent],
    providers: [HttpClient])

    cy.get('input[formcontrolname="username"]').type('UserNameUnitTest')

    cy.get('button[name="signup"]').should("be.visible")
  })
