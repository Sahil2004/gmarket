import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  Validators,
  ReactiveFormsModule,
  type ValidationErrors,
} from '@angular/forms';
import { Router, RouterLink } from '@angular/router';

import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'login',
  templateUrl: 'login.html',
  imports: [
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    ReactiveFormsModule,
    RouterLink,
  ],
})
export class Login {
  private router = inject(Router);
  private fb = inject(FormBuilder);
  private userService = inject(UserService);
  private _snackBar = inject(MatSnackBar);

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
    let email = this.email?.value;
    let password = this.password?.value;
    if (!email || !password) return;
    const res = this.userService.login(email, password);
    if (res) {
      this.router.navigate(['/watchlist']);
    } else {
      let snackBarRef = this._snackBar.open('Invalid email or password', 'Close', {
        duration: 3000,
      });
      snackBarRef.onAction().subscribe(() => {
        snackBarRef.dismiss();
      });
    }
  }
}
