import { Component } from '@angular/core';

import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';

@Component({
  selector: 'login',
  templateUrl: 'login.html',
  imports: [MatCardModule, MatFormFieldModule],
})
export class Login {}
