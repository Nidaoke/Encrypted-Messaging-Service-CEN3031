import { LoginComponent } from "src/app/components/login/login.component"
import { HttpClient, HttpClientModule } from '@angular/common/http';

it('checks that all input boxes and buttons (signup and login) are visible', () => {
  cy.mount(LoginComponent, {
    imports: [HttpClientModule],
    declarations: [LoginComponent],
    providers: [HttpClient],
  })

  cy.get('input[formcontrolname="username"]').type('UserNameUnitTest')
  cy.get('input[formcontrolname="username"]').should("be.visible")
  cy.get('input[formcontrolname="email"]').should("be.visible")
  cy.get('input[formcontrolname="password"]').should("be.visible")
  cy.get('button[name="signup"]').should("be.visible")

  cy.get('input[formcontrolname="loginUsername"]').should("be.visible")
  cy.get('input[formcontrolname="loginPassword"]').should("be.visible")
  cy.get('button[name="login"]').should("be.visible")
})

//Unit test to sign up
it('Sign up', () => {
    cy.mount(LoginComponent, {
      imports: [HttpClientModule],
      declarations: [LoginComponent],
      providers: [HttpClient],
    })

    cy.get('input[formcontrolname="username"]').type('UserNameUnitTest')
    cy.get('input[formcontrolname="email"]').type("usernameunittest@test.com")
    cy.get('input[formcontrolname="password"]').type("Password123")
    cy.get('button[name="signup"]').click()
})

//Continuation of Signup
it('Login', () => {
  cy.mount(LoginComponent, {
    imports: [HttpClientModule],
    declarations: [LoginComponent],
    providers: [HttpClient],
  })

  cy.get('input[formcontrolname="loginUsername"]').type('UserNameUnitTest')
  cy.get('input[formcontrolname="loginPassword"]').type("Password123")
  cy.get('button[name="login"]').click()
})

//Unit test to sign up
it('Sign up', () => {
  cy.mount(LoginComponent, {
    imports: [HttpClientModule],
    declarations: [LoginComponent],
    providers: [HttpClient],
  })

  cy.get('input[formcontrolname="username"]').type('UserNameUnitTest')
  cy.get('input[formcontrolname="email"]').type("usernameunittest@test.com")
  cy.get('input[formcontrolname="password"]').type("Password123")
  cy.get('button[name="signup"]').click()
})

//Continuation of Signup
it('Login', () => {
cy.mount(LoginComponent, {
  imports: [HttpClientModule],
  declarations: [LoginComponent],
  providers: [HttpClient],
})

cy.get('input[formcontrolname="loginUsername"]').type('UserNameUnitTest')
cy.get('input[formcontrolname="loginPassword"]').type("Password123")
cy.get('button[name="login"]').click()
})
