describe('Real e2e testing UwU', () => {
  it('sign up does stuff', () => {
    cy.visit('http://localhost:4200/:')
    cy.reload()

    cy.get('<input _ngcontent-mko-c116="" formcontrolname="username" name="nam" placeholder="Username" ng-reflect-name="username" class="ng-pristine ng-invalid ng-touched">').as('username')
    cy.wait(0)
    cy.get('@username').click({force:true}).type('username2')

  })
})
