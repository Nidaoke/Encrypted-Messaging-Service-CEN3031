import { ProfileComponent } from "src/app/components/profile/profile.component"
import { HttpClient, HttpClientModule } from '@angular/common/http';

it('checks that all buttons (edit profile, logout, and message) are visible', () => {
  cy.mount(ProfileComponent, {
    imports: [HttpClientModule],
    declarations: [ProfileComponent],
    providers: [HttpClient],
  })

    cy.get('img[name="pp"]').should("be.visible")
    cy.get('button[name="edit"]').should("be.visible")
    cy.get('button[name="logout"]').should("be.visible")
    cy.get('button[name="message"]').should("be.visible")
  })
