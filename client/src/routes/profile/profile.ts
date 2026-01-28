import { Component, computed, inject, signal } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { ChangePasswordDialog, ConfirmationDialog, InputImage } from '../../components';
import { UserService } from '../../services/user.service';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import type { IUserData } from '../../types/user-data.types';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute, Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { toSignal } from '@angular/core/rxjs-interop';

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
  readonly router = inject(Router);
  readonly userService = inject(UserService);
  readonly fb = inject(FormBuilder);
  readonly _snackBar = inject(MatSnackBar);
  readonly _dialog = inject(MatDialog);

  private route = inject(ActivatedRoute);
  private data = toSignal(this.route.data);

  private userUpdate = signal<IUserData | null>(null);
  user = computed(() => {
    if (this.userUpdate() !== null) {
      return this.userUpdate() as IUserData;
    } else {
      console.log(this.data()?.['userData']);
      return this.data()?.['userData'] as IUserData;
    }
  });

  profilePhotoUri = signal<string | null>(this.user().profile_picture_url || null);

  get nameInitials(): string {
    if (!this.user()) return '';
    const nameArr = this.user().name.split(' ');
    const initials = nameArr.map((n) => n.charAt(0).toUpperCase());
    return initials.join('');
  }

  profileForm = this.fb.group({
    name: [this.user().name, [Validators.required]],
    email: [this.user().email, [Validators.required, Validators.email]],
    phoneNumber: [this.user().phone_number, [Validators.minLength(10), Validators.maxLength(10)]],
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

  updateProfilePhoto() {
    return (newImage: string) => {
      this.userService.updateProfile({ profile_picture_url: newImage }).subscribe({
        next: (res) => {
          let snackBarRef = this._snackBar.open('Profile photo updated successfully', 'Close', {
            duration: 3000,
          });
          snackBarRef.onAction().subscribe(() => {
            snackBarRef.dismiss();
          });
          this.userUpdate.set(res as IUserData);
          this.profilePhotoUri.set(this.user().profile_picture_url || null);
        },
      });
      return true;
    };
  }

  updateProfileHandler() {
    if (this.profileForm.valid) {
      this.userService
        .updateProfile({
          name: this.name?.value as string,
          email: this.email?.value as string,
          phone_number: this.phoneNumber?.value as number | undefined,
        })
        .subscribe({
          next: (res) => {
            let snackBarRef = this._snackBar.open('Profile updated successfully', 'Close', {
              duration: 3000,
            });
            snackBarRef.onAction().subscribe(() => {
              snackBarRef.dismiss();
            });
          },
        });
    }
  }

  openChangePasswordDialog() {
    this._dialog.open(ChangePasswordDialog, {
      data: {
        onSubmit: (oldPassword: string, newPassword: string) => {
          this.userService.changePassword(oldPassword, newPassword).subscribe({
            next: (res) => {
              let snackBarRef = this._snackBar.open('Password changed successfully', 'Close', {
                duration: 3000,
              });
              snackBarRef.onAction().subscribe(() => {
                snackBarRef.dismiss();
              });
            },
          });
        },
      },
    });
  }

  logoutHandler() {
    this.userService.logout().subscribe({
      next: (res) => {
        this.router.navigate(['/login']);
      },
    });
  }

  deleteAccountHandler() {
    this._dialog.open(ConfirmationDialog, {
      data: {
        title: 'Delete Account',
        message: 'Are you sure you want to delete your account? This action cannot be undone.',
        confirmText: 'Delete',
        cancelText: 'Cancel',
        onConfirm: () => {
          this.userService.deleteAccount().subscribe({
            next: (res) => {
              this.router.navigate(['/login']);
            },
          });
        },
      },
    });
  }
}
