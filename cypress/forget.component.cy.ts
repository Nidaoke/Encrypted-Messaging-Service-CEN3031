import { ForgetComponent } from "src/app/components/forget/forget.component"
import { HttpClient, HttpClientModule } from '@angular/common/http';

it('checks that all input boxes (name and email) and the recover button are visible', () => {
  cy.mount(ForgetComponent, {
    imports: [HttpClientModule],
    declarations: [ForgetComponent],
    providers: [HttpClient],
  })

    cy.get('input[name="nam"]').should("be.visible")
    cy.get('input[name="emal"]').should("be.visible")
    cy.get('button[name="recover"]').should("be.visible")
  })
