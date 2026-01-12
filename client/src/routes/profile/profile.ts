import { Component, inject } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { InputImage } from '../../components';
import { UserService } from '../../services/user.service';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import type { IUserDataClient } from '../../types/user-data.types';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'profile',
  templateUrl: 'profile.html',
  imports: [
    MatGridListModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    ReactiveFormsModule,
    InputImage,
  ],
})
export class Profile {
  readonly userService = inject(UserService);
  readonly fb = inject(FormBuilder);
  readonly _snackBar = inject(MatSnackBar);

  get currentUser() {
    return this.userService.currentUser as IUserDataClient;
  }

  get nameInitials(): string {
    if (!this.currentUser) return '';
    const nameArr = this.currentUser.name.split(' ');
    const initials = nameArr.map((n) => n.charAt(0).toUpperCase());
    return initials.join('');
  }

  profileForm = this.fb.group({
    name: [this.currentUser.name, [Validators.required]],
    email: [this.currentUser.email, [Validators.required, Validators.email]],
    phoneNumber: [
      this.currentUser.phoneNumber,
      [Validators.minLength(10), Validators.maxLength(10)],
    ],
  });

  get name() {
    return this.profileForm.get('name');
  }

  get email() {
    return this.profileForm.get('email');
  }

  get phoneNumber() {
    return this.profileForm.get('phoneNumber');
  }

  getErrorMessage(controlName: string, errors: any): string {
    if (errors) {
      switch (Object.keys(errors)[0]) {
        case 'required':
          return `${controlName} is required`;
        case 'email':
          return `Not a valid email`;
        case 'minlength':
          return `${controlName} must be at least ${errors['minlength'].requiredLength} characters long`;
        case 'maxlength':
          return `${controlName} cannot be more than ${errors['maxlength'].requiredLength} characters long`;
        default:
          return '';
      }
    }
    return '';
  }

  updateProfileHandler() {
    if (this.profileForm.valid) {
      this.userService.updateProfile({
        name: this.name?.value as string,
        email: this.email?.value as string,
        phoneNumber: this.phoneNumber?.value as number | undefined,
      });
      let snackBarRef = this._snackBar.open('Profile updated successfully', 'Close', {
        duration: 3000,
      });
      snackBarRef.onAction().subscribe(() => {
        snackBarRef.dismiss();
      });
    }
  }
}
