import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  Validators,
  ReactiveFormsModule,
  type ValidationErrors,
} from '@angular/forms';

import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'login',
  templateUrl: 'login.html',
  imports: [
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    ReactiveFormsModule,
  ],
})
export class Login {
  private fb = inject(FormBuilder);
  loginForm = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', Validators.required],
  });
  get email() {
    return this.loginForm.get('email');
  }
  get password() {
    return this.loginForm.get('password');
  }
  getErrorMessage(controlName: string, errors: ValidationErrors | null | undefined): string {
    if (errors) {
      switch (Object.keys(errors)[0]) {
        case 'required':
          return `${controlName} is required`;
        case 'email':
          return `Not a valid email`;
        default:
          return '';
      }
    }
    return '';
  }
  loginHandler() {
    console.log(this.loginForm.value);
  }
}
