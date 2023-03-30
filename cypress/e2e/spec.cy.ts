describe('Real e2e testing UwU', () => {
  it('sign up goes to profile', () => {
    cy.visit('http://localhost:4200/')
    cy.reload()

    cy.get('input[formcontrolname="username"]')
    .type('username2')

    cy.get('input[formcontrolname="email"]')
    .type('nana@emal.com')

    cy.get('input[formcontrolname="password"]')
    .type('p@55wOrD')

    cy.get('button[name="signup"]').click()

    cy.url()
      .should('be.equal', 'https://localhost:4200/profile')
  })
})
