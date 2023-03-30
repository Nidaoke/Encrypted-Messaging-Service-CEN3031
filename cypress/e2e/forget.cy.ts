describe('Click around to Forget Password', () => {
  it('lets pray this works', () => {
    cy.visit('http://localhost:4200/')
    cy.reload()

    cy.contains("Login").click()
    cy.get('button[name="login"]').click()

    // cy.get('input[formcontrolname="username"]')
    // .type('username2')

    // cy.get('input[formcontrolname="email"]')
    // .type('nana@emal.com')

    // cy.get('input[formcontrolname="password"]')
    // .type('p@55wOrD')

  })
})
