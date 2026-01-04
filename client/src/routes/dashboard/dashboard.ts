import { Component, inject } from '@angular/core';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { DESIGN_SYSTEM } from '../../config';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'dashboard',
  templateUrl: 'dashboard.html',
  imports: [MatToolbarModule, MatIconModule, MatButtonModule, RouterLink],
})
export class Dashboard {
  readonly ds = inject(DESIGN_SYSTEM);
  readonly title = this.ds.app.APP_NAME;
}
