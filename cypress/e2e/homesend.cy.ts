describe('Go to home message page and type something to friend and show up', () => {
  it('sign up goes to profile', () => {
    cy.visit('http://localhost:4200/home')
    cy.reload()

    cy.get('input[formcontrolname="username"]')
    .type('username2')

    cy.get('input[formcontrolname="email"]')
    .type('nana@emal.com')

    cy.get('input[formcontrolname="password"]')
    .type('p@55wOrD')

    cy.get('button[name="signup"]').click()

    cy.url()
      .should('be.equal', 'http://localhost:4200/profile')
  })
})
