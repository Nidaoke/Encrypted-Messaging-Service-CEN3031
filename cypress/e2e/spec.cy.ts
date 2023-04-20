describe('Real e2e testing UwU', () => {
  it('sign up and start talking to Tester2', () => {
    cy.visit('http://localhost:4200/')
    cy.reload()

    cy.get('input[formcontrolname="username"]')
    .type('e2etesting')

    cy.get('input[formcontrolname="email"]')
    .type('e2etesting@test.com')

    cy.get('input[formcontrolname="password"]')
    .type('p@55wOrD!')

    cy.get('button[name="signup"]').click()

    cy.url().should('be.equal', 'http://localhost:4200/profile')

    cy.get('button[name="message"]') // find the button with name "message"
    .click()

    cy.url().should('be.equal', 'http://localhost:4200/home')

    cy.get('input[formcontrolname="messageControll"]')
    .type("Cypress E2E says Hello")
    .trigger('keydown', {
      key: 'Enter',
    });
  })
})
