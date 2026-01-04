import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  Validators,
  ReactiveFormsModule,
  type ValidationErrors,
} from '@angular/forms';
import { Router } from '@angular/router';

import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'register',
  templateUrl: 'register.html',
  imports: [
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    ReactiveFormsModule,
  ],
})
export class Register {
  private router = inject(Router);
  private fb = inject(FormBuilder);
  private userService = inject(UserService);
  private _snackBar = inject(MatSnackBar);

  registerForm = this.fb.group({
    name: ['', Validators.required],
    email: ['', [Validators.required, Validators.email]],
    password: [
      '',
      [
        Validators.required,
        Validators.minLength(6),
        Validators.pattern('^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{6,}$'),
      ],
    ],
  });
  get name() {
    return this.registerForm.get('name');
  }
  get email() {
    return this.registerForm.get('email');
  }
  get password() {
    return this.registerForm.get('password');
  }
  getErrorMessage(controlName: string, errors: ValidationErrors | null | undefined): string {
    if (errors) {
      switch (Object.keys(errors)[0]) {
        case 'required':
          return `${controlName} is required`;
        case 'email':
          return `Not a valid email`;
        case 'minlength':
          return `${controlName} must be at least 6 characters long`;
        case 'pattern':
          return `${controlName} must contain at least one letter and one number`;
        default:
          return '';
      }
    }
    return '';
  }
  registerHandler() {
    let name = this.name?.value;
    let email = this.email?.value;
    let password = this.password?.value;
    if (!name || !email || !password) return;
    const res = this.userService.register(name, email, password);
    if (res) {
      this.router.navigate(['/dashboard']);
    } else {
      let snackBarRef = this._snackBar.open('Unable to register. Please try again.', 'Close', {
        duration: 3000,
      });
      snackBarRef.onAction().subscribe(() => {
        snackBarRef.dismiss();
      });
    }
  }
}
